package control

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"linktree_server/server/result"
	"linktree_server/server/result/code"
	"linktree_server/utils/logger"
)

func UploadOne(c *gin.Context) {
	file, _ := c.FormFile("file")
	toFilePath := fmt.Sprintf("file/%v", file.Filename)
	logger.Log().Info(file.Filename+"-to-", toFilePath)
	// 上传文件至指定目录
	err := c.SaveUploadedFile(file, toFilePath)
	if err != nil {
		result.APIResponse(c, code.ErrUpload, "")
		return
	}
	result.APIResponse(c, code.OK, file.Filename)
}

func UploadList(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["files"]
	resList := make(map[string]map[int]string)
	errList := make(map[int]string)
	sucList := make(map[int]string)
	for i, file := range files {
		toFilePath := fmt.Sprintf("file/%v", file.Filename)
		err := c.SaveUploadedFile(file, toFilePath)
		if err != nil {
			errList[i] = fmt.Sprintf("%v Error:%v", file.Filename, err)
		} else {
			sucList[i] = file.Filename
		}
	}
	resList["errList"] = errList
	resList["sucList"] = sucList
	result.APIResponse(c, code.OK, resList)
}
