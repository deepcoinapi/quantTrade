package indicator

import (
	"fmt"
	"quantTrade/core/data"
	"quantTrade/core/strategy"
)

// BollingerBandsStrategy 布林带突破策略
type BollingerBandsStrategy struct {
	symbol   string
	period   string  // k线周期    1m,15m,1H,4H,1D,1W
	n        int     // MA 周期
	k        float64 // 标准差倍数
	position string  // 当前持仓状态: "LONG" / "SHORT" / ""
	prices   []data.Bar
}

// Init 初始化策略
func (s *BollingerBandsStrategy) Init(params map[string]interface{}) error {
	if sym, ok := params["symbol"].(string); ok {
		s.symbol = sym
	} else {
		return fmt.Errorf("missing symbol param")
	}

	if p, ok := params["n"].(int); ok {
		s.n = p
	} else {
		s.n = 20
	}

	if k, ok := params["k"].(float64); ok {
		s.k = k
	} else {
		s.k = 2.0
	}

	if k, ok := params["period"].(string); ok {
		s.period = k
	} else {
		s.period = "1m"
	}

	if k, ok := params["klines"].([]data.Bar); ok {
		s.prices = k
	} else {
		s.prices = []data.Bar{}
	}

	if k, ok := params["position"].(string); ok {
		s.position = k
	} else {
		s.position = ""
	}
	return nil
}

// OnBar 使用 K 线数据驱动策略
func (s *BollingerBandsStrategy) OnBar(bar data.Bar) []strategy.Signal {
	signals := []strategy.Signal{}

	//fmt.Println("1 bar:", bar)
	// 维护内存
	l := len(s.prices)
	if d := bar.Ts - s.prices[l-1].Ts; d == 0 {
		//同一分钟
		//fmt.Println("3 bar:", bar, s.prices[l-1].Ts)
		s.prices[l-1] = bar
	} else if d == 60 {
		s.prices = append(s.prices, bar)
	}

	l = len(s.prices)
	if l < s.n {
		//fmt.Println("4 bar:", bar)
		return signals // 数据不足，暂不生成信号
	}

	s.prices = s.prices[l-s.n:]
	// 取最后 N 个价格
	//window := s.prices[len(s.prices)-s.n:]

	// 计算均值
	//sum := 0.0
	//for _, p := range window {
	//	sum += p
	//}
	//ma := sum / float64(s.n)
	//
	//// 计算标准差
	//variance := 0.0
	//for _, p := range window {
	//	variance += (p - ma) * (p - ma)
	//}
	//std := math.Sqrt(variance / float64(s.n))

	ma, std := calcSMAStd(s.prices)

	upper := ma + s.k*std
	lower := ma - s.k*std

	price := bar.Close

	fmt.Println("up:", upper, "low:", lower, "md:", ma)

	if s.position == "" {
		if price > upper {
			signals = append(signals, strategy.Signal{
				Action: "BUY_SWAP",
				Symbol: s.symbol,
				Price:  price,
				Size:   1.0,
			})
			s.position = "LONG"
			fmt.Printf("[Bollinger] 突破上轨 %.2f, 买入 %.2f\n", upper, price)
		} else if price < lower {
			signals = append(signals, strategy.Signal{
				Action: "SELL_SWAP",
				Symbol: s.symbol,
				Price:  price,
				Size:   1.0,
			})
			s.position = "SHORT"
			fmt.Printf("[Bollinger] 跌破下轨 %.2f, 卖出 %.2f\n", lower, price)
		}
	} else if s.position == "LONG" {
		if price <= ma {
			signals = append(signals, strategy.Signal{
				Action: "SELL_SWAP",
				Symbol: s.symbol,
				Price:  price,
				Size:   1.0,
			})
			s.position = ""
			fmt.Printf("[Bollinger] 回到中轨 %.2f, 平多 %.2f\n", ma, price)
		}
	} else if s.position == "SHORT" {
		if price >= ma {
			signals = append(signals, strategy.Signal{
				Action: "BUY_SWAP",
				Symbol: s.symbol,
				Price:  price,
				Size:   1.0,
			})
			s.position = ""
			fmt.Printf("[Bollinger] 回到中轨 %.2f, 平空 %.2f\n", ma, price)
		}
	}

	return signals
}

func (s *BollingerBandsStrategy) OnTick(tick data.Tick) []strategy.Signal {
	return nil
}

func (s *BollingerBandsStrategy) Name() string {
	return "BollingerBandsStrategy"
}
