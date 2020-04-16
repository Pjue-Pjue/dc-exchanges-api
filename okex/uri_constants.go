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
	FUTURE_ORDERBOOK = "/api/futures/v3/instruments/{instrument_id}/book"
	FUTURE_INSTRUMENTS = "/api/futures/v3/instruments"
	FUTURE_POSITION = "/api/futures/v3/{instrument_id}/position"
	FUTURE_ACCOUNT = "/api/futures/v3/accounts/{underlying}"
	FUTURE_LEVERAGE = "/api/futures/v3/accounts/{underlying}/leverage"
	// 永续
)
