package socket

import (
	"linuxNet/model/socket/wsPool"
)

var WsPool wsPool.Pool

func InitWsPool(num int) {
	WsPool = wsPool.CreatePool(num)
}
