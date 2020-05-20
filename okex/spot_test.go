package okex

import "testing"

func TestSpotOKEXClient_GetSpotOrderBook(t *testing.T) {
	c := NewSpotOKEXClient("", "", "","https://www.okex.me/")
	params := map[string]string{}
	params["size"] = "10"
	params["depth"] = "0.1"
	res, err := c.GetSpotOrderBook(10, 0.1,"BTC-USDT")
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Logf("%#v", res)
}
