package logger

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
	"sync"
	"time"
)

const (
	// LevelError 错误
	LevelError = iota
	// LevelWarning 警告
	LevelWarning
	// LevelInformational 提示
	LevelInformational
	// LevelDebug 除错
	LevelDebug
)

var GlobalLogger *Logger
var Level = LevelDebug

type Logger struct {
	level int
	// 互斥量
	lock sync.Mutex
}

// 日志颜色
var colors = map[string]func(a ...interface{}) string{
	"Warning": color.New(color.FgYellow).Add(color.Bold).SprintFunc(),
	"Panic":   color.New(color.BgRed).Add(color.Bold).SprintFunc(),
	"Error":   color.New(color.FgRed).Add(color.Bold).SprintFunc(),
	"Info":    color.New(color.FgCyan).Add(color.Bold).SprintFunc(),
	"Debug":   color.New(color.FgHiMagenta).Add(color.Bold).SprintFunc(),
	"Gin":     color.New(color.FgHiCyan).Add(color.Bold).SprintFunc(),
	"Flag":    color.New(color.FgYellow).Add(color.Bold).SprintFunc(),
	"Socket":  color.New(color.FgHiCyan).Add(color.Bold).SprintFunc(),
	"Mqtt":    color.New(color.FgHiCyan).Add(color.Bold).SprintFunc(),
}

// 不同级别前缀与时间的间隔，保持宽度一致
var spaces = map[string]string{
	"Warning": "",
	"Panic":   "  ",
	"Error":   "  ",
	"Info":    "   ",
	"Debug":   "  ",
	"Gin":     "    ",
	"Flag":    "   ",
	"Socket":  " ",
	"Mqtt":    "   ",
}

func logFile(path string) *os.File {
	src, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModeAppend)
	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}
	return src
}

// Println 打印
func (log *Logger) Println(prefix string, msg string) {
	// 发送完成后解锁
	defer log.lock.Unlock()
	color.NoColor = false
	// 发送的时候锁定
	log.lock.Lock()

	_ = fmt.Sprintf(
		"%s%s %s %s\n",
		"["+prefix+"]",
		spaces[prefix],
		time.Now().Format("2006-01-02 15:04:05"),
		msg,
	)
	ctrlOut := fmt.Sprintf(
		"%s%s %s %s\n",
		colors[prefix]("["+prefix+"]"),
		spaces[prefix],
		time.Now().Format("2006-01-02 15:04:05"),
		msg,
	)
	//_, err := logFile(ServerPath).WriteString(logOut)
	//if err != nil {
	//	return
	//}
	fmt.Printf(ctrlOut)
}

// Panic 极端错误
func (log *Logger) Panic(format string, v ...interface{}) {
	if LevelError > log.level {
		return
	}
	msg := fmt.Sprintf(format, v...)
	log.Println("Panic", msg)
	//panic(any(msg))
}

// Error 错误
func (log *Logger) Error(format string, v ...interface{}) {
	if LevelError > log.level {
		return
	}
	msg := fmt.Sprintf(format, v...)
	log.Println("Error", msg)
}

// Warning 警告
func (log *Logger) Warning(format string, v ...interface{}) {
	if LevelWarning > log.level {
		return
	}
	msg := fmt.Sprintf(format, v...)
	log.Println("Warning", msg)
}

// Info 信息
func (log *Logger) Info(format string, v ...interface{}) {
	if LevelInformational > log.level {
		return
	}
	msg := fmt.Sprintf(format, v...)
	log.Println("Info", msg)
}

// Debug 校验
func (log *Logger) Debug(format string, v ...interface{}) {
	if LevelDebug > log.level {
		return
	}
	msg := fmt.Sprintf(format, v...)
	log.Println("Debug", msg)
}

// Flag 控制台输入
func (log *Logger) Flag(format string, v ...interface{}) {
	if LevelError > log.level {
		return
	}
	msg := fmt.Sprintf(format, v...)
	log.Println("Flag", msg)
}
// Socket 控制台输入
func (log *Logger) Socket(format string, v ...interface{}) {
	if LevelError > log.level {
		return
	}
	msg := fmt.Sprintf(format, v...)
	log.Println("Socket", msg)
}
// Mqtt 控制台输入
func (log *Logger) Mqtt(format string, v ...interface{}) {
	if LevelError > log.level {
		return
	}
	msg := fmt.Sprintf(format, v...)
	log.Println("Mqtt", msg)
}

// Log 返回日志对象
func Log() *Logger {
	if GlobalLogger == nil {
		l := Logger{
			level: Level,
		}
		GlobalLogger = &l
	}
	return GlobalLogger
}

func Debug(format string, v ...interface{})  {
	msg := fmt.Sprintf(format, v...)
	Log().Debug(msg)
}