package okex

import (
	dc_exchanges_api "dc-exchanges-api"
	"strconv"
)

type AccountClient struct {
	client *dc_exchanges_api.Client
}

func NewAccountOKEXClient(apiKey, secretKey, passphrase string) *AccountClient {
	return &AccountClient{client: newOKExClient(apiKey, secretKey, passphrase)}
}

func (c *AccountClient) GetAssetValuation(accountType int, valuationCurrency string) (AssetValuation, error) {
	r := AssetValuation{}
	params := map[string]string{}
	if accountType >= 0 {
		params["account_type"] = strconv.Itoa(accountType)
	}
	if valuationCurrency != "" {
		params["valuation_currency"] = valuationCurrency
	}

	uri := BuildParams(TOTAL_ASSEST,params)
	if _, _, err := c.client.Request("GET", uri, nil, &r); err != nil {
		return AssetValuation{}, err
	}
	return r,nil
}