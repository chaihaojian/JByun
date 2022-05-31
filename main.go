package main

import (
	"JByun/config"
	"JByun/dao/mysql"
	"JByun/dao/redis"
	"JByun/logger"
	"JByun/pkg/snowflake"
	"JByun/routes"
	"fmt"
	"go.uber.org/zap"
)

func main() {
	//加载服务端配置
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

	//初始化Redis连接
	if err := redis.Init(); err != nil {
		fmt.Printf("Init redis faild, err:%v\n", err)
	}
	defer redis.Close()

	//初始化ID生成器
	if err := snowflake.Init(); err != nil {
		fmt.Printf("Init snowflake faild, err:%v\n", err)
	}

	//注册路由
	r := routes.Setup()

	r.Run()
}
