package bootstrap

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)
/*
	读取配置文件
*/
var Conf = new(Config)

// Config json转换工具 https://oktools.net/json2go
type Config struct {
	Service struct {
		Port string `json:"port"`
	} `json:"server"`
	Mqtt struct {
		Host string `json:"host"`
		Port string `json:"port"`
		LoginName string `json:"loginName"`
		Password string `json:"password"`
	} `json:"mqtt"`
	Redis struct {
		Host string `json:"host"`
		Port string `json:"port"`
		LoginName string `json:"loginName"`
		Password string `json:"password"`
	} `json:"redis"`
	Emqx struct {
		Host string `json:"host"`
		Port string `json:"port"`
		LoginName string `json:"loginName"`
		Password string `json:"password"`
		ClientID string `json:"clientId"`
	} `json:"emqx"`
}

// ReadConfig 读取配置文件
func ReadConfig(configPath string) *Config {
	var conf = new(Config)
	config, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Print(err)
	}
	// yaml文件内容影射到结构体中
	err1:=yaml.Unmarshal(config,&conf)
	if err1!=nil{
		fmt.Println("error")
	}
	return conf
}

func InitConfig(configPath string)  {
	Conf = ReadConfig(configPath)
}