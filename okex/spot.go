package okex

import dc_exchanges_api "dc-exchanges-api"

type SpotClient struct {
	client *dc_exchanges_api.Client
}

func (c *SpotClient) GetSpotAccounts() (Accounts, error) {
	var accounts Accounts
	// todo
	return accounts, nil
}
