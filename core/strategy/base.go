package strategy

import "quantTrade/core/data"

type Strategy interface {
	Init(params map[string]interface{}) error
	OnTick(tick data.Tick) []Signal
	OnBar(bar data.Bar) []Signal
	Name() string
}

type Signal struct {
	Action string // "BUY" or "SELL"
	Symbol string // "BTC-USDT"
	Price  float64
	Size   float64
}
