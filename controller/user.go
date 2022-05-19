package controller

import (
	"JByun/logic"
	"JByun/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//RegisterHandler 用户注册
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

//LoginHandler 用户登陆
func LoginHandler(c *gin.Context) {
	//获取参数及参数校验
	p := new(models.ParamSignIn)
	err := c.ShouldBindJSON(&p)
	if err != nil {
		zap.L().Error("login with invalid param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParam, "login with invalid param")
		return
	}
	//登陆业务
	token, err := logic.Login(p)
	if err != nil {
		zap.L().Error("login failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeError, "login failed")
		return
	}
	//返回响应
	ResponseSuccess(c, token)
}
