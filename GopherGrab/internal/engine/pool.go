package engine

import (
	"GopherGrab/internal/models"
	"sync"
)

func worker(jobs <-chan models.DownloadJob, results chan<- models.DownloadResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		res := DownloadFromURL(job)
		results <- res
	}
}

func StartPool(workerCount int, jobsList []models.DownloadJob) <-chan models.DownloadResult {
	jobs := make(chan models.DownloadJob, len(jobsList))
	//results := make(chan models.DownloadResult, len(jobsList))
	results := make(chan models.DownloadResult)
	var wg sync.WaitGroup
	for w := 0; w < workerCount; w++ {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}

	for _, job := range jobsList {
		jobs <- job
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()
	return results

}
