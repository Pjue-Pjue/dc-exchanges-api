package okex

import "testing"

func TestAccountClient_GetAssetValuation(t *testing.T) {
	c := NewAccountOKEXClient("", "", "")
	res, err := c.GetAssetValuation(0, VALUATEUSD)
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Logf("%#v", res)
}