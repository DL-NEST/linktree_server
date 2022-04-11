package test

import (
	"fmt"
	"linuxNet/server/config"
	"testing"
)

/*--- 一个测试文件 ---*/
// go的协程

// 读取配置文件测试
func TestReadConfig(t *testing.T){
	Config := config.ReadConfig("./config.yaml")
	fmt.Println(Config.Service.Port)
	t.Log(Config.Service.Port)
}
