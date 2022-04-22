package controller

import (
	"JByun/logic"
	"JByun/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	//1.获取参数及参数校验
	p := new(models.ParamSignUp)
	err := c.ShouldBind(&p)
	if err != nil {
		zap.L().Error("sign up with invalid param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParam, "sign up with invalid param")
		return
	}
	//2.注册业务
	if err := logic.SignUp(p); err != nil {
		if err.Error() == "user already exist" {
			zap.L().Error("user already exist", zap.Error(err))
			ResponseErrorWithMsg(c, CodeError, "user already exist")
			return
		}
		zap.L().Error("sign up failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeError, "sign up failed")
	}
	//3.返回信息
	ResponseSuccess(c, nil)
}
