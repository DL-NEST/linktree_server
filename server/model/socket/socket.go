package socket

import "linuxNet/server/model/socket/wsPool"

var WsPool = wsPool.CreatePool(10)

//func InitWsPool(num int)  {
//	WsPool = wsPool.CreatePool(num)
//}