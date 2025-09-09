package main

import (
	"quantTrade/core/data"
	"quantTrade/core/strategy/arbitrage"
)

func main() {
	strat := &arbitrage.FutureSpotArbitrage{}
	strat.Init(map[string]interface{}{
		"symbol":    "BTC-USDT",
		"threshold": 0.005,
	})

	ticks := []data.Tick{
		{Symbol: "BTC-USDT", SpotPrice: 20000, FuturePrice: 20150},
		{Symbol: "BTC-USDT", SpotPrice: 20050, FuturePrice: 20100},
		{Symbol: "BTC-USDT", SpotPrice: 20100, FuturePrice: 20105},
		{Symbol: "BTC-USDT", SpotPrice: 20200, FuturePrice: 20210},
		{Symbol: "BTC-USDT", SpotPrice: 20300, FuturePrice: 20300}, // 差价消失，平仓
	}

	for _, tick := range ticks {
		signals := strat.OnTick(tick)
		for _, sig := range signals {
			println(sig.Action, sig.Symbol, sig.Price, sig.Size)
		}
	}
}
