package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type PriceParser interface {
	Parse(data []byte) (string, error)
}
type BinanceHandler struct{}

func (h BinanceHandler) Parse(data []byte) (string, error) {
	var resp struct {
		Price string `json:"price"`
	}
	err := json.Unmarshal(data, &resp)
	return resp.Price, err
}

type KuCoinHandler struct{}

func (h KuCoinHandler) Parse(data []byte) (string, error) {
	var resp struct {
		Data struct {
			Price string `json:"price"`
		} `json:"data"`
	}
	err := json.Unmarshal(data, &resp)
	return resp.Data.Price, err
}

type BybitHandler struct{}

func (h BybitHandler) Parse(data []byte) (string, error) {
	var resp struct {
		Result struct {
			List []struct {
				LastPrice string `json:"lastPrice"`
			} `json:"list"`
		} `json:"result"`
	}
	err := json.Unmarshal(data, &resp)
	if err != nil {
		return "", err
	}
	if len(resp.Result.List) > 0 {
		return resp.Result.List[0].LastPrice, err
	}
	return "", fmt.Errorf("No result")

}

var Handlers = map[string]PriceParser{
	"Binance":     BinanceHandler{},
	"Binance-ETH": BinanceHandler{},
	"KuCoin":      KuCoinHandler{},
	"ByBit":       BybitHandler{},
}

// Assign task for worker
type PriceJob struct {
	Exchange string
	URL      string
}

// Return
type PriceResult struct {
	Exchange string
	Price    string
	Duration time.Duration
	Err      error
}
