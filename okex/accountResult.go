package okex

type AssetValuation struct {
	AccountType       string  `json:"account_type"`
	Balance           float64 `json:"balance"`
	Timestamp         string  `json:"timestamp"`
	ValuationCurrency string  `json:"valuation_currency"`
}

type TransferResult struct {
	Amount     string `json:"amount"`
	Currency   string `json:"currency"`
	From       string `json:"from"`
	Result     bool   `json:"result"`
	To         string `json:"to"`
	TransferID string `json:"transfer_id"`
}