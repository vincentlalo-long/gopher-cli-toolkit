package main

import (
	"io"
	"net/http"
	"sync"
	"time"
)

func fetchPriceWorker(id int, jobs <-chan PriceJob, results chan<- PriceResult, wg *sync.WaitGroup) {
	defer wg.Done()

	client := http.Client{
		Timeout: 60 * time.Second,
	}
	for job := range jobs {
		start := time.Now()
		resp, err := client.Get(job.URL)
		result := PriceResult{Exchange: job.Exchange, Err: err}
		if err == nil {
			//encode json direct from response

			body, _ := io.ReadAll(resp.Body)
			resp.Body.Close()

			if handler, ok := Handlers[job.Exchange]; ok {
				price, parseErr := handler.Parse(body)
				result.Price = price
				result.Err = parseErr
			} else {
				result.Err = err
			}

		}
		result.Duration = time.Since(start)
		results <- result
	}

}
