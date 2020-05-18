package okex

import "time"

type FutureOrderbookResult struct {
	Asks      [][]string `json:"asks"`
	Bids      [][]string `json:"bids"`
	Timestamp string     `json:"timestamp"`
}

// price	String	价格
// size	String	数量
// num_orders	String	组成此条深度的订单数量

type BookItem struct {
	Price     float64 `json:"price"`
	Size      float64 `json:"size"`
	NumOrders int     `json:"num_orders"`
}

type OrderBook struct {
	Asks      []BookItem `json:"asks,string"`
	Bids      []BookItem `json:"bids,string"`
	Timestamp time.Time  `json:"timestamp"`
}

type BizWarmTips struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Msg     string `json:"msg"`
}

type Result struct {
	Result bool `json:"result"`
}

type PageResult struct {
	From  int
	To    int
	Limit int
}

type FuturesPosition struct {
	BizWarmTips
	Result
	MarginMode    string
	CrossPosition []FuturesCrossPositionHolding
	FixedPosition []FuturesFixedPositionHolding
}

type FuturesCrossPosition struct {
	Result
	MarginMode    string                        `json:"margin_mode"`
	CrossPosition []FuturesCrossPositionHolding `json:"holding"`
}

type FuturesFixedPosition struct {
	Result
	MarginMode    string                        `json:"margin_mode"`
	FixedPosition []FuturesFixedPositionHolding `json:"holding"`
}

type FuturesCrossPositionHolding struct {
	FuturesPositionBase
	LiquidationPrice float64 `json:"liquidation_price,string"`
	Leverage         float64 `json:"leverage,string"`
}

type FuturesFixedPositionHolding struct {
	FuturesPositionBase
	LongMargin      float64 `json:"long_margin,string"`
	LongLiquiPrice  float64 `json:"long_liqui_price,string"`
	LongPnlRatio    float64 `json:"long_pnl_ratio,string"`
	LongLeverage    float64 `json:"long_leverage,string"`
	ShortMargin     float64 `json:"short_margin,string"`
	ShortLiquiPrice float64 `json:"short_liqui_price,string"`
	ShortPnlRatio   float64 `json:"short_pnl_ratio,string"`
	ShortLeverage   float64 `json:"short_leverage,string"`
}

type FuturesPositionBase struct {
	LongQty              float64 `json:"long_qty,string"`
	LongAvailQty         float64 `json:"long_avail_qty,string"`
	LongAvgCost          float64 `json:"long_avg_cost,string"`
	LongSettlementPrice  float64 `json:"long_settlement_price,string"`
	RealizedPnl          float64 `json:"realized_pnl,string"`
	ShortQty             float64 `json:"short_qty,string"`
	ShortAvailQty        float64 `json:"short_avail_qty,string"`
	ShortAvgCost         float64 `json:"short_avg_cost,string"`
	ShortSettlementPrice float64 `json:"short_settlement_price,string"`
	InstrumentId         string  `json:"instrument_id"`
	CreatedAt            string  `json:"created_at"`
	UpdatedAt            string  `json:"updated_at"`
}

type FutureAccount struct {
	TotalAvailBalance float64                        `json:"total_avail_balance,string"` // 账户余额（账户静态权益）
	Contracts         []FuturesFixedAccountContracts `json:"contracts"`
	Equity            float64                        `json:"equity,string"`         // 账户权益（账户动态权益）
	MarginMode        string                         `json:"margin_mode"`           // 账户类型 全仓：crossed 逐仓: fixed
	AutoMargin        int                            `json:"auto_margin,string"`    // 是否自动追加保证金 1: 自动追加已开启 0: 自动追加未开启
	LiquiMode         string                         `json:"liqui_mode"`            // 强平模式：tier（梯度强平）
	CanWithdraw       float64                        `json:"can_withdraw,string"`   // 可划转数量
	RealizedPnl       float64                        `json:"realized_pnl,string"`   // 全仓模式 已实现盈亏
	UnRealizedPnl     float64                        `json:"unrealized_pnl,string"` // 全仓模式 未实现盈亏
	Margin            float64                        `json:"margin,string"`         // 保证金（挂单冻结+持仓已用）
}

