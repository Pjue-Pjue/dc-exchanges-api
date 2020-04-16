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

func NewFutureOKEXClient(apiKey, secretKey, passphrase string) *FutureClient {
	return &FutureClient{client: newOKExClient(apiKey, secretKey, passphrase)}
}

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
	requestPath := BuildParams(GetInstrumentIdUri(FUTURE_ORDERBOOK, instrumentId), params)
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

func (c *FutureClient) GetFutureInstrumentByCurrency(base string, quote string) ([]Instrument, error) {
	var instruments []Instrument
	_, _, err := c.client.Request("GET", FUTURE_INSTRUMENTS, nil, &instruments)
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

func (c *FutureClient) GetPositionByInstrumentId(instrumentId string) (FuturesPosition, error) {
	var result FuturesPosition
	_, response, err := c.client.Request("GET", GetInstrumentIdUri(FUTURE_POSITION, instrumentId), nil, nil)
	if err != nil {
		return FuturesPosition{}, err
	}
	result, err = parsePositions(response, err)
	return result, err
}

func (c *FutureClient) GetFutureAccountByUnderlying(underlying string) (FutureAccount, error) {
	var result FutureAccount
	_, _, err := c.client.Request("GET", GetUnderlyingUri(FUTURE_ACCOUNT, underlying), nil, &result)
	return result, err
}

func (c *FutureClient) GetFutureLeverageByUnderlying(underlying string) (FutureLeverage, error) {
	var result FutureLeverage
	_, response, err := c.client.Request("GET", GetUnderlyingUri(FUTURE_LEVERAGE, underlying), nil, nil)
	if err != nil {
		return FutureLeverage{}, err
	}
	//fmt.Println(response.Header.Get("resultDataJsonString"))
	result, err = c.parseFutureLeverage(response, err, underlying)
	return result, err
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