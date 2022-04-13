package middleware

import (
	"github.com/gin-gonic/gin"
	"regexp"
)

// JsType js文件请求添加Content-Type头
func JsType() gin.HandlerFunc {
	return func(c *gin.Context) {
		getPath := c.Request.URL.Path
		cs, _ := regexp.MatchString("(^)?\\.js", getPath)
		if cs {
			c.Header("Content-Type", "application/javascript; charset=utf-8")
		}
		c.Next()
	}
}
