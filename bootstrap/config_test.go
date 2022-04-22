package bootstrap

import "testing"

// 无配置文件时新建一个配置文件
func TestCreateConfig(t *testing.T) {
	CreateConfig("test/conf.yaml")
	t.Log("创建成功")
}

func TestGetDbConf(t *testing.T) {
	//GetDbConf("test/conf.yaml")
	t.Log(Conf.Service.Port)
}
