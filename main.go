package main

import (
	"fmt"
	"linuxNet/bootstrap"
	"linuxNet/model/emqx"
	sock "linuxNet/model/socket"
	"linuxNet/server"
	"linuxNet/utils/logger"
	//"github.com/tensorflow/tensorflow/tensorflow/go"
)

func init() {
	// 获取控制台输入
	bootstrap.InitFlag()
	// 比main先加载
	bootstrap.InitApp()
	// 读取配置文件
	bootstrap.InitConfig()
}

func main() {
	// 设置日记级别
	logger.Level = logger.LevelDebug
	// 初始化socket连接池和数量
	sock.InitWsPool(10)
	// 初始化mqtt
	emqx.InitMqtt()
	// 获取服务
	service := server.InitRouter()
	// 启动服务
	port := fmt.Sprint(":", bootstrap.Conf.Service.Port)
	logger.Log().Info("监听服务端口" + port)
	logger.Log().Info("服务启动成功：http://localhost" + port + "\n")
	serverErr := service.Run(port)
	if serverErr != nil {
		logger.Log().Error("服务启动失败:%v", serverErr)
		return
	}
}
