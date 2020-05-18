package okex

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
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

func GetInstrumentIdUri(uri string, instrumentId string) string {
	return strings.Replace(uri, "{instrument_id}", instrumentId, -1)
}

func GetUnderlyingUri (uri string, underlying string) string {
	return strings.Replace(uri, "{underlying}", underlying, -1)
}

func GetResponseDataJsonString(response *http.Response) string {
	return response.Header.Get("resultDataJsonString")
}

func JsonString2Struct(jsonString string, result interface{}) error {
	jsonBytes := []byte(jsonString)
	err := json.Unmarshal(jsonBytes, result)
	return err
}
func GetInstrumentOrderIdUri(uri string, instrumentId string, id string) string {
	uri = strings.Replace(uri, "{instrument_id}", instrumentId, -1)
	uri = strings.Replace(uri, "{order_id}", id, -1)
	return uri
}

func GetCurrencyUri(uri string, currency string) string {
	return strings.Replace(uri, "{currency}", currency, -1)
}

func GetOrderIdUri(uri string, id string) string {
	return strings.Replace(uri, "{order_id}", id, -1)
}