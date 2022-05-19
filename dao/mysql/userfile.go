package mysql

import (
	"JByun/models"
	"go.uber.org/zap"
)

func InsertUserFile(userFile *models.UserFile) error {
	sqlStr := `insert into user_file(user_id, user_name, file_sha1, file_name, file_size) values (?, ?, ?, ?, ?)`
	_, err := db.Exec(sqlStr, userFile.UserID, userFile.UserName, userFile.FileSha1, userFile.FileName, userFile.FileSize)
	if err != nil {
		zap.L().Error("db.Exec failed", zap.Error(err))
		return err
	}
	return nil
}
