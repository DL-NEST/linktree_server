package device

/*
	设备连接池
*/

// DevPool 设备名称是map索引
type DevPool map[string]Dev

// Dev 设备
type Dev struct {
	Name 	string		// 设备名称
	Mqtt 	MqttList	// 设备订阅的主题列表
	Data  	string
}

type Devs interface {
	OnMsg()		// 消息的回调
	SendMsg() 	// 消息发送
}

type MqttList map[string]Mqtt

type Mqtt struct {
	Topic 	string	// 订阅主题
	Qos 	int
	Rules   string	// 消息解析规则
}

func LinkDev()  {

}




