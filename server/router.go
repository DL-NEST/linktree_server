package server

import (
	"github.com/gin-contrib/cors"
	_ "github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "linktree_server/docs"
	"linktree_server/server/api/fixed"
	v1 "linktree_server/server/api/v1"
	"linktree_server/server/middleware"
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
	server.Use(middleware.JsType())

	server.LoadHTMLGlob("web/dist/*.html")
	server.StaticFS("/assets", http.Dir("web/dist/assets"))
	server.StaticFile("/", "web/dist/index.html")
	// swagger生成的文档，更新 CMD swag init
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 固定的api
	fixed.Fixed(server)
	// 注册api
	v1.InjectV1(server)

	return server
}

// InitCORS   初始化跨域配置
func InitCORS(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*/*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
}
