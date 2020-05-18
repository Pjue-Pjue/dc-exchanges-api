package okex

import (
	dc_exchanges_api "dc-exchanges-api"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

type FutureClient struct {
	client *dc_exchanges_api.Client
}

func NewFutureOKEXClient(apiKey, secretKey, passphrase string, endPoint string) *FutureClient {
	return &FutureClient{client: newOKExClient(apiKey, secretKey, passphrase, endPoint)}
}

// 获取盘口
func (c *FutureClient) GetOrderBook(instrumentId string, options map[string]string) (OrderBook, error) {
	var ob OrderBook
	var orderBook FutureOrderbookResult
	params := map[string]string{}
	if options != nil {
		if v, ok := options["size"]; ok {
			params["size"] = v
		}
		if v, ok := options["depth"]; ok {
			params["depth"] = v
		}
	}
	requestPath := BuildParams(GetInstrumentIdUri(FUTURES_ORDERBOOK, instrumentId), params)
	_, _, err := c.client.Request("GET", requestPath, nil, &orderBook)

	timestamp, _ := time.Parse("2006-01-02T15:04:05.000Z", orderBook.Timestamp)
	ob.Timestamp = timestamp.Local()
	for _, v := range orderBook.Asks {
		ob.Asks = append(ob.Asks, BookItem{
			Price:     ParseFloat(v[0]),
			Size:      ParseFloat(v[1]),
			NumOrders: ParseInt(v[2]),
		})
	}
	for _, v := range orderBook.Bids {
		ob.Bids = append(ob.Bids, BookItem{
			Price:     ParseFloat(v[0]),
			Size:      ParseFloat(v[1]),
			NumOrders: ParseInt(v[2]),
		})
	}
	sort.Slice(ob.Asks, func(i, j int) bool {
		return ob.Asks[i].Price < ob.Asks[j].Price
	})
	sort.Slice(ob.Bids, func(i, j int) bool {
		return ob.Bids[i].Price > ob.Bids[j].Price
	})
	return ob, err
}

// 获取指数信息
func (c *FutureClient) GetIndexInfo(instrumentId string) (IndexInfo, error) {
	var result IndexInfo
	_, _, err := c.client.Request("GET", GetInstrumentIdUri(INDEX, instrumentId), nil, &result)
	return result, err
}

// 获取法币汇率
func (c *FutureClient) GetRate() (RateResult, error) {
	var result RateResult
	_,_, err := c.client.Request("GET", RATE, nil, &result)
	return result, err
}

// 获取预估交割价
// 交割前1小时才有返回值
func (c *FutureClient) GetEstimatedPrice(instrumentId string) (EstimatedPriceResult, error) {
	var result EstimatedPriceResult
	_, _, err := c.client.Request("GET",GetInstrumentIdUri(FUTURES_ESTIMATED_PRICE, instrumentId), nil, &result)
	return result, err
}

// 获取合约信息
func (c *FutureClient) GetFutureInstrumentByCurrency(base string, quote string) ([]Instrument, error) {
	var instruments []Instrument
	_, _, err := c.client.Request("GET", FUTURES_INSTRUMENTS, nil, &instruments)
	if err != nil {
		return []Instrument{}, err
	}
	var result []Instrument
	for _, v := range instruments {
		if strings.EqualFold(v.BaseCurrency, base) && strings.EqualFold(v.QuoteCurrency, quote) {
			result = append(result, v)
		}
	}
	return result, nil
}

// 单个合约持仓信息
func (c *FutureClient) GetPositionByInstrumentId(instrumentId string) (FuturesPosition, error) {
	var result FuturesPosition
	_, response, err := c.client.Request("GET", GetInstrumentIdUri(FUTURES_POSITION, instrumentId), nil, nil)
	if err != nil {
		return FuturesPosition{}, err
	}
	result, err = parsePositions(response, err)
	return result, err
}

// 单个币种合约账户信息
func (c *FutureClient) GetFutureAccountByUnderlying(underlying string) (FutureAccount, error) {
	var result FutureAccount
	_, _, err := c.client.Request("GET", GetUnderlyingUri(FUTURES_ACCOUNT, underlying), nil, &result)
	return result, err
}

// 获取合约币种杠杆倍数
func (c *FutureClient) GetFutureLeverageByUnderlying(underlying string) (FutureLeverage, error) {
	var result FutureLeverage
	_, response, err := c.client.Request("GET", GetUnderlyingUri(FUTURES_LEVERAGE, underlying), nil, nil)
	if err != nil {
		return FutureLeverage{}, err
	}
	//fmt.Println(response.Header.Get("resultDataJsonString"))
	result, err = c.parseFutureLeverage(response, err, underlying)
	return result, err
}

// 设定合约账户模式
func (c *FutureClient) PostFuturesAccountsMarginNode(underlying string, marginMode string) (PostMarginModeResult, error) {
	params := make(map[string]string)
	params["underlying"] = underlying
	params["margin_mode"] = marginMode
	var r PostMarginModeResult
	_, _, err := c.client.Request("POST", FUTURES_POST_MARGIN_MODE, params, &r)
	return r, err
}

// 设定合约杠杆
func (c *FutureClient) PostFutureLeverageByUnderlying(currency string, leverage int, optionalParams map[string]string) error {
	uri := GetUnderlyingUri(FUTURES_LEVERAGE, currency)
	params := make(map[string]string)
	params["leverage"] = strconv.Itoa(leverage)

	if optionalParams != nil && len(optionalParams) > 0 {
		params["instrument_id"] = optionalParams["instrument_id"]
		params["direction"] = optionalParams["direction"]
	}

	r := new(map[string]interface{})
	_, _, err := c.client.Request("POST", uri, params, r)
	return err
}

func parsePositions(response *http.Response, err error) (FuturesPosition, error) {
	var position FuturesPosition
	if err != nil {
		return position, err
	}
	var result Result
	result.Result = false
	jsonString := GetResponseDataJsonString(response)
	if strings.Contains(jsonString, "\"margin_mode\":\"fixed\"") {
		var fixedPosition FuturesFixedPosition
		err = JsonString2Struct(jsonString, &fixedPosition)
		if err != nil {
			return position, err
		} else {
			position.Result = fixedPosition.Result
			position.MarginMode = fixedPosition.MarginMode
			position.FixedPosition = fixedPosition.FixedPosition
		}
	} else if strings.Contains(jsonString, "\"margin_mode\":\"crossed\"") {
		var crossPosition FuturesCrossPosition
		err = JsonString2Struct(jsonString, &crossPosition)
		if err != nil {
			return position, err
		} else {
			position.Result = crossPosition.Result
			position.MarginMode = crossPosition.MarginMode
			position.CrossPosition = crossPosition.CrossPosition
		}
	} else if strings.Contains(jsonString, "\"code\":") {
		JsonString2Struct(jsonString, &position)
		position.Result = result
	} else {
		position.Result = result
	}

	return position, nil
}

func (c *FutureClient) parseFutureLeverage(response *http.Response, err error, underlying string) (FutureLeverage, error) {
	var leverageInfo FutureLeverage
	if err != nil {
		return leverageInfo, err
	}
	jsonString := GetResponseDataJsonString(response)
	//fmt.Println(jsonString)
	if strings.Contains(jsonString, "\"margin_mode\":\"fixed\"") {
		currencies := strings.Split(underlying, "-")
		instruments, err := c.GetFutureInstrumentByCurrency(currencies[0], currencies[1])
		if err != nil {
			return leverageInfo, err
		}
		//fmt.Printf("%#v", instruments)
		var m map[string]interface{}
		err = json.Unmarshal([]byte(jsonString), &m)
		if err != nil {
			fmt.Println("grwhrew: ", err.Error())
			return leverageInfo, err
		}
		leverageInfo.MarginMode = m["margin_mode"].(string)
		var data []FixedLeverageItem
		for _, v := range instruments {
				if v1, ok := m[v.InstrumentId]; ok {
					v2 := v1.(map[string]interface{})
					//fmt.Printf(fmt.Sprintf("%v", v3))
					var item FixedLeverageItem
					item.LongLeverage, _ = strconv.ParseFloat(v2["long_leverage"].(string), 64)
					item.ShortLeverage, _ = strconv.ParseFloat(v2["short_leverage"].(string), 64)
					item.InstrumentId = v.InstrumentId
					//fmt.Printf("%#v \n", item)
					data = append(data, item)
				}
		}
		leverageInfo.FixedLeverages = data
	} else if strings.Contains(jsonString, "\"margin_mode\":\"crossed\"") {
		var crossLeverage FuturesCrossLeverage
		err = JsonString2Struct(jsonString, &crossLeverage)
		if err != nil {
			return leverageInfo, err
		} else {
			leverageInfo.MarginMode = crossLeverage.MarginMode
			leverageInfo.CrossLeverage.Leverage =crossLeverage.Leverage
		}
	} else if strings.Contains(jsonString, "\"code\":") {
		JsonString2Struct(jsonString, &leverageInfo)
		return leverageInfo, errors.New(jsonString)
	}

	return leverageInfo, nil
}

// 市价全平
// 2次/2s （根据underlying，分别限速）
func (c *FutureClient) FutureCrossPosition(instrumentId string, direction string) ([]byte, CrossPositionResult, error) {
	var futureCrossPositionResult CrossPositionResult
	params := make(map[string]string)
	params["instrument_id"] = instrumentId
	params["direction"] = direction
	var respBody []byte
	respBody, _, err := c.client.Request("POST", FURURES_CROSS_POSITION, params, &futureCrossPositionResult)
	return respBody, futureCrossPositionResult, err
}

// 获取历史结算/交割记录
// 1次/ 60s
func (c *FutureClient) GetFutureSettlementHistory(instrumentId string, start, end time.Time, limit int) ([]SettlementItem, error) {
	var result []SettlementItem
	params := make(map[string]string)
	params["instrument_id"] = instrumentId
	params["limit"] = strconv.Itoa(limit)
	params["start"] = fmt.Sprintf("%v",start.Add(8 * time.Hour).UTC().Format(time.RFC3339))
	params["end"] = fmt.Sprintf("%v", end.Add(8 * time.Hour).UTC().Format(time.RFC3339))
	requestPath := BuildParams(FUTURES_SETTLEMENT_HISTORY, params)
	_, _, err := c.client.Request("GET", requestPath, nil, &result)
	return result, err
}

// 合约下单
//限速规则：60次/2s （1）不同合约之间限速不累计；2）同一合约的当周次周季度之间限速累计；3）同一合约的币本位和USDT保证金之间限速不累计）
// type 1:开多 | 2:开空 | 3:平多 | 4:平空
// order_type 0: 普通委托（order type不填或填0都是普通委托）| 1: 只做Maker（Post only）| 2: 全部成交或立即取消（FOK）| 3: 立即成交并取消剩余（IOC）| 4: 市价委托
// match_price  是否以对手价下单(0:不是; 1:是)，默认为0，当取值为1时，price字段无效。当以对手价下单，order_type只能选择0（普通委托）
func (c *FutureClient) PlaceFutureOrder(instrumentId string, pType int, orderType int, price float64, size float64, matchPrice int, clientOid string) ([]byte, FutureNewOrderResult, error) {
	var newOrderResult FutureNewOrderResult
	var params FutureNewOrderParams
	params.InstrumentId = instrumentId
	params.ClientOid = clientOid
	params.Price = fmt.Sprintf("%v", price)
	params.Size = fmt.Sprintf("%v", size)
	params.OrderType = fmt.Sprintf("%v", orderType)
	params.Type = fmt.Sprintf("%v", pType)
	params.MatchPrice = fmt.Sprintf("%v", matchPrice)
	var respBody []byte
	respBody, _, err := c.client.Request("POST", FUTURES_ORDER, params, &newOrderResult)
	return respBody, newOrderResult, err
}

// 撤单
// id 可以为client_id 也可以是orderID
func (c *FutureClient) CancelFuturesOrder(instrumentId string, id string) ([]byte, CancelResult, error) {
	var result CancelResult
	respBody, _, err := c.client.Request("POST",GetInstrumentOrderIdUri(FUTURES_CANCEL,instrumentId,id), nil, &result)
	return respBody, result, err
}

