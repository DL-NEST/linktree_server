package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"linuxNet/utils/logger"
	"time"
)

// Logger 自定义日记中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		stateTime := time.Now()

		c.Next()

		cost := time.Since(stateTime)

		msg := logger.GinMsg{
			Status: c.Writer.Status(),
			Proto:  strAlign(15, cost.String()),
			Host:   c.ClientIP(),
			Method: c.Request.Method,
			Path:   c.Request.URL.Path,
		}
		logger.Log().GinLog(msg)
	}
}

func strAlign(num int, str string) string {
	var setUp string
	for i := 0; i < num-len(str); i++ {
		setUp += " "
	}
	return fmt.Sprintf("%v%v", setUp, str)
}
