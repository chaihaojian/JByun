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
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("signup with invalid param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParam, "signup with invalid param")
		return
	}
	//2.注册业务
	logic.SignUp(p)
	//3.返回信息
	ResponseSuccess(c, nil)
}
