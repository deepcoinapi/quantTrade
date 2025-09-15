package main

import (
	"quantTrade/core/data"
	"quantTrade/core/strategy/indicator"
)

func main() {
	strat := &indicator.BollingerBandsStrategy{}
	strat.Init(map[string]interface{}{
		"symbol": "BTC/USDT",
		"period": 20,
		"k":      2.0,
	})

	bars := []data.Bar{
		{Symbol: "BTC/USDT", Close: 20000},
		{Symbol: "BTC/USDT", Close: 20100},
		{Symbol: "BTC/USDT", Close: 20250},
		{Symbol: "BTC/USDT", Close: 20500},
		{Symbol: "BTC/USDT", Close: 21000}, // 可能突破
		{Symbol: "BTC/USDT", Close: 20800}, // 可能平仓
		{Symbol: "BTC/USDT", Close: 19900}, // 可能下轨突破
		{Symbol: "BTC/USDT", Close: 20100}, // 平空
	}

	for _, bar := range bars {
		signals := strat.OnBar(bar)
		for _, sig := range signals {
			println(sig.Action, sig.Symbol, sig.Price, sig.Size)
		}
	}
}
