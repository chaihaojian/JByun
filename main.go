package main

import (
	"JByun/config"
	"JByun/dao/mysql"
	"JByun/logger"
	"JByun/routes"
	"fmt"
	"go.uber.org/zap"
)

func main() {
	//加载配置
	if err := config.Init(); err != nil {
		fmt.Printf("Init config faild, err:%v\n", err)
	}
	//初始化日志
	if err := logger.Init(); err != nil {
		fmt.Printf("Init logger faild, err:%v\n", err)
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")
	//初始化MySQL连接
	if err := mysql.Init(); err != nil {
		fmt.Printf("Init mysql faild, err:%v\n", err)
	}
	defer mysql.Close()
	//注册路由
	r := routes.Setup()

	r.Run()
}
