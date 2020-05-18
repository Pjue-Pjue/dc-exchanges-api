package okex

import dc_exchanges_api "dc-exchanges-api"

type MarginOKEXClient struct {
	client *dc_exchanges_api.Client
}

func NewMarginOKEXClient(apiKey, secretKey, passphrase string, endPoint string) *MarginOKEXClient {
	return &MarginOKEXClient{client: newOKExClient(apiKey, secretKey, passphrase, endPoint)}
}