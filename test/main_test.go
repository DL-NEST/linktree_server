package test

import (
	"fmt"
	"linktree_server/bootstrap"
	"testing"
)

/*--- 一个测试文件 ---*/
// go的协程

// 读取配置文件测试
func TestReadConfig(t *testing.T) {
	Config := bootstrap.ReadConfig("./config.yaml")
	fmt.Println(Config.Service.Port)
	t.Log(Config.Service.Port)
}
