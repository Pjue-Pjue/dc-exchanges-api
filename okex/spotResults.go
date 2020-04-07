package okex

type Accounts []Account

type Account struct {
	Currency  string  `json:"currency"`
	Balance   float64 `json:"balance,string"`   // 余额
	Hold      float64 `json:"hold,string"`      // 冻结
	Available float64 `json:"available,string"` // 可用
}
