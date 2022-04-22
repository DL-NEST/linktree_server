package main

import (
	"linktree_server/bootstrap"
	"linktree_server/models/DB"
	"linktree_server/models/emqx"
	"linktree_server/server"
	"linktree_server/utils/logger"
	//"github.com/tensorflow/tensorflow/tensorflow/go"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/
//go:generate swag init

func init() {
	// 获取控制台输入
	bootstrap.InitFlag()
	// 比main先加载
	bootstrap.InitApp()
	// 读取配置文件
	bootstrap.InitConfig()
	// 连接化数据库
	DB.CreateDBLink()
	// 初始化socket连接池和数量
	//sock.InitWsPool(1)
	// 初始化mqtt
	emqx.InitMqtt()
}

func main() {
	// 获取服务
	service := server.InitRouter()
	// 启动服务
	port := bootstrap.OutInfo()
	serverErr := service.Run(port)
	if serverErr != nil {
		logger.Log().Error("服务启动失败:%v", serverErr)
		return
	}
}
