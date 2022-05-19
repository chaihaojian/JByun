package models

import "time"

type UserFile struct {
	UserID         int64
	FileSize       int64
	UserName       string
	FileSha1       string
	FileName       string
	UpLoadTime     time.Time
	LastUpDateTime time.Time
}
