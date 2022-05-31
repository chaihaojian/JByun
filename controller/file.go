package controller

import (
	"JByun/logic"
	"JByun/models"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func FileUpLoadHandler(c *gin.Context) {
	// FormFile方法会读取参数“upload”后面的文件名，返回值是一个File指针，和一个FileHeader指针，和一个err错误。
	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		zap.L().Error("c.Request.FormFile failed", zap.Error(err))
		return
	}

	//获取当前用户信息
	userid, _ := c.Get(CtxUserID)
	username, _ := c.Get(CtxUserName)

	if err := logic.FileUpLoad(file, header, userid.(int64), username.(string)); err != nil {
		zap.L().Error("logic.FileUpLoad failed", zap.Error(err))
		return
	}

	ResponseSuccess(c, nil)
}

func FastFileUpLoadHandler(c *gin.Context) {
	//获取参数
	f := new(models.File)
	err := c.ShouldBindJSON(&f)
	if err != nil {
		zap.L().Error("c.ShouldBindJSON failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParam, "fast upload failed")
	}
	//获取当前用户信息
	userid, _ := c.Get(CtxUserID)
	username, _ := c.Get(CtxUserName)
	//秒传业务逻辑
	if err = logic.FastFileUpLoad(f, userid.(int64), username.(string)); err != nil {
		if err == errors.New("file not exist") {
			ResponseError(c, CodeFastUpLoadFailed)
		}
		zap.L().Error("logic.FastFileUpLoad failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}
	//返回响应
	ResponseSuccess(c, nil)
}

func ChunkInitHandler(c *gin.Context) {
	//获取参数及参数校验

	//获取文件信息
	f := new(models.ChunkInitParam)
	if err := c.ShouldBindJSON(&f); err != nil {
		zap.L().Error("c.ShouldBindJSON failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParam, "chunk upload with invalid param failed")
	}
	//获取当前用户
	userid, _ := c.Get(CtxUserID)
	username, _ := c.Get(CtxUserName)
	u := &models.User{
		UserID:   userid.(int64),
		Username: username.(string),
	}
	//业务逻辑
	if err := logic.ChunkInit(u, f); err != nil {
		zap.L().Error("logic.ChunkInit failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeServerBusy, "chunk upload init failed")
	}
	//返回初始化信息
	ResponseSuccess(c, f)
}

func ChunkUpLoadHandler(c *gin.Context) {

}

func ChunkCompleteHandler(c *gin.Context) {

}
