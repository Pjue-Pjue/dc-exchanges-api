package okex

type Account struct {
	Currency  string  `json:"currency"`
	Balance   float64 `json:"balance,string"`   // 余额
	Hold      float64 `json:"hold,string"`      // 冻结
	Available float64 `json:"available,string"` // 可用
}

type SpotMarketOrderParams struct {
	Notional     string `json:"notional"`
	BaseSpotOrderParams
}

type BaseSpotOrderParams struct {
	ClientOid    string `json:"client_oid"`
	InstrumentID string `json:"instrument_id"`
	OrderType    string `json:"order_type"`
	Side         string `json:"side"`
	Size         string `json:"size"`
	Type         string `json:"type"`
}

type SpotLimitOrderParams struct {
	Price        string `json:"price"`
	BaseSpotOrderParams
}

type SpotCommonResult struct {
	ClientOid    string `json:"client_oid"`
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	OrderID      string `json:"order_id"`
	Result       bool   `json:"result"`
}