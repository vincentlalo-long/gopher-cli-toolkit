package crawler

import (
	"gopherExplore/internal/models"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var sem = make(chan struct{}, 20)

func Crawler(dir, ext string, results chan<- models.FileResult, wg *sync.WaitGroup) {
	defer wg.Done()
	sem <- struct{}{}
	defer func() {
		<-sem
	}()
	entries, err := os.ReadDir(dir)
	if err != nil {
		return
	}
	for _, entry := range entries {
		path := filepath.Join(dir, entry.Name())
		//var wg sync.WaitGroup()
		if entry.IsDir() {
			wg.Add(1)
			go Crawler(path, ext, results, wg)
		} else if strings.EqualFold(strings.ToLower(filepath.Ext(path)), strings.ToLower(ext)) {
			if info, err := entry.Info(); err == nil {
				results <- models.FileResult{
					Path:    path,
					Size:    uint64(info.Size()),
					ModTime: info.ModTime(),
				}
			}

		}

	}

}
