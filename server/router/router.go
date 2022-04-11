package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"linuxNet/bootstrap"
	"linuxNet/server/control"
	"linuxNet/server/handler"
	sock "linuxNet/server/model/socket"
	"linuxNet/server/model/socket/wsPool"
	"linuxNet/server/result"
	"linuxNet/server/result/code"
	"linuxNet/server/service"
	ctrl "linuxNet/utils/ctrlOut"
	"net/http"
	"os"
)

func InitGin() {
	gin.ForceConsoleColor()
	// 设置日记的输出,io.MultiWriter// io的写入,文件流和控制台
	// gin.DefaultErrorWriter 和 gin.DefaultWriter
	gin.DefaultErrorWriter = io.MultiWriter(handler.LogFile(), os.Stdout)
	gin.SetMode(gin.ReleaseMode)
}
func ginConfig(server *gin.Engine) {
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	// 上传文件大小
	server.MaxMultipartMemory = 8 << 20 // 8 MiB
}

func InitRouter() *gin.Engine {
	InitGin()
	// 配置gin 使用自定义中间件
	server := gin.Default()
	ginConfig(server)
	// 全局请求中间件
	//server.Use(handler.GlobalAuth())
	//server.Use(handler.JsType())
	// 网页
	server.LoadHTMLGlob("web/dist/*.html")
	server.StaticFS("/assets", http.Dir("web/dist/assets"))

	server.GET("/", handler.JsType(), func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	// 声明请求组
	// user
	user := server.Group("/user")
	{
		user.POST("/login", control.UserLogin)
		user.POST("/register", control.UserRegister)
	}
	// sys /*系统状态的获取和操作
	sys := server.Group("/sys")
	{
		sys.GET("/getCpuInfo", control.GetSysInfo)
		sys.GET("/getMemInfo", control.GetMemInfo)
		sys.GET("/getDiskInfo", control.GetDiskInfo)
		sys.GET("/getHostInfo", control.GetHostInfo)
		sys.GET("/getNetInfo", control.GetNetInfo)
		sys.GET("/poolList", func(context *gin.Context) {
			result.APIResponse(context, code.OK, 	sock.WsPool.GetPoolList())
		})
	}
	// file /*上传文件
	file := server.Group("/file")
	{
		file.POST("/upload", control.UploadOne)
		file.POST("/uploadList", control.UploadList)
	}
	// websocket
	socket := server.Group("/socket")
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
	// 启动服务
	port := fmt.Sprint(":",bootstrap.Conf.Service.Port)
	ctrl.Success(fmt.Sprint("服务启动成功：http://localhost",port))
	serverErr := server.Run(port)
	if serverErr != nil {
		ctrl.Error(fmt.Sprintf("服务启动失败:%v", serverErr))
		return nil
	}
	return server
}
