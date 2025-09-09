package execution

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

// 订阅消息
type DcSubWSMsg struct {
	SendTopicAction struct {
		Action      string
		FilterValue string
		LocalNo     int
		TopicID     string
		ResumeNo    int
	} `json:"SendTopicAction"`
}

type MarketOrder struct {
	Table string `json:"table"`
	Data  struct {
		ExchangeID   string  `json:"ExchangeID"`
		InstrumentID string  `json:"InstrumentID"`
		Direction    string  `json:"Direction"`
		Price        float64 `json:"Price"`
		Volume       float64 `json:"Volume"`
		Orders       int     `json:"Orders"`
	} `json:"data"`
}

type MarketTrade struct {
	Table string `json:"table"`
	Data  struct {
		TradeID      string  `json:"TradeID"`
		ExchangeID   string  `json:"ExchangeID"`
		InstrumentID string  `json:"InstrumentID"`
		Direction    string  `json:"Direction"`
		Price        float64 `json:"LastPrice"`
		Volume       float64 `json:"Volume"`
		TradeTime    int     `json:"TradeTime"`
	} `json:"data"`
}

// ob 结构体
type DcResponseWSMsg struct {
	Action     string          `json:"action"`
	ErrorMsg   string          `json:"errorMsg"`
	Index      string          `json:"index"`
	BNo        int64           `json:"bNo"`
	ChangeType string          `json:"changeType"`
	Result     json.RawMessage `json:"result"`
}

func RunPublicWS(ctx context.Context, url string, sub DcSubWSMsg, exc Exchange) error {
	//dialer := websocket.Dialer{HandshakeTimeout: 10 * time.Second}
	fmt.Println("url:", url)
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("ws dial: %w", err)
	}
	defer c.Close()

	//for _, sub := range subs {
	fmt.Println("sub:", sub)
	if err := c.WriteJSON(sub); err != nil {
		fmt.Println("sub:", sub, "err:", err)
		return err
	}
	//}

	_ = c.SetReadDeadline(time.Now().Add(30 * time.Second))
	c.SetPongHandler(func(appData string) error { _ = c.SetReadDeadline(time.Now().Add(30 * time.Second)); return nil })

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			_ = c.WriteControl(websocket.PingMessage, []byte("ping"), time.Now().Add(3*time.Second))
		default:
			c.SetReadDeadline(time.Now().Add(30 * time.Second))
			_, msg, err := c.ReadMessage()
			if err != nil {
				fmt.Println("err:", err)
				return err
			}
			//fmt.Println("msg:", string(msg))
			m := &DcResponseWSMsg{}
			if err := json.Unmarshal(msg, m); err != nil {
				fmt.Println("Unmarshal err:", err)
				continue
			}
			//fmt.Println("msg:", m.Action, m.ErrorMsg)
			if m.Action == "RecvTopicAction" && m.ErrorMsg != "Success" {
				return errors.New(m.ErrorMsg)
			}

			if len(m.Result) == 0 {
				continue
			}
			wsDataHandler(m, exc)
		}
	}
}
