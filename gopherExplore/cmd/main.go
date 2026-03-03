package main

import (
	"flag"
	"fmt"
	"sync"

	//"time"
	"gopherExplore/internal/crawler"
	"gopherExplore/internal/models"
)

func main() {
	pathStr := flag.String("path", ".", "Root Directory")
	extStr := flag.String("ext", ".pdf", "File Extension")
	flag.Parse()
	results := make(chan models.FileResult, 100)
	var wg sync.WaitGroup

	wg.Add(1)
	go crawler.Crawler(*pathStr, *extStr, results, &wg)
	go func() {
		wg.Wait()
		close(results)
	}()
	for res := range results {
		fmt.Printf("Found : %s (%d Byte)\n", res.Path, res.Size)
	}
}
