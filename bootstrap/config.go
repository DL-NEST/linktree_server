package bootstrap

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"linktree_server/utils/logger"
	"log"
	"os"
)

/*
	读取配置文件
*/
var Conf = new(Config)

// ReadConfig 读取配置文件
func ReadConfig(configPath string) *Config {
	var conf = new(Config)
	config, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Print(err)
	}
	// yaml文件内容影射到结构体中
	err1 := yaml.Unmarshal(config, &conf)
	if err1 != nil {
		fmt.Printf("配置文件读取失败")
	}
	return conf
}

// CreateConfig 创建配置文件
func CreateConfig(path string) {
	var conf = new(Config)
	conf.Service.Port = 5241
	confData, err := yaml.Marshal(&conf)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	e := os.WriteFile(path, confData, os.ModeAppend)
	if e != nil {
		return
	}
}


func InitConfig() {
	if GlobalConsole.Conf != "null" {
		// 判断文件是否存在
		_, err := os.Stat(GlobalConsole.Conf)
		if err != nil {
			logger.Log().Warning("配置文件不存在,在路径下创建配置文件")
			CreateConfig(GlobalConsole.Conf)
		}
		logger.Log().Info("配置文件存在,读取配置文件")
		Conf = ReadConfig(GlobalConsole.Conf)
	} else {
		rootPath := "./conf.yaml"
		logger.Log().Info("未指定配置文件")
		_, err := os.Stat(rootPath)
		if err != nil {
			logger.Log().Warning("配置文件不存在,在根路径下创建配置文件")
			CreateConfig(rootPath)
		} else {
			logger.Log().Info("配置文件存在,读取配置文件")
			Conf = ReadConfig(rootPath)
		}
	}
}

func GetDbConf() bool {
	if Conf.DB.Port != "" && Conf.DB.Host != "" {
		return true
	}
	return false
}