package DB

import (
	uuid "github.com/satori/go.uuid"
	"testing"
	"time"
)

func TestLinkSqlLite(t *testing.T) {
	db := LinkSqlLite("../../param/linkTree.sqlite")
	// 插入值
	db.Create(&User{
		ID:         0,
		Name:       "root",
		Tel:        13242351251,
		CreateTime: time.Now(),
		UUID: 		uuid.NewV4(),
	})
}


func TestLinkMySql(t *testing.T) {
	db := LinkMySql()
	// 插入值
	db.Create(&User{
		ID:         0,
		Name:       "root",
		Tel:        13242351251,
		CreateTime: time.Now(),
		UUID: 		uuid.NewV4(),
	})
}
