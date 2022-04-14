package control

import (
	"github.com/gin-gonic/gin"
	"linktree_server/server/result"
	"linktree_server/server/result/code"
	"linktree_server/server/service"
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

// UserLogin
// @Tags User
// @Summary 用户登录
// @Description 用户登录权鉴
// @Param data body UserLoginService true "用户名, 密码, 验证码"
// @Accept application/json
// @Produce  application/json
// @Success 200 {object} UserRegisterService
// @Router /user/login [post]
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
