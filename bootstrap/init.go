package bootstrap

import (
	"fmt"
	"linktree_server/utils/logger"
)

// InitApp 初始化程序打印版本号
func InitApp() {
	// TODO 在这获取有没有版本更新
	appVersion := "3.14.5"
	gitUrl := "https://github.com/DL-NEST/linktree_server"
	newVersion := "4.8.5"
	var Version string

	if appVersion != newVersion {
		Version = appVersion + " \u001B[;34m-> " + newVersion + "\u001B[0m"
	}

	fmt.Printf("\u001B[;32m"+` 
                     _      _         _   _____                
                    | |    |_|       | | |_   _|               
                    | |     _  _ __  | | __| | _ __  ___   ___ 
                    | |    | || '_ \ | |/ /| || '__|/ _ \ / _ \
                    | |____| || | | ||   < | || |  |  __/|  __/
                    \_____/|_||_| |_||_|\_\\_/|_|   \___| \___|`+"\n\n\t"+
		`Version:`+"\u001B[;35m %v  \u001B[0m"+"\u001B[;32m"+"Github: %v"+"\u001B[0m\n", Version, gitUrl)
	fmt.Printf("\u001B[;32m" + "=====================================================================================\u001B[0m\n\n")

	GlobalConsole.OutFlag()
}

func OutInfo() string {
	port := fmt.Sprint(":", Conf.Service.Port)
	// 输出信息
	logger.Log().Info("监听服务端口" + port)
	logger.Log().Info("服务启动成功: http://localhost" + port)
	logger.Log().Info("swag文档地址: http://localhost" + port + "/swagger/index.html")
	fmt.Printf("\n")
	return port
}