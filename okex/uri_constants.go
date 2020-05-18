package okex

const (
	VALUATEBTC = "BTC"
	VALUATEUSD = "USD"
	VALUATECNY = "CNY"
	VALUATEJPY = "JPY"
	VALUATEKRW = "KRW"
	VALUATERUB = "RUB"
	// 资金账户
	TOTAL_ASSEST = "/api/account/v3/asset-valuation"
	// 币币
	// 杠杆
	// 交割
	FUTURES_ORDERBOOK = "/api/futures/v3/instruments/{instrument_id}/book"
	FUTURES_INSTRUMENTS = "/api/futures/v3/instruments"
	FUTURES_POSITION = "/api/futures/v3/{instrument_id}/position"
	FUTURES_ACCOUNT = "/api/futures/v3/accounts/{underlying}"
	FUTURES_LEVERAGE = "/api/futures/v3/accounts/{underlying}/leverage"
	FURURES_CROSS_POSITION = "/api/futures/v3/close_position"
	FUTURES_SETTLEMENT_HISTORY = "/api/futures/v3/settlement/history"
	FUTURES_POST_MARGIN_MODE = "/api/futures/v3/accounts/margin_mode"
	FUTURES_ORDER = "/api/futures/v3/order"
	FUTURES_CANCEL = "/api/futures/v3/cancel_order/{instrument_id}/{order_id}"
	FUTURES_ESTIMATED_PRICE = "/api/futures/v3/instruments/{instrument_id}/estimated_price"
	RATE = "/api/futures/v3/rate"
	INDEX = "/api/futures/v3/instruments/{instrument_id}/index"
	// 永续
)
