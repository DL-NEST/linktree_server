package socket

import (
	"linktree_server/model/socket/wsPool"
)

var WsPool wsPool.Pool

func InitWsPool(num int) {
	WsPool = wsPool.CreatePool(num)
}
