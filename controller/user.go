package controller

import (
	"JByun/logic"
	"JByun/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RegisterHandler(c *gin.Context) {
	//1.获取参数及参数校验
	p := new(models.ParamSignUp)
	err := c.ShouldBindJSON(&p)
	if err != nil {
		zap.L().Error("register with invalid param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParam, "sign up with invalid param")
		return
	}
	//2.注册业务
	if err := logic.Register(p); err != nil {
		if err.Error() == "user already exist" {
			zap.L().Error("user already exist", zap.Error(err))
			ResponseErrorWithMsg(c, CodeError, "user already exist")
			return
		}
		zap.L().Error("register failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeError, "sign up failed")
	}
	//3.返回信息
	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	//获取参数及参数校验
	p := new(models.ParamSignIn)
	//登陆逻辑
	err := c.ShouldBindJSON(&p)
	if err != nil {
		zap.L().Error("login with invalid param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParam, "sign up with invalid param")
		return
	}
	//返回响应
}
