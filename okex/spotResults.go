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

type SpotOrderResult struct {
	ClientOid      string `json:"client_oid"`
	CreatedAt      string `json:"created_at"`
	FilledNotional string `json:"filled_notional"`
	FilledSize     string `json:"filled_size"`
	Funds          string `json:"funds"`
	InstrumentID   string `json:"instrument_id"`
	Notional       string `json:"notional"`
	OrderID        string `json:"order_id"`
	OrderType      string `json:"order_type"`
	Price          string `json:"price"`
	PriceAvg       string `json:"price_avg"`
	ProductID      string `json:"product_id"`
	Side           string `json:"side"`
	Size           string `json:"size"`
	State          string `json:"state"`
	Status         string `json:"status"`
	Timestamp      string `json:"timestamp"`
	Type           string `json:"type"`
}

type SpotFillItem struct {
	CreatedAt    string `json:"created_at"`
	Currency     string `json:"currency"`
	ExecType     string `json:"exec_type"`
	Fee          string `json:"fee"`
	InstrumentID string `json:"instrument_id"`
	LedgerID     string `json:"ledger_id"`
	Liquidity    string `json:"liquidity"`
	OrderID      string `json:"order_id"`
	Price        string `json:"price"`
	ProductID    string `json:"product_id"`
	Side         string `json:"side"`
	Size         string `json:"size"`
	Timestamp    string `json:"timestamp"`
}

type SpotInstrumentsResult struct {
	BaseCurrency  string `json:"base_currency"` // 交易货币币种
	InstrumentID  string `json:"instrument_id"` //币对名称
	MinSize       string `json:"min_size"` //最小交易数量
	QuoteCurrency string `json:"quote_currency"` // 计价货币币种
	SizeIncrement string `json:"size_increment"` // 交易货币数量精度
	TickSize      string `json:"tick_size"` // 交易价格精度
}