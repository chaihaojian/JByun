package models

import "time"

type File struct {
	FileID     int16     `json:"id"`
	FileSize   int64     `json:"file_size" db:"file_size"`
	Status     int16     `json:"status"`
	FileSha1   string    `json:"file_sha1" binding:"required"`
	FileName   string    `json:"file_name" binding:"required"`
	FileAddr   string    `json:"file_addr"`
	UpLoadTime time.Time `json:"upload_time"`
	UpDateTime time.Time `json:"update_time"`
}
