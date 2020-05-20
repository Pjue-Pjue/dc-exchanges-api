package okex

import (
	dc_exchanges_api "dc-exchanges-api"
	"fmt"
	"sort"
	"time"
)

type SpotOKEXClient struct {
	client *dc_exchanges_api.Client
}

func NewSpotOKEXClient(apiKey, secretKey, passphrase string, endPoint string) *SpotOKEXClient {
	return &SpotOKEXClient{client: newOKExClient(apiKey, secretKey, passphrase, endPoint)}
}

func (c *SpotOKEXClient) GetSpotOrderBook(size int, depth float64, instrumentId string) (OrderBook, error) {
	var ob OrderBook
	var orderBook OrderBookResult
	params := map[string]string{}
	params["size"] = fmt.Sprintf("%v", size)   // "10"
	params["depth"] = fmt.Sprintf("%v", depth) // "0.1"

	requestPath := BuildParams(GetInstrumentIdUri(SPOT_ORDERBOOK, instrumentId), params)
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

// 币币 所有币对信息
func (c *SpotOKEXClient) GetSpotInstruments() ([]SpotInstrumentsResult, error) {
	var result []SpotInstrumentsResult
	_,_, err := c.client.Request("GET", SPOT_INSTRUMENTS, nil, &result)
	return result,err
}

func (c *SpotOKEXClient) GetSpotAccounts() ([]Account, error) {
	var accounts []Account
	_, _, err := c.client.Request("GET",SPOT_ACCOUNTS, nil, &accounts)
	return accounts, err
}

func (c *SpotOKEXClient) GetSpotAccountByCurrency(currency string) (Account, error) {
	var account Account
	_, _, err := c.client.Request("GET",GetCurrencyUri(SPOT_ACCOUNT_CURRENCY, currency), nil, &account)
	return account, err
}

// 市价单
// type limit-限价  market-市价
// side buy/sell
// order_type: [0-普通委托（order type不填或填0都是普通委托)] [1:只做Maker（Post only)] [2-全部成交或立即取消（FOK)] [3-立即成交并取消剩余（IOC)]
// size 卖出数量  notional- 买入金额
func (c *SpotOKEXClient) PlaceMarketSpotOrder(instrumentId,clientOid, side string,orderType int, size, notional float64) ([]byte, SpotCommonResult, error) {
	var result SpotCommonResult
	var params SpotMarketOrderParams
	params.Type = "market"
	params.OrderType = fmt.Sprintf("%v", orderType)
	params.Side = side
	params.ClientOid = clientOid
	params.InstrumentID = instrumentId
	params.Size = fmt.Sprintf("%v",size)
	params.Notional = fmt.Sprintf("%v",notional)
	respBody, _, err := c.client.Request("POST", SPOT_ORDER, params, &result)
	return respBody, result, err
}

// 限价单
func (c *SpotOKEXClient) PlaceLimitSpotOrder(instrumentId,clientOid, side string, orderType int, size, price float64) ([]byte, SpotCommonResult, error) {
	var result SpotCommonResult
	var params SpotLimitOrderParams
	params.Type = "limit"
	params.OrderType = fmt.Sprintf("%v", orderType)
	params.Side = side
	params.ClientOid = clientOid
	params.InstrumentID = instrumentId
	params.Size = fmt.Sprintf("%v",size)
	params.Price = fmt.Sprintf("%v",price)
	respBody, _, err := c.client.Request("POST", SPOT_ORDER, params, &result)
	return respBody, result, err
}

// id: client_oid 、 orderId 二选一
func (c *SpotOKEXClient) CancelOrder(instrumentId string, id string) ([]byte, SpotCommonResult, error) {
	var result SpotCommonResult
	params := struct {
		InstrumentId string `json:"instrument_id"`
	}{
		InstrumentId: instrumentId,
	}
	respBody, _, err := c.client.Request("POST", GetOrderIdUri(SPOT_CANCEL_ORDER, id), params, &result)
	return respBody, result, err
}

// id: client_oid 、 orderId 二选一
func (c *SpotOKEXClient) GetSpotOrderInfo(instrumentId string, id string) (SpotOrderResult, error) {
	var result SpotOrderResult
	params := make(map[string]string)
	params["instrument_id"] = instrumentId
	requestPath := BuildParams(GetOrderIdUri(SPOT_ORDER_INFO, id), params)
	_, _, err := c.client.Request("GET", requestPath, nil, &result)
	return result, err
}

// 成交明细
// orderId 不填表示当前币对所有的成交明细
func (c *SpotOKEXClient) GetSpotFills(orderId string, instrumentId string, limit int, before string, after string) ([]SpotFillItem, error) {
	var result []SpotFillItem
	params := make(map[string]string)
	params["instrument_id"] = instrumentId
	params["order_id"] = orderId
	params["limit"] = fmt.Sprintf("%v", limit)
	if before != "" {
		params["before"] = before
	}
	if after != "" {
		params["after"] = after
	}
	requestPath := BuildParams(SPOT_ORDER_FILLS, params)
	_, _, err := c.client.Request("GET", requestPath, nil, &result)
	return result, err
}