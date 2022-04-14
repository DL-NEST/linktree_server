package control

import (
	"github.com/gin-gonic/gin"
	"linktree_server/model/emqx"
	"linktree_server/server/result"
	"linktree_server/server/result/code"
	"linktree_server/server/service"
)

// SysServer /* 获取系统信息

func GetSysInfo(c *gin.Context) {
	result.APIResponse(c, code.OK, service.GetSysInfo())
}
func GetMemInfo(c *gin.Context) {
	result.APIResponse(c, code.OK, service.GetMemInfo())
}
func GetDiskInfo(c *gin.Context) {
	result.APIResponse(c, code.OK, service.GetDiskInfo())
}
func GetHostInfo(c *gin.Context) {
	result.APIResponse(c, code.OK, service.GetHostInfo())
}
func GetNetInfo(c *gin.Context) {
	result.APIResponse(c, code.OK, service.GetNetInfo())
}
func GetMqttList(c *gin.Context) {
	result.APIResponse(c, code.OK, service.GetMqttList(emqx.MqttClient))
}
