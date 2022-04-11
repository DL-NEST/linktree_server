package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		out := Log()
		Regime := time.Now()
		// 放行
		out.Printf(" | %d |  %v  | %v | %v | %v | %v \n",
			c.Writer.Status(),
			time.Since(Regime),
			c.Request.Method,
			c.Request.Proto,
			c.Request.Host,
			c.Request.URL.Path,
		)
		out.WithFields(logrus.Fields{
			"name": "hanyun",
		}).Info("记录一下日志", "Info")
		c.Next() // 请求后
	}
}

func LogFile() *os.File {
	src, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, os.ModeAppend)
	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}
	return src
}

func Log() *logrus.Logger {
	logger := logrus.New()
	logger.Out = LogFile()

	return logger
}
