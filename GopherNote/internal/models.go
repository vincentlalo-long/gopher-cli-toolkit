package internal

import "time"

type Note struct {
	Id       uint64
	FileName string
	Size     uint64
	ModTime  time.Time
}
