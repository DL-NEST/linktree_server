package main

import (
	"fmt"
	"linuxNet/bootstrap"
	"linuxNet/server/model/emqx"
	"linuxNet/server/router"
	"time"
	//"github.com/tensorflow/tensorflow/tensorflow/go"
)

func init() {
	// 比main先加载
	bootstrap.InitApp()
	bootstrap.InitConfig("config/config.yaml")
	// 初始化socket连接池
	//socket.InitWsPool(10)
	// 初始化mqtt
	emqx.InitMqtt()
}

func main() {
	//red.SetKey("zz","ssf",1*time.Minute)
	//val,err := red.GetKey("ji")
	//if err {
	//	fmt.Print(val)
	//}else {
	//	fmt.Print("df")
	//}
	//time.Sleep(time.Second * 1) // 暂停一秒等待 subscribe 完成
	////emqx.Publish(mqt,0)
	//fmt.Println(utils.GetUUID())

	// 初始化MQTT服务器
	// 初始化全局对象,参数
	router.InitRouter()
}

func goService() {
	fmt.Printf("协程测试\n")
	done := make(chan string)
	// 更多的子协程
	for i := 0; i < 10; i++ {
		go child(done, i)
	}
	done <- "1"
	done <- "2"
	done <- "3"
	done <- "4"
}

func child(done chan string, num int) {
	data := <-done
	fmt.Printf("child%d:%v\n", num, data)
	time.Sleep(4 * time.Second)
	// 继续改变这个参数会让还没有运行的子线程接收到然后完成打印,直到消耗掉所以子线程
	done <- "回去"
}
