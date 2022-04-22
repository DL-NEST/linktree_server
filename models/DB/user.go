package DB

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"linktree_server/utils/logger"
	"time"
)

// User 用户表
type User struct {
	ID          uuid.UUID  `gorm:"primaryKey"`
	Name   		string
	Tel  		uint
	CreateTime 	time.Time
}

func AddUser(user *User){
	DB.Create(user)
}

func FindUser(){
	var user User
	DB.Model(&User{}).First(&user)
	logger.Log().Info(fmt.Sprintf("%v",user))
}