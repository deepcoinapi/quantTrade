package arbitrage

import (
	"fmt"
	"quantTrade/core/data"
	"quantTrade/core/strategy"
)

// FutureSpotArbitrage 期现套利策略
type FutureSpotArbitrage struct {
	symbol       string
	threshold    float64 // 套利门槛，比如 0.5% 以上才操作
	positionOpen bool
}

// Init 初始化参数
func (s *FutureSpotArbitrage) Init(params map[string]interface{}) error {
	if sym, ok := params["symbol"].(string); ok {
		s.symbol = sym
	} else {
		return fmt.Errorf("missing symbol param")
	}
	if th, ok := params["threshold"].(float64); ok {
		s.threshold = th
	} else {
		s.threshold = 0.005 // 默认 0.5%
	}
	s.positionOpen = false
	return nil
}

// OnTick 处理行情数据
func (s *FutureSpotArbitrage) OnTick(tick data.Tick) []strategy.Signal {
	signals := []strategy.Signal{}

	spread := (tick.FuturePrice - tick.SpotPrice) / tick.SpotPrice

	if !s.positionOpen {
		if spread > s.threshold {
			// 期货溢价，做多现货 + 做空期货
			signals = append(signals, strategy.Signal{
				Action: "BUY_SPOT",
				Symbol: s.symbol,
				Price:  tick.SpotPrice,
				Size:   1.0,
			})
			signals = append(signals, strategy.Signal{
				Action: "SELL_SWAP",
				Symbol: s.symbol,
				Price:  tick.FuturePrice,
				Size:   1.0,
			})
			s.positionOpen = true
			fmt.Printf("[Arb] 开仓: 现货买入 %.2f, 期货卖出 %.2f, 差价 %.2f%%\n",
				tick.SpotPrice, tick.FuturePrice, spread*100)
		}
	} else {
		// 平仓条件：套利空间消失或变为负
		if spread < 0.001 {
			signals = append(signals, strategy.Signal{
				Action: "SELL_SPOT",
				Symbol: s.symbol,
				Price:  tick.SpotPrice,
				Size:   1.0,
			})
			signals = append(signals, strategy.Signal{
				Action: "SELL_SWAP",
				Symbol: s.symbol,
				Price:  tick.FuturePrice,
				Size:   1.0,
			})
			s.positionOpen = false
			fmt.Printf("[Arb] 平仓: 现货卖出 %.2f, 期货买入 %.2f, 差价 %.2f%%\n",
				tick.SpotPrice, tick.FuturePrice, spread*100)
		}
	}

	return signals
}

func (s *FutureSpotArbitrage) OnBar(bar data.Bar) []strategy.Signal {
	return nil
}

func (s *FutureSpotArbitrage) Name() string {
	return "FutureSpotArbitrage"
}
