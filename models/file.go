package models

import "time"

type File struct {
	FileID     int16
	FileSize   int64
	Status     int16
	FileSha1   string
	FileName   string
	FileAddr   string
	UpLoadTime time.Time
	UpDateTime time.Time
}
