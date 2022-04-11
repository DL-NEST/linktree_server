package bootstrap

import (
	"fmt"
)

// InitApp 初始化程序打印版本号
func InitApp() {
	appVersion := "3.14.5"
	fmt.Printf("\u001B[;32m"+` 
	 _      _         _   _____                
	| |    |_|       | | |_   _|               
	| |     _  _ __  | | __| | _ __  ___   ___ 
	| |    | || '_ \ | |/ /| || '__|/ _ \ / _ \
	| |____| || | | ||   < | || |  |  __/|  __/
	\_____/|_||_| |_||_|\_\\_/|_|   \___| \___|
    Version:`+"%v\t Github:https://www.baidu.com/"+"\u001B[0m\n",appVersion)

}
