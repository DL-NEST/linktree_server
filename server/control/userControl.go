package control

import (
	"github.com/gin-gonic/gin"
	"linuxNet/server/result"
	"linuxNet/server/result/code"
	"linuxNet/server/service"
)

type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

type UserRegisterService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
	Time     string `form:"time" json:"time" binding:"required,min=8,max=40"`
}

func UserLogin(c *gin.Context) {
	var userLoginService UserLoginService
	if err := c.ShouldBind(&userLoginService); err != nil {
		result.APIResponse(c, code.ErrAccessRight, service.GetNetInfo())
		return
	}
	c.JSON(200, gin.H{"user_name": userLoginService.UserName, "password": userLoginService.Password})
}

func UserRegister(c *gin.Context) {
	var userLoginService UserRegisterService
	if err := c.ShouldBind(&userLoginService); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	c.JSON(200, gin.H{"user_name": userLoginService.UserName, "password": userLoginService.Password})
}
