package mysql

import (
	"JByun/models"
	"go.uber.org/zap"
)

func OnFileUpLoad(fileMeta *models.File) error {
	sqlStr := `insert into file(file_sha1, file_name, file_size, file_addr) values (?, ?, ?, ?)`
	_, err := db.Exec(sqlStr, fileMeta.FileSha1, fileMeta.FileName, fileMeta.FileSize, fileMeta.FileAddr)
	if err != nil {
		zap.L().Error("db.Exec failed", zap.Error(err))
		return err
	}
	return nil
}
