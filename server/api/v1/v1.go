package v1

import (
	"github.com/gin-gonic/gin"
	sock "linktree_server/model/socket"
	"linktree_server/model/socket/wsPool"
	"linktree_server/server/control"
	"linktree_server/server/result"
	"linktree_server/server/result/code"
	"linktree_server/server/service"
)

func InjectV1(server *gin.Engine) {
	// api版本
	V1 := server.Group("/v1")

	user := V1.Group("/user")
	{
		user.POST("/login", control.UserLogin)
		user.POST("/register", control.UserRegister)
	}
	// sys /*系统状态的获取和操作
	sys := V1.Group("/sys")
	{
		sys.GET("/getCpuInfo", control.GetSysInfo)
		sys.GET("/getMemInfo", control.GetMemInfo)
		sys.GET("/getDiskInfo", control.GetDiskInfo)
		sys.GET("/getHostInfo", control.GetHostInfo)
		sys.GET("/getNetInfo", control.GetNetInfo)
		sys.GET("/poolList", func(context *gin.Context) {
			result.APIResponse(context, code.OK, sock.WsPool.GetPoolList())
		})
	}
	// file /*上传文件
	file := V1.Group("/file")
	{
		file.POST("/upload", control.UploadOne)
		file.POST("/uploadList", control.UploadList)
	}
	// websocket
	socket := V1.Group("/socket")
	socket.GET("/linkSocket", control.UpgradeSocket)
	// 广播到单个
	socket.GET("/getTheOne", func(context *gin.Context) {
		sock.WsPool.WriteOne("3bd68f68-dfdc-400f-bf3b-f7840edc0385", wsPool.Text, "发送到单个")
	})
	// 广播到用户组
	socket.GET("/getTheMore", func(context *gin.Context) {
		userList := []string{"3bd68f68-dfdc-400f-bf3b-f7840edc0385", ""}
		sock.WsPool.WriteMore(userList, wsPool.Text, "发送到用户组")
	})
	// 广播到全部
	socket.GET("/broadcast", func(context *gin.Context) {
		sock.WsPool.Broadcast(wsPool.Json, service.GetCpuInfo())
	})
	// 获取已经连接的用户列表
	socket.GET("/poolList", func(context *gin.Context) {
		result.APIResponse(context, code.OK, sock.WsPool.GetPoolList())
	})
}
