package engine

import (
	"GopherGrab/internal/models"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadFromURL(job models.DownloadJob) models.DownloadResult {
	result := models.DownloadResult{URL: job.URL}

	//Make save folder
	err := os.MkdirAll(job.SavePath, os.ModePerm)
	if err != nil {
		result.Error = err
		return result
	}

	// full path for file
	fullPath := filepath.Join(job.SavePath, job.FileName)
	result.FullPath = fullPath

	// Request for http get
	resp, err := http.Get(job.URL)
	if err != nil {
		result.Error = err
		return result
	}
	defer resp.Body.Close()
	//Check status http code
	if resp.StatusCode != http.StatusOK {
		result.Error = os.ErrNotExist
		return result
	}

	//  make file in hardware
	out, err := os.Create(fullPath)
	if err != nil {
		result.Error = err
		return result
	}
	defer out.Close()

	//streaming
	n, err := io.Copy(out, resp.Body)
	if err != nil {
		result.Error = err
		return result
	}
	result.Size = uint64(n)

	return result

}