type FuturesFixedAccountContracts struct {
	AvailableQty      float64 `json:"available_qty,string"`
	FixedBalance      float64 `json:"fixed_balance,string"`
	InstrumentId      string  `json:"instrument_id"`
	MarginFixed       float64 `json:"margin_fixed,string"`
	MarginForUnfilled float64 `json:"margin_for_unfilled,string"`
	MarginFrozen      float64 `json:"margin_frozen,string"`
	RealizedPnl       float64 `json:"realized_pnl,string"`
	UnrealizedPnl     float64 `json:"unrealizedPnl,string"`
}

type FutureLeverage struct {
	MarginMode     string              `json:"margin_mode"`
	CrossLeverage  CrossLeverageItem   `json:"cross_leverage"`
	FixedLeverages []FixedLeverageItem `json:"fixed_leverages"`
}

type CrossLeverageItem struct {
	Underlying string  `json:"underlying"`
	Leverage   float64 `json:"leverage,string"`
}

type FixedLeverageItem struct {
	InstrumentId  string  `json:"instrument_id"`
	LongLeverage  float64 `json:"long_leverage"`
	ShortLeverage float64 `json:"short_leverage"`
}

type FuturesCrossLeverage struct {
	CrossLeverageItem
	MarginMode string `json:"margin_mode"`
}

//type FuturesFixedLeverage struct {
//
//}

type Instrument struct {
	InstrumentId        string  `json:"instrument_id"`
	UnderlyingIndex     string  `json:"underlying_index"`
	QuoteCurrency       string  `json:"quote_currency"`
	TickSize            float64 `json:"tick_size,string"`
	ContractVal         float64 `json:"contract_val,string"`
	Listing             string  `json:"listing"`
	Delivery            string  `json:"delivery"`
	TradeIncrement      float64 `json:"trade_increment,string"`
	Alias               string  `json:"alias"`
	Underlying          string  `json:"underlying"`
	BaseCurrency        string  `json:"base_currency"`
	SettlementCurrency  string  `json:"settlement_currency"`
	IsInverse           bool    `json:"is_inverse,string"`
	ContractValCurrency string  `json:"contract_val_currency"`
}

type CrossPositionResult struct {
	Direction    string `json:"direction"`
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	InstrumentID string `json:"instrument_id"`
	Result       bool   `json:"result"`
}

type SettlementItem struct {
	ClawbackLoss    string `json:"clawback_loss"`
	ClawbackRate    string `json:"clawback_rate"`
	InstrumentID    string `json:"instrument_id"`
	Reserve         string `json:"reserve"`
	SettlementPrice string `json:"settlement_price"`
	Timestamp       string `json:"timestamp"`
	Type            string `json:"type"`
}

type PostMarginModeResult struct {
	MarginMode string `json:"margin_mode"`
	Result     bool   `json:"result"`
	Underlying string `json:"underlying"`
}

type FutureNewOrderParams struct {
	InstrumentId string `json:"instrument_id"`
	ClientOid    string `json:"client_oid"`
	Type         string `json:"type"`
	OrderType    string `json:"order_type"`
	Price        string `json:"price"`
	Size         string `json:"size"`
	MatchPrice   string `json:"match_price"`
}

type FutureNewOrderResult struct {
	ClientOid    string `json:"client_oid"`
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	OrderID      string `json:"order_id"`
	Result       bool   `json:"result"`
}

type CancelResult struct {
	InstrumentId string `json:"instrument_id"`
	FutureNewOrderResult
}

type EstimatedPriceResult struct {
	SettlementPrice string `json:"settlement_price"`
	BaseInstrumentItem
}

type RateResult struct {
	Rate         string `json:"rate"`
	BaseInstrumentItem
}

type IndexInfo struct {
	Index string `json:"index"`
	BaseInstrumentItem
}

type BaseInstrumentItem struct {
	InstrumentID    string `json:"instrument_id"`
	Timestamp       string `json:"timestamp"`
}