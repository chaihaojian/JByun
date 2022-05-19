package logic

import (
	"JByun/dao/mysql"
	"JByun/models"
	"JByun/pkg/jwt"
	"JByun/pkg/snowflake"
	"errors"
	"go.uber.org/zap"
	"time"
)

func Register(p *models.ParamSignUp) error {
	//1.校验用户是否已存在
	exist, err := mysql.CheckUserExist(p.Phone)
	if err != nil {
		zap.L().Error("mysql.CheckUserExist failed", zap.Error(err))
		return err
	}
	if exist {
		err = errors.New("user already exist")
		return err
	}
	//2.生成ID,构造用户结构体
	id := snowflake.GenID()
	user := models.User{
		UserID:     id,
		PhoneValid: 0,
		EmailValid: 0,
		Gender:     0,
		Status:     0,
		Username:   p.Username,
		Password:   p.Password,
		Phone:      p.Phone,
		Email:      "",
		SignUpTime: time.Time{},
		LastActive: time.Time{},
	}
	//3.插入数据库
	if err = mysql.InsertUser(&user); err != nil {
		zap.L().Error("mysql.InsertUser failed", zap.Error(err))
		return err
	}
	//4.返回信息
	return err
}

func Login(p *models.ParamSignIn) (token string, err error) {
	user := &models.User{
		Username: p.Username,
	}
	if err = mysql.QueryUserByName(user); err != nil {
		zap.L().Error("mysql.QueryUserByName failed", zap.Error(err))
		return "", err
	}
	if p.Password != user.Password {
		err = errors.New("wrong password")
		return "", err
	}
	token, err = jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		zap.L().Error("jwt.GenToken failed", zap.Error(err))
		return "", err
	}
	return token, nil
}
