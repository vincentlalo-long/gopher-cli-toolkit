package main

import (
	//"fmt"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

func main() {
	/*client := http.Client{
		Timeout: 2 * time.Second,
	}
	client2 := http.Client{
		Timeout: 2 * time.Second,
	}
	resp, _ := client.Get("https://api.binance.com/api/v3/ticker/price?symbol=BTCUSDT")
	resp2, _ := client2.Get("https://api.binance.com/api/v3/ticker/price?symbol=ETHUSDT")
	defer resp.Body.Close()
	defer resp2.Body.Close()
	body1, _ := io.ReadAll(resp.Body)
	body2, _ := io.ReadAll(resp2.Body)
	fmt.Println(string(body1))
	fmt.Println(string(body2))*/

	targets := []PriceJob{
		{"Binance", "https://api.binance.com/api/v3/ticker/price?symbol=BTCUSDT"},
		{"Binance-ETH", "https://api.binance.com/api/v3/ticker/price?symbol=ETHUSDT"},
		{"KuCoin", "https://api.kucoin.com/api/v1/market/orderbook/level1?symbol=BTC-USDT"},
		{"ByBit", "https://api.bybit.com/v5/market/tickers?category=linear&symbol=BTCUSDT"},
	}

	http.HandleFunc("/price", func(w http.ResponseWriter, r *http.Request) {
		jobs := make(chan PriceJob, len(targets))
		results := make(chan PriceResult, len(targets))
		var wg sync.WaitGroup

		for w := 1; w <= 3; w++ {
			wg.Add(1)
			go fetchPriceWorker(w, jobs, results, &wg)
		}

		for _, t := range targets {
			jobs <- t
		}
		close(jobs)

		go func() {
			wg.Wait()
			close(results)
		}()
		var finalResults []PriceResult
		for res := range results {
			finalResults = append(finalResults, res)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(finalResults)
	})
	fmt.Println("Server đang chạy tại http://localhost:8080/price")
	http.ListenAndServe(":8080", nil)
	/*fmt.Printf("%-15s | %-12s | %-10s\n", "EXCHANGE", "PRICE (USD)", "LATENCY")
	fmt.Println("--------------------------------------------")
	for r := range results {
		if r.Err != nil {
			fmt.Printf("%-15s | %-12s | %v\n", r.Exchange, "ERROR", r.Err)
		} else {
			fmt.Printf("%-15s | %-12s | %v\n", r.Exchange, r.Price, r.Duration)
		}
	}*/
}
