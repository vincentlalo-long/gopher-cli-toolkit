package collector

import (
	"GopherGrab/internal/models"
	"bufio"

	//"errors"
	//"go/scanner"
	"os"
	"path/filepath"
	"strings"
)

// Colector Folder scan file.txt to parse URL
func CollectFromFolder(folderPath string, saveDir string) ([]models.DownloadJob, error) {
	var allJobs []models.DownloadJob
	err := filepath.WalkDir(folderPath, func(path string, dir os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !dir.IsDir() && strings.ToLower(filepath.Ext(path)) == ".txt" {
			jobs, _ := CollectFromFile(path, saveDir)
			allJobs = append(allJobs, jobs...)
		}

		return nil
	})
	return allJobs, err
}

func CollectFromFile(filePath string, saveDir string) ([]models.DownloadJob, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var jobs []models.DownloadJob
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := strings.TrimSpace(scanner.Text())
		if url == "" || !strings.HasPrefix(url, "http") {
			continue
		}
		fileName := filepath.Base(url)
		jobs = append(jobs, models.DownloadJob{
			URL:      url,
			SavePath: saveDir,
			FileName: fileName,
		})
	}
	return jobs, err
}
