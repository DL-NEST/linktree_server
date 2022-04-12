package control

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	sock "linuxNet/model/socket"
	"linuxNet/model/socket/wsPool"
	"net/http"
)

type WsServer struct {
	upgrade *websocket.Upgrader
}

func checkOrigin(re *http.Request) bool {
	//fmt.Println(re)
	return true
}

// UpgradeSocket 升级请求为websocket
func UpgradeSocket(c *gin.Context) {
	// 获取请求的Token
	token := c.GetHeader("Token")
	ws := new(WsServer)
	ws.upgrade = &websocket.Upgrader{
		ReadBufferSize:  2048,
		WriteBufferSize: 2048,
		CheckOrigin:     checkOrigin,
	}
	Conn, err := ws.upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Printf("%s连接失败\n", token)
		return
	}
	// 添加到连接池
	// TODO 连接池溢出错误管理
	sock.WsPool.AddWs(token, wsPool.Ws{
		Conn:     Conn,
		LinkTime: 0,
		Ip:       "",
	})
}
