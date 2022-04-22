package control

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"linktree_server/models/DB"
	"linktree_server/server/result"
	"linktree_server/server/result/code"
	"time"
)

// UserLoginParam 登录接收参数
type UserLoginParam struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
	Tel      int `form:"tel" json:"tel" binding:"required,min=1,max=40"`
}

type UserRegisterParam struct {
	UserName string `form:"user_name" json:"user_name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Tel      int `form:"tel" json:"tel" binding:"required"`
}


// UserLogin
// @Tags User
// @Summary 用户登录
// @Description 用户登录权鉴
// @Param data  body UserLoginService true "用户名, 密码, 验证码"
// @Accept application/json
// @Produce  application/json
// @Success 200 {object} UserLoginInfo
// @Router /user/login [post]
func UserLogin(c *gin.Context) {
	var userLoginParam UserLoginParam
	if err := c.ShouldBind(&userLoginParam); err != nil {
		result.APIResponse(c, code.ErrParam, err)
		return
	}
	result.APIResponse(c, code.OK, &UserLoginParam{
		UserName: "efw",
		Password: "sfv",
		Tel: 324525,
	})
}

// UserRegister 注册
func UserRegister(c *gin.Context) {
	var userRegisterParam UserRegisterParam
	if err := c.ShouldBind(&userRegisterParam); err != nil {
		result.APIResponse(c, code.ErrParam, err)
		return
	}
	DB.AddUser(&DB.User{
		ID:         uuid.NewV4(),
		Name:       userRegisterParam.UserName,
		Tel:        34234,
		CreateTime: time.Now(),
	})
	result.APIResponse(c, code.OK, "res")
}
