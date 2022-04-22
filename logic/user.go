package logic

import (
	"JByun/dao/mysql"
	"JByun/models"
)

func SignUp(p *models.ParamSignUp) {
	//1.校验用户是否已存在
	mysql.CheckUserExist(p.Username)
	//2.生成ID,构造用户结构体
	//3.插入数据库
	//4.返回信息
}
