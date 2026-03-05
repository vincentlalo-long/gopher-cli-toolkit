package crawler

import (
	"gopherExplore/internal/models"
	"os"
	"path/filepath"

	//"strings"
	"sync"
)

var sem = make(chan struct{}, 20)

func Crawler(config models.FilterConfig, dir string, results chan<- models.FileResult, wg *sync.WaitGroup) {
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
			go Crawler(config, path, results, wg)
		} else {
			info, err := entry.Info()
			if err != nil {
				continue
			}
			if config.Matches(entry.Name(), info.Size(), filepath.Ext(path)) {
				results <- models.FileResult{
					Path: path,
					Size: uint64(info.Size()),
				}
			}
		}

	}

}
