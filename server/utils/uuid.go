package utils

import (
	"github.com/satori/go.uuid"
)

func GetUUID() string {
	u1 := uuid.NewV4()
	return u1.String()
}
