package controller

import (
	"JByun/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func FileUpLoadHandler(c *gin.Context) {
	// FormFile方法会读取参数“upload”后面的文件名，返回值是一个File指针，和一个FileHeader指针，和一个err错误。
	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		zap.L().Error("c.Request.FormFile failed", zap.Error(err))
		return
	}

	if err := logic.FileUpLoad(file, header); err != nil {
		zap.L().Error("logic.FileUpLoad failed", zap.Error(err))
		return
	}

	ResponseSuccess(c, nil)
}
