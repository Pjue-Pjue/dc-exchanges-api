package okex

import (
	"testing"
	"time"
)

func TestFutureClient_GetOrderBook(t *testing.T) {
	c := NewFutureOKEXClient("", "", "","https://www.okex.me/")
	params := map[string]string{}
	params["size"] = "10"
	params["depth"] = "0.1"
	res, err := c.GetOrderBook("BTC-USD-200424", params)
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Logf("%#v", res)
}

// todo test
func TestFutureClient_GetPositionByInstrumentId(t *testing.T) {
	c := NewFutureOKEXClient("", "", "","https://www.okex.me/")
	res, err := c.GetPositionByInstrumentId("BTC-USD-200424")
	if err != nil {
		t.Logf("err: %v", err.Error())
		return
	}
	t.Logf("%#v", res)
}

func TestFutureClient_GetFutureLeverageByUnderlying(t *testing.T) {
	c := NewFutureOKEXClient("", "", "","https://www.okex.me/")
	res, err := c.GetFutureLeverageByUnderlying("XRP-USD")
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Logf("%v", res)
}

func TestFutureClient_GetFutureInstrumentByCurrency(t *testing.T) {
	c := NewFutureOKEXClient("", "", "","https://www.okex.me/")
	res, err := c.GetFutureInstrumentByCurrency("BTC", "USD")
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(res)
}

func TestFutureClient_GetFutureSettlementHistory(t *testing.T) {
	c := NewFutureOKEXClient("", "", "","https://www.okex.me/")
	res, err := c.GetFutureSettlementHistory("BTC-USD-200626", time.Now().Add(-120 * time.Hour), time.Now(), 10)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(res)
}