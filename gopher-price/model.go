package main

import "time"

type BinanceResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type KuCoinResponse struct {
	Data struct {
		Price string `json:"price"`
	} `json:"data"`
}

type BybitResponse struct {
	Result []struct {
		Price string `json:"lastPrice"`
	} `json:"result"`
}

//Assign task for worker
type PriceJob struct {
	Exchange string
	URL      string
}

//Return
type PriceResult struct {
	Exchange string
	Price    string
	Duration time.Duration
	Err      error
}
