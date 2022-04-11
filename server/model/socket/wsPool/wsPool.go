package wsPool

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

/*
	一个websocket的管理池子
*/


type Ws struct {
	Conn     *websocket.Conn
	LinkTime int
	Ip       string
}

// Pool 使用上的map的参数为用户登录的token作为识别,Ws是创建的参数
type Pool map[string]Ws

// CreatePool 创建一个池
func CreatePool(linkSize int) Pool {
	wsPool := make(Pool, linkSize)
	return wsPool
}

// GetWs 通过查找map表获取ws
func (p Pool) GetWs(wsName string) (Ws, bool) {
	ws, boo := p[wsName]
	return ws, boo
}

// AddWs 添加ws连接
func (p Pool) AddWs(name string, ws Ws) bool {
	p[name] = ws
	msg := fmt.Sprintf("%v:连接成功!", name)
	err := ws.Conn.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		return false
	}
	// 在协程里面创建监听
	go onMessage(name, ws.Conn, p)
	return true
}

// CleanWs 删掉池内连接对象
func (p Pool) CleanWs(wsName string) {
	delete(p, wsName)
}

// WriteOne 发给单个
func (p Pool) WriteOne(wsName string, state int, msg interface{}) {
	// 获取发送用户的连接句柄
	ws, boo := p.GetWs(wsName)
	// 判断发送的数据类型
	switch state {
	case Text:
		if boo {
			err := ws.Conn.WriteMessage(websocket.TextMessage, []byte(msg.(string)))
			if err != nil {
				fmt.Printf("%v发送错误", wsName)
				return
			}
		} else {
			fmt.Println("连接已丢失或未连接")
		}
	case Json:
		if boo {
			err := ws.Conn.WriteJSON(msg)
			if err != nil {
				fmt.Printf("%v发送错误", wsName)
				return
			}
		} else {
			fmt.Println("连接已丢失或未连接")
		}
	}
}

// WriteMore 发给多个[list]用户列表
func (p Pool) WriteMore(userList []string, state int, msg interface{}) {
	// 请求提交多个user
	for userName := range userList {
		p.WriteOne(userList[userName], state, msg)
	}
}

// Broadcast 全局广播
func (p Pool) Broadcast(state int, msg interface{}) {
	for s := range p {
		p.WriteOne(s, state, msg)
	}
}

// GetPoolList 返回已连接池的列表
func (p Pool) GetPoolList() Pool {
	return p
}

// onMessage 收到消息的回调
func onMessage(name string, conn *websocket.Conn, pool Pool) {
	for {
		// 读取数据
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Printf("%v:的连接中断监听关闭", name)
			// 在这里删除池内对
			pool.CleanWs(name)
			return
		}
		// 返回原数据
		if err1 := conn.WriteMessage(messageType, p); err1 != nil {
			log.Println(err1)
			return
		}
		fmt.Printf("%v : %v\n", name, string(p))
	}
}
