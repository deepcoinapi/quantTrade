package execution

import (
	"encoding/json"
	"errors"
	"fmt"
	"quantTrade/core/data"
	"quantTrade/core/execution/dc"
	"strconv"
	"strings"
	"time"
)

func wsDataHandler(msg *DcResponseWSMsg, client Exchange) {
	switch msg.Action {
	case "PushMarketTrade":
		if ss := strings.Split(msg.Index, "_"); len(ss) > 1 {
			var r []MarketTrade
			if err := json.Unmarshal(msg.Result, &r); err != nil {
				fmt.Println("Unmarshal err:", err)
			} else {
				if len(r) > 0 {
					p := r[0].Data.Price
					if strings.Contains(ss[1], "/") {
						client.UpdatePrice("SPOT", p)
					} else {
						client.UpdatePrice("SWAP", p)
					}
				}
			}
		}
	}
}

type DcClient struct {
	Sign *dc.Sign
	Tick data.Tick
}

func NewDcClient() *DcClient {
	sign := dc.NewSign()
	return &DcClient{
		Sign: sign,
	}
}

func (o *DcClient) UpdatePrice(ts string, price float64) {
	switch ts {
	case "SPOT":
		o.Tick.SpotPrice = price
	case "SWAP":
		o.Tick.FuturePrice = price
	}
}

func (o *DcClient) GetTicker(symbol string) (data.Tick, error) {
	tick := o.Tick
	if tick.SpotPrice <= 0 || tick.FuturePrice <= 0 {
		return tick, errors.New("price is 0")
	}
	tick.Symbol = symbol
	return o.Tick, nil
}

// 下单（简化，只示例市价单）
func (o *DcClient) PlaceOrder(symbol string, action, otype string, px float64, sz float64, isfuture bool) (Order, error) {
	id := fmt.Sprintf("Dc%d", time.Now().UnixNano())

	//action SELL_SWAP
	infos := strings.Split(action, "_")
	if len(infos) < 1 {
		return Order{}, errors.New("action:" + action + " is wrong")
	}
	req := dc.OrderRequest{
		InstId:      symbol,
		TdMode:      dc.CROSS,
		ClOrdId:     id,
		OrdType:     otype,
		Px:          strconv.FormatFloat(px, 'f', 1, 64),
		Sz:          strconv.FormatFloat(sz, 'f', 2, 64),
		MrgPosition: dc.MERGE,
	}

	side := strings.ToLower(infos[0])
	req.Side = side
	if isfuture {
		pside := dc.POSITION_SIDE_LONG
		if req.Side != dc.SIDE_BUY {
			pside = dc.POSITION_SIDE_SHORT
		}
		req.PosSide = pside
		req.InstId = symbol + "-SWAP"
		//req.OrdType = dc.ORDER_TYPE_POST_ONLY
	} else {
		req.TdMode = dc.CASH
		req.TgtCcy = "base_ccy"
		req.Ccy = "USDT"
	}

	fmt.Printf("post order:%+v \n", req)
	if err := dc.Order(req, o.Sign); err != nil {
		return Order{}, err
	}
	return Order{
		ID:     id,
		Symbol: symbol,
		Side:   side,
		Price:  px,
		Size:   sz,
		Status: "placed",
	}, nil
}

func (o *DcClient) CancelOrder(orderID, symbol string) error {
	cancel := dc.CancelOrderRequest{
		InstId:  symbol,
		ClOrdId: orderID,
	}
	fmt.Printf("Cancel order:%+v \n", cancel)
	if err := dc.CancelOrder(cancel, o.Sign); err != nil {
		return err
	}
	return nil
}
