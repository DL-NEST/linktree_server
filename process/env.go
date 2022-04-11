package process

import (
	"linuxNet/bootstrap"
	"linuxNet/server/model/emqx"
	"linuxNet/server/model/socket/wsPool"
)

var Env = new(Process)

type Process struct {
	WsPool wsPool.Pool
	MqState *emqx.MqState
	Config *bootstrap.Config
}

//// InitEnv 初始化全局变量
//func InitEnv(env *Process) {
//	Env = env
//}