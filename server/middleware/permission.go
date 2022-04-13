package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"linuxNet/server/result"
	"linuxNet/server/result/code"
	"linuxNet/server/service"
)

func GlobalAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求path
		path := c.Request.URL.Path
		// 拦截列表
		interceptList := []string{"/user/login", "/user/register", "/", "/index.html"}
		token := c.GetHeader("Token")
		fmt.Println(token)
		if intercept(path, interceptList) {
			c.Next()
		} else {
			// 判断token
			//service.JudgeToken()
			result.APIResponse(c, code.ErrAccessRight, service.GetNetInfo())
			c.Abort()
		}
		return
	}
}

func intercept(path string, reList []string) bool {
	for i := range reList {
		if reList[i] == path {
			return true
		}
	}
	return false
}
