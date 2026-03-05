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
	queryStr := flag.String("query", "", "key ")
	minSize := flag.Int64("min-size", 0, "min size of file crawl")
	flag.Parse()
	config := models.FilterConfig{
		Root:    *pathStr,
		Ext:     *extStr,
		Query:   *queryStr,
		MinSize: *minSize,
	}
	results := make(chan models.FileResult, 100)
	var wg sync.WaitGroup

	wg.Add(1)
	go crawler.Crawler(config, *pathStr, results, &wg)
	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Println("==== Scanning ====")
	for res := range results {
		fmt.Printf("Found : %s (%d Byte)\n", res.Path, res.Size)
	}
}
