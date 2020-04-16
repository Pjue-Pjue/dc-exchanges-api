package okex

import (
	dc_exchanges_api "dc-exchanges-api"
	"strconv"
)

//type OkClient struct {
//	SpotClient    *SpotClient
//	MarginClient  *MarginClient
//	FutureClient  *FutureClient
//	SwapClient    *SwapClient
//	AccountClient *AccountClient
//}


func newOKExClient(apiKey, secretKey, passphrase string) *dc_exchanges_api.Client {
	var config dc_exchanges_api.Config
	//config.Endpoint = "https://www.okex.com/"
	config.Endpoint = "https://www.okex.me/"
	config.ApiKey = apiKey
	config.SecretKey = secretKey
	config.Passphrase = passphrase
	config.TimeoutSecond = 45
	config.IsPrint = false
	config.I18n = "en_US"
	config.ProxyURL = "" //cfg.ProxyURL

	client := dc_exchanges_api.NewClient(config)
	return client
}

func ParseFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func ParseInt(s string) int {
	i, _ := strconv.ParseInt(s, 10, 64)
	return int(i)
}