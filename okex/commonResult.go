package okex

import "time"

type OrderBookResult struct {
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
