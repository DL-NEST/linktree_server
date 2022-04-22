package DB

import (
	"fmt"
	"github.com/fatih/color"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"linktree_server/bootstrap"
	"linktree_server/utils/logger"
)


var DB *gorm.DB

// CreateDBLink 创建数据库连接
func CreateDBLink() {
	// 判断数据库的使用类型 sqlite and mysql
	if GetDbMode() {
		// 使用mysql
		logger.Log().Info("连接数据库 "+color.New(color.FgBlue).Add(color.Bold).Sprint(":mysql"))
		DB = LinkMySql()
	}else {
		// 使用sqlite
		logger.Log().Info("连接数据库 "+color.New(color.FgBlue).Add(color.Bold).Sprint(":sqlite"))
		DB = LinkSqlLite("param/linkTree.sqlite")
	}
}

// GetDbMode 判断使用的数据库类型
func GetDbMode() bool {
	if bootstrap.GetDbConf() {
		return true
	}else {
		return false
	}
}

type Mysql struct {
	UserName    string
	Password    string
	Host        string
	Port        string
	soa         string
	Charset     string
}

func LinkMySql() *gorm.DB {
	_ = "linktree:dl2002123@tcp(1.14.96.5:3306)/linktree?charset=utf8"
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		"linktree",
		"dl2002123",
		"1.14.96.5",
		"3306",
		"linktree",
		"utf8")
	// 打开数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		CreateBatchSize: 1000,
	})
	if err != nil {
		panic(any("连接数据库失败"))
	}

	DB = db
	// 数据库迁徙
	migration()

	return db
}

func LinkSqlLite(path string) *gorm.DB {
	// 打开数据库
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{
		CreateBatchSize: 1000,
	})
	if err != nil {
		panic(any("连接数据库失败"))
	}

	DB = db
	// 数据库迁徙
	migration()

	return db
}
