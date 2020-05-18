package okex

import dc_exchanges_api "dc-exchanges-api"

type SwapOKEXClient struct {
	client *dc_exchanges_api.Client
}

func NewSwapOKEXClient(apiKey, secretKey, passphrase string, endPoint string) *SwapOKEXClient {
	return &SwapOKEXClient{client: newOKExClient(apiKey, secretKey, passphrase, endPoint)}
}

func (c *SwapOKEXClient) GerSwapOrderBook(instrumentId string, options map[string]string) () {

}