package controller

import (
	"JByun/logic"
	"JByun/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"strconv"
)

func FileUpLoadHandler(c *gin.Context) {
	// FormFile方法会读取参数“upload”后面的文件名，返回值是一个File指针，和一个FileHeader指针，和一个err错误。
	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		zap.L().Error("c.Request.FormFile failed", zap.Error(err))
		ResponseError(c, CodeError)
	}

	//获取当前用户信息
	userid, _ := c.Get(CtxUserID)
	username, _ := c.Get(CtxUserName)

	if err := logic.FileUpLoad(file, header, userid.(int64), username.(string)); err != nil {
		zap.L().Error("logic.FileUpLoad failed", zap.Error(err))
		ResponseError(c, CodeError)
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
	//获取参数及校验
	// FormFile方法会读取参数“upload”后面的文件名，返回值是一个File指针，和一个FileHeader指针，和一个err错误。
	//获取分块文件
	file, header, err := c.Request.FormFile("file_block")
	if err != nil {
		zap.L().Error("c.Request.FormFile failed", zap.Error(err))
		ResponseError(c, CodeError)
	}
	uploadID := c.Request.Form.Get("upload_id")
	blockIdx := c.Request.Form.Get("block_idx")

	//获取当前用户信息
	userid, _ := c.Get(CtxUserID)
	username, _ := c.Get(CtxUserName)

	if err := logic.ChunkUpLoad(file, header, uploadID, blockIdx, userid.(int64), username.(string)); err != nil {
		zap.L().Error("logic.ChunkUpLoad failed", zap.Error(err))
		ResponseError(c, CodeError)
	}

	ResponseSuccess(c, nil)
}

func ChunkCompleteHandler(c *gin.Context) {
	//获取参数及校验
	uploadID := c.Request.Form.Get("upload_id")
	fileName := c.Request.Form.Get("file_name")
	filesize := c.Request.Form.Get("file_size")
	fileSize, _ := strconv.ParseInt(filesize, 10, 64)
	fileSha1 := c.Request.Form.Get("file_sha1")
	file := &models.File{
		FileSize: fileSize,
		FileSha1: fileSha1,
		FileName: fileName,
		FileAddr: "",
	}
	fmt.Println(file)

	userid, _ := c.Get(CtxUserID)
	username, _ := c.Get(CtxUserName)
	userFile := &models.UserFile{
		UserID:   userid.(int64),
		FileSize: fileSize,
		UserName: username.(string),
		FileSha1: fileSha1,
		FileName: fileName,
	}
	fmt.Println(userFile)

	//业务逻辑
	if err := logic.ChunkComplete(file, userFile, uploadID); err != nil {
		zap.L().Error("logic.ChunkComplete failed", zap.Error(err))
		ResponseError(c, CodeError)
	} else {
		//返回响应
		ResponseSuccess(c, nil)
	}
}
