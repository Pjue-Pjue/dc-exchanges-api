package okex

import dc_exchanges_api "dc-exchanges-api"

type SpotOKEXClient struct {
	client *dc_exchanges_api.Client
}

func NewSpotOKEXClient(apiKey, secretKey, passphrase string, endPoint string) *SpotOKEXClient {
	return &SpotOKEXClient{client: newOKExClient(apiKey, secretKey, passphrase, endPoint)}
}

func (c *SpotOKEXClient) GetSpotAccounts() (Accounts, error) {
	var accounts Accounts
	// todo
	return accounts, nil
}
