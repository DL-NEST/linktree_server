package server

import (
	"github.com/gin-contrib/cors"
	_ "github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	v1 "linuxNet/server/api/v1"
	"linuxNet/server/middleware"
	"net/http"
	"time"
)

func InitGin() {
	gin.ForceConsoleColor()
	// 设置日记的输出,io.MultiWriter// io的写入,文件流和控制台
	// gin.DefaultErrorWriter 和 gin.DefaultWriter
	//gin.DefaultErrorWriter = io.MultiWriter(middleware.LogFile(), os.Stdout)
	gin.SetMode(gin.ReleaseMode)
}

func InitRouter() *gin.Engine {
	InitGin()
	// 配置gin 使用自定义中间件
	server := gin.New()
	server.Use(middleware.Logger(), gin.Recovery())
	// Gzip
	server.Use(gzip.Gzip(gzip.DefaultCompression))
	// 全局请求中间件
	//server.Use(middleware.GlobalAuth())
	//server.Use(middleware.JsType())
	server.LoadHTMLGlob("web/dist/*.html")
	server.StaticFS("/assets", http.Dir("web/dist/assets"))
	server.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	// 注册api
	v1.InjectV1(server)

	return server
}

// InitCORS   初始化跨域配置
func InitCORS(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*/*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
}
