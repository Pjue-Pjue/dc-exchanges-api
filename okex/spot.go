package okex

import (
	dc_exchanges_api "dc-exchanges-api"
	"fmt"
)

type SpotOKEXClient struct {
	client *dc_exchanges_api.Client
}

func NewSpotOKEXClient(apiKey, secretKey, passphrase string, endPoint string) *SpotOKEXClient {
	return &SpotOKEXClient{client: newOKExClient(apiKey, secretKey, passphrase, endPoint)}
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