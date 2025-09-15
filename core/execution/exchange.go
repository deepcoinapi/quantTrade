package execution

import "quantTrade/core/data"

//type OrderBook struct {
//}

type Order struct {
	ID     string
	Symbol string
	Side   string
	Price  float64
	Size   float64
	Status string
}

type Exchange interface {
	UpdatePrice(ts string, price float64)
	UpdateBar(bar data.Bar)
	GetTicker(symbol string) (tick data.Tick, err error)
	GetNewBar(symbol string) data.Bar
	PlaceOrder(symbol string, action, otype string, price float64, size float64, isfuture bool) (Order, error)
	CancelOrder(orderID, symbol string) error
	//GetBalance(asset string) (float64, error)
	//GetOrderBook(symbol string) (OrderBook, error)
}
