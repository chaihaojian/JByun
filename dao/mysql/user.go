package mysql

import (
	"JByun/models"
	"go.uber.org/zap"
)

func CheckUserExist(phone string) (exist bool, err error) {
	sqlStr := "select count(*) from user where phone = ?"
	var count int
	err = db.Get(count, sqlStr, phone)
	if err != nil {
		zap.L().Error("db.Get failed", zap.Error(err))
		return false, err
	}
	if count >= 1 {
		return true, nil
	}
	return false, nil
}

func InsertUser(user *models.User) error {
	sqlStr := "insert into user(id, username, password, phone) values(?,?,?,?)"
	_, err := db.Exec(sqlStr, user.UserID, user.Username, user.Password, user.Phone)
	if err != nil {
		zap.L().Error("db.Exec failed", zap.Error(err))
	}
	return err
}
