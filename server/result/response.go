package result

import (
	"github.com/gin-gonic/gin"
	"linktree_server/server/result/code"
	"net/http"
)

// Response ...
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// APIResponse ....
func APIResponse(Ctx *gin.Context, start *code.Co, data interface{}) {
	co, message := code.Decode(start)

	if start == code.OK {
		Ctx.JSON(http.StatusOK, Response{
			Code: co,
			Msg:  message,
			Data: data,
		})
	} else {
		Ctx.JSON(http.StatusOK, Response{
			Code: co,
			Msg:  message,
			Data: nil,
		})
	}
}
