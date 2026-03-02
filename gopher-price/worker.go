package main

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

func fetchPriceWorker(id int, jobs <-chan PriceJob, results chan<- PriceResult, wg *sync.WaitGroup) {
	defer wg.Done()

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	for job := range jobs {
		start := time.Now()
		resp, err := client.Get(job.URL)
		result := PriceResult{Exchange: job.Exchange, Err: err}
		if err == nil {
			// Xử lý JSON từ các exchange khác nhau
			if job.Exchange == "Binance" || job.Exchange == "Binance-ETH" {
				var data BinanceResponse
				decodeErr := json.NewDecoder(resp.Body).Decode(&data)
				if decodeErr == nil {
					result.Price = data.Price
				} else {
					result.Err = decodeErr
				}
			} else if job.Exchange == "KuCoin" {
				var data KuCoinResponse
				decodeErr := json.NewDecoder(resp.Body).Decode(&data)
				if decodeErr == nil {
					result.Price = data.Data.Price
				} else {
					result.Err = decodeErr
				}
			} else if job.Exchange == "Bybit" {
				var data BybitResponse
				decodeErr := json.NewDecoder(resp.Body).Decode(&data)
				if decodeErr == nil && len(data.Result) > 0 {
					result.Price = data.Result[0].Price
				} else {
					result.Err = decodeErr
				}
			}
		}
		result.Duration = time.Since(start)
		results <- result
	}

}
