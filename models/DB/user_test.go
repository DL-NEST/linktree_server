package DB

import (
	uuid "github.com/satori/go.uuid"
	"testing"
	"time"
)

func TestAddUser(t *testing.T)  {
	LinkSqlLite("../../param/linkTree.sqlite")
	AddUser(&User{
		ID:         uuid.NewV4(),
		Name:       "root",
		Tel:        13242351251,
		CreateTime: time.Now(),
	})
}

func TestUpdateUser(t *testing.T) {
	LinkSqlLite("../../param/linkTree.sqlite")
	FindUser()
}