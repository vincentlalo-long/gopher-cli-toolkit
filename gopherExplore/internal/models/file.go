package models

import "time"

type FileResult struct {
	Path    string
	Size    uint64
	ModTime time.Time
}
