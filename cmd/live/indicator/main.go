package main

import (
	"context"
	"fmt"
	"quantTrade/core/execution"
	"quantTrade/core/execution/dc"
	"quantTrade/core/strategy/indicator"
	"strings"
)

func main() {
	cli := execution.NewDcClient()

	symbol := "BTC-USDT"
	strat := &indicator.BollingerBandsStrategy{}
	parms := map[string]interface{}{
		"symbol": symbol,
		"period": "1m",
		"n":      20,
		"k":      2,
	}
	d, err := cli.GetKlines("1m", symbol, "20")
	if err != nil {
		fmt.Println("get klines err:", err)
		return
	}
	if len(d) == 0 {
		fmt.Println("get 0 kline ?????")
		return
	}
	fmt.Println("klines:", d)
	parms["klines"] = d

	strat.Init(parms)

	//k 线订阅
	filterSpotValue := "DeepCoin_" + strings.ReplaceAll(symbol, "-", "") + "_1m"
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
			TopicID:     "11",
			ResumeNo:    -1,
		},
	}

	go execution.RunPublicWS(context.TODO(), dc.WS_SWAP_ADDR, sub, cli)

	for {
		bar := cli.GetNewBar(symbol)
		signals := strat.OnBar(bar)
		//continue
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
		//time.Sleep(5 * time.Second) // 每 5 秒轮询
	}
}
