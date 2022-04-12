package logger

import (
	"fmt"
	"strconv"
)

//前景色   背景色
//30  	40	  黑色
//31  	41	  红色
//32  	42	  绿色
//33  	43    黄色
//34  	44    蓝色
//35  	45 	  紫色
//36  	46 	  青色
//37  	47	  白色

func Success(txt string) {
	fmt.Printf("\x1B[42;30m Success \x1B[;32m %v\x1B[0m\n", txt)
}
func Info(txt string) {
	fmt.Printf("\x1B[44;30m INFO \x1B[;m %v\x1B[0m\n", txt)
}
func Error(txt string) {
	fmt.Printf("\x1B[41;30m ERROR \x1B[;31m %v\x1B[0m\n", txt)
}
func Warn(txt string) {
	fmt.Printf("\x1B[43;30m WARN \x1B[;33m %v\x1B[0m\n", txt)
}
func Clear() {
	fmt.Printf("\x1B[0m\x1B[0m")
}

func f2i(f float32) int {
	i, _ := strconv.Atoi(fmt.Sprintf("%1.0f", f))
	return i
}
func bar(now float32, total float32, bf float32) string {
	var bar string
	for i := now; i < 50; i++ {
		bar += "="
	}
	return bar
}
func Bar(txt string, now float32, total float32) {
	bf := now / total * 100
	fmt.Printf("\x1B[46;30m ProGressBar \x1B[;36m %v: %0.1f%% 处理的文件 [%v] %v/%v\x1B[0m",
		txt, bf, bar(now, total, bf/2), now, total)
}
