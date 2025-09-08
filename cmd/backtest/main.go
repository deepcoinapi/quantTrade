package main

import (
	"context"
	"fmt"
	"quantTrade/core/execution"
	"quantTrade/core/execution/dc"
	"quantTrade/core/strategy/arbitrage"
	"strings"
	"time"
)

func main() {
	cli := execution.NewDcClient()

	symbol := "BTC-USDT"
	strat := &arbitrage.FutureSpotArbitrage{}
	strat.Init(map[string]interface{}{
		"symbol":    "BTC-USDT",
		"threshold": 0.005, // 0.5% 差价
	})

	filterSpotValue := "DeepCoin_" + strings.ReplaceAll(symbol, "-", "/")
	sub := execution.DcSubWSMsg{
		SendTopicAction: struct {
			Action      string
			FilterValue string
			LocalNo     int
			TopicID     string
			ResumeNo    int
		}{
			Action:      "1",
			FilterValue: filterSpotValue,
			LocalNo:     1,
			TopicID:     "25",
			ResumeNo:    -1,
		},
	}
	go execution.RunPublicWS(context.TODO(), dc.WS_SPOT_ADDR, sub, cli)

	filterSwapValue := "DeepCoin_" + strings.ReplaceAll(symbol, "-", "")
	sub.SendTopicAction.FilterValue = filterSwapValue
	go execution.RunPublicWS(context.TODO(), dc.WS_SWAP_ADDR, sub, cli)

	for {
		tick, err := cli.GetTicker(symbol)
		if err != nil {
			fmt.Println("获取行情失败:", err)
			time.Sleep(2 * time.Second)
			continue
		}

		signals := strat.OnTick(tick)
		for _, sig := range signals {
			isFuture := false
			if sig.Action == "SELL_SWAP" || sig.Action == "BUY_SWAP" {
				isFuture = true
			}
			order, err := cli.PlaceOrder(sig.Symbol, sig.Action, dc.ORDER_TYPE_MARKET, sig.Price, sig.Size, isFuture)
			if err != nil {
				fmt.Println("下单失败:", err)
			} else {
				fmt.Println("下单成功:", order)
			}
		}

		time.Sleep(5 * time.Second) // 每 5 秒轮询
	}
}
