package main

import (
	"GopherGrab/internal/collector"
	"GopherGrab/internal/engine"
	"GopherGrab/internal/models"
	"GopherGrab/internal/ui"
	"flag"
	"fmt"
	"time"
)

func main() {

	var allJobs []models.DownloadJob
	filePtr := flag.String("file", "", "file .txt link")
	dirPtr := flag.String("dir", "", "link dir")
	outPtr := flag.String("out", "./downloads", "output")
	workerPtr := flag.Int("worker", 5, "number of worker")
	flag.Parse()

	if *filePtr != "" {
		jobs, _ := collector.CollectFromFile(*filePtr, *outPtr)
		allJobs = append(allJobs, jobs...)
	}

	if *dirPtr != "" {
		jobs, _ := collector.CollectFromFolder(*dirPtr, *outPtr)
		allJobs = append(allJobs, jobs...)
	}
	if len(allJobs) == 0 {
		fmt.Println("Error no file found")
	}

	start := time.Now()

	ui.PrintHeader(*workerPtr, len(allJobs))

	results := engine.StartPool(*workerPtr, allJobs)

	var successCount, failCount int
	var totalBytes int64
	for res := range results {
		if res.Error != nil {
			failCount++
			fmt.Printf("%s[FAILED]%s %s - Error: %v\n", ui.ColorRed, ui.ColorReset, res.URL, res.Error)
		} else {
			successCount++
			fmt.Printf("%s[SUCCESSED]%s %s\n", ui.ColorGreen, ui.ColorReset, res.URL)
		}
	}
	ui.PrintDivider()
	fmt.Printf("%s Complete in : %v%s\n", ui.ColorYellow, time.Since(start), ui.ColorReset)
	fmt.Printf("Success: %d |  Fail : %d | Total : %s\n",
		successCount, failCount, ui.FormatBytes(totalBytes))

}
