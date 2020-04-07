package okex

type AssetValuation struct {
	AccountType       string  `json:"account_type"`
	Balance           float64 `json:"balance"`
	Timestamp         string  `json:"timestamp"`
	ValuationCurrency string  `json:"valuation_currency"`
}
