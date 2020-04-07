package okex

import (
	"net/url"
)

/*
 Get api requestPath + requestParams
	params := NewParams()
	params["depth"] = "200"
	params["conflated"] = "0"
	url := BuildParams("/api/futures/v3/products/BTC-USD-0310/book", params)
 return eg:/api/futures/v3/products/BTC-USD-0310/book?conflated=0&depth=200
*/
func BuildParams(requestPath string, params map[string]string) string {
	urlParams := url.Values{}
	for k := range params {
		urlParams.Add(k, params[k])
	}
	return requestPath + "?" + urlParams.Encode()
}
