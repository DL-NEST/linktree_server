package socket

import (
	"linktree_server/models/socket/wsPool"
)

var WsPool = make(wsPool.Pool, 2)

//func InitWsPool(num int) {
//	logger.Log().Info("设置websocket连接池")
//	logger.Log().Info(fmt.Sprintf("websocket池的最大值: %v",num))
//	WsPool =
//}
