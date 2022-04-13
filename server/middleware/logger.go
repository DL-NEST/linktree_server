package middleware

import (
	"github.com/gin-gonic/gin"
	"linuxNet/utils/logger"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		stateTime := time.Now()

		c.Next()

		cost := time.Since(stateTime)
		msg := logger.GinMsg{
			Status: c.Writer.Status(),
			Proto:  cost.String(),
			Host:   c.ClientIP(),
			Method: c.Request.Method,
			Path:   c.Request.URL.Path,
		}
		logger.Log().GinLog(msg)
	}
}
