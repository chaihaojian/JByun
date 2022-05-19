package mysql

import (
	"JByun/models"
	"database/sql"
	"go.uber.org/zap"
)

func InsertFile(fileMeta *models.File) error {
	sqlStr := `insert into file(file_sha1, file_name, file_size, file_addr) values (?, ?, ?, ?)`
	_, err := db.Exec(sqlStr, fileMeta.FileSha1, fileMeta.FileName, fileMeta.FileSize, fileMeta.FileAddr)
	if err != nil {
		zap.L().Error("db.Exec failed", zap.Error(err))
		return err
	}
	return nil
}

func CheckFileExist(file *models.File) bool {
	sqlStr := `select file_size from file where file_sha1 = ?`
	err := db.Get(file, sqlStr, file.FileSha1)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		zap.L().Error("db.Get failed", zap.Error(err))
		return false
	}
	return true
}
