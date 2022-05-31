package logic

import (
	"JByun/dao/mysql"
	"JByun/dao/redis"
	"JByun/models"
	"JByun/pkg/snowflake"
	"JByun/util"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"io"
	"math"
	"mime/multipart"
	"os"
)

func FileUpLoad(file multipart.File, header *multipart.FileHeader, userid int64, username string) error {
	// header调用Filename方法，就可以得到文件名
	filename := header.Filename

	// 创建一个文件，文件名为filename，这里的返回值newFile也是一个File指针
	addr := viper.GetString("file.path") + filename
	newFile, err := os.Create(addr)
	if err != nil {
		zap.L().Error("os.Create failed", zap.Error(err))
		return err
	}

	defer newFile.Close()

	// 将file的内容拷贝到newFile
	_, err = io.Copy(newFile, file)
	if err != nil {
		zap.L().Error("io.Copy failed", zap.Error(err))
		return err
	}

	//将文件信息保存进数据库
	fileMeta := &models.File{
		FileSize: util.GetFileSize(addr),
		Status:   0,
		FileSha1: util.GetFileSha1(addr),
		FileName: filename,
		FileAddr: addr,
	}
	fmt.Println(fileMeta.FileSha1)
	userFile := &models.UserFile{
		UserID:   userid,
		FileSize: fileMeta.FileSize,
		UserName: username,
		FileSha1: fileMeta.FileSha1,
		FileName: fileMeta.FileName,
	}
	if err = mysql.InsertFile(fileMeta); err != nil {
		zap.L().Error("mysql.InsertFile failed", zap.Error(err))
		return err
	}
	if err = mysql.InsertUserFile(userFile); err != nil {
		zap.L().Error("mysql.InsertUserFile failed", zap.Error(err))
		return err
	}

	return nil
}

func FastFileUpLoad(file *models.File, userid int64, username string) error {
	//查询文件表若该文件已存在，触发秒传，插入用户文件表
	if mysql.CheckFileExist(file) {
		userFile := &models.UserFile{
			UserID:   userid,
			FileSize: file.FileSize,
			UserName: username,
			FileSha1: file.FileSha1,
			FileName: file.FileName,
		}
		if err := mysql.InsertUserFile(userFile); err != nil {
			zap.L().Error("mysql.InsertUserFile failed", zap.Error(err))
			return err
		}
		return nil
	}
	//若不存在，秒传失败
	return errors.New("file not exist")
}

func ChunkInit(user *models.User, f *models.ChunkInitParam) error {
	//检查文件是否已存在，若已经存在，直接触发秒传
	file := &models.File{
		FileSha1: f.FileSha1,
		FileSize: f.FileSize,
		FileName: f.FileName,
	}
	if mysql.CheckFileExist(file) {
		userFile := &models.UserFile{
			UserID:   user.UserID,
			FileSize: file.FileSize,
			UserName: user.Username,
			FileSha1: file.FileSha1,
			FileName: file.FileName,
		}
		if err := mysql.InsertUserFile(userFile); err != nil {
			zap.L().Error("mysql.InsertUserFile failed", zap.Error(err))
			return err
		}
		return nil
	}
	//若不存在，初始化分块上传信息，并返回
	//初始化分块信息
	f.UpLoadID = snowflake.GenID()
	f.ChunkSize = viper.Get("chunk.size").(int64)
	f.ChunkCount = int64(math.Ceil(float64(f.FileSize) / float64(f.ChunkSize)))
	//将分块信息保存进redis
	if err := redis.InsertChunkInfo(f); err != nil {
		zap.L().Error("redis.InsertChunkInfo failed", zap.Error(err))
		return err
	}
	return nil
}
