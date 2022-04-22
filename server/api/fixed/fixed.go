package fixed

import (
	"github.com/gin-gonic/gin"
	"linktree_server/server/control"
	"linktree_server/server/result"
	"linktree_server/server/result/code"
)

// Fixed 不需要验证就能的请求的api
func Fixed(server *gin.Engine) {
	// 一些检测函数
	server.GET("healthy", func(context *gin.Context) {
		result.APIResponse(context, code.OK, "")
	})
	// 公开,无Auth的接口
	server.POST("/login", control.UserLogin)
	server.POST("/register", control.UserRegister)

}
