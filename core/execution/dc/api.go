package dc

import (
	"encoding/json"
	"errors"
	"fmt"
)

type resp struct {
	Code string          `json:"code"`
	Msg  string          `json:"msg"`
	Data json.RawMessage `json:"data"`
}

type Position struct {
	InstType    string `json:"instType"`
	MgnMode     string `json:"mgnMode"`
	InstId      string `json:"instId"`
	PosId       string `json:"posId"`
	PosSide     string `json:"posSide"`
	Pos         string `json:"pos"`
	AvgPx       string `json:"avgPx"`
	Lever       string `json:"lever"`
	LiqPx       string `json:"liqPx"`
	UseMargin   string `json:"useMargin"`
	MrgPosition string `json:"mrgPosition"`
	Ccy         string `json:"ccy"`
	UTime       string `json:"uTime"`
	CTime       string `json:"cTime"`
}

/*
	{
	    "instId": "BTC-USDT-SWAP",
	    "tdMode": "cross",
	    "side": "buy",
	    "ordType": "market",
	    "sz": "5",
	    "posSide": "long",
	    "mrgPosition": "merge",
	}
*/
type OrderRequest struct {
	InstId      string `json:"instId"`
	TdMode      string `json:"tdMode"`
	Ccy         string `json:"ccy,omitempty"`
	ClOrdId     string `json:"clOrdId,omitempty"`
	Tag         string `json:"tag,omitempty"`
	Side        string `json:"side"`
	PosSide     string `json:"posSide,omitempty"`
	MrgPosition string `json:"mrgPosition,omitempty"`
	ClosePosId  string `json:"closePosId,omitempty"`
	OrdType     string `json:"ordType"`
	Sz          string `json:"sz"`
	Px          string `json:"px"`
	ReduceOnly  string `json:"reduceOnly,omitempty"`
	TgtCcy      string `json:"tgtCcy"`
	TpTriggerPx string `json:"tpTriggerPx,omitempty"`
	SlTriggerPx string `json:"slTriggerPx,omitempty"`
}

type CancelOrderRequest struct {
	InstId  string `json:"instId,omitempty"`
	OrdId   string `json:"ordId,omitempty"`
	ClOrdId string `json:"clOrdId,omitempty"`
}

func Positions(intID string, s *Sign) ([]Position, error) {
	requestURL := fmt.Sprintf(URL+TRADE_POSITION+"?instType=%s&instId=%s", SWAP, intID)
	requestPath := fmt.Sprintf(TRADE_POSITION+"?instType=%s&instId=%s", SWAP, intID)
	if body, err := DoHttp(requestURL, HTTP_METHOD_GET, requestPath, "", s); err != nil {
		return nil, err
	} else {
		var m resp
		if err := json.Unmarshal(body, &m); err != nil {
			return nil, err
		} else {
			var pos []Position
			if err := json.Unmarshal(m.Data, &pos); err != nil {
				return nil, err
			} else {
				return pos, nil
			}
		}
	}
}

func Order(o OrderRequest, s *Sign) error {
	requestBody, err := json.Marshal(o)
	if err != nil {
		fmt.Println(err)
		return err
	}

	requestURL := s.Url + TRADE_ORDER
	if body, err := DoHttp(requestURL, HTTP_METHOD_POST, TRADE_ORDER, string(requestBody), s); err != nil {
		return err
	} else {
		var m resp
		if err := json.Unmarshal(body, &m); err != nil {
			return err
		} else {
			if m.Code == "0" {
				return nil
			} else {
				return errors.New(m.Msg)
			}
		}
	}
}

func CancelOrder(cancelOrder CancelOrderRequest, s *Sign) error {
	requestBody, err := json.Marshal(cancelOrder)
	if err != nil {
		fmt.Println(err)
		return err
	}

	requestURL := s.Url + TRADE_CANCEL_ORDER
	if body, err := DoHttp(requestURL, HTTP_METHOD_POST, TRADE_CANCEL_ORDER, string(requestBody), s); err != nil {
		return err
	} else {
		var m resp
		if err := json.Unmarshal(body, &m); err != nil {
			return err
		} else {
			if m.Code == "0" {
				return nil
			} else {
				return errors.New(m.Msg)
			}
		}
	}
}
