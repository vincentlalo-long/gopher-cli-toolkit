package models

type DownloadJob struct {
	URL      string
	SavePath string
	FileName string
}

type DownloadResult struct {
	URL      string
	FileName string
	Size     uint64
	Error    error
}
