package logger

import (
	"fmt"
	"github.com/fatih/color"
	"net/http"
	"time"
)

const (
	green   = "\033[97;42m"
	white   = "\033[90;47m"
	yellow  = "\033[90;43m"
	red     = "\033[97;41m"
	blue    = "\033[97;44m"
	magenta = "\033[97;45m"
	cyan    = "\033[97;46m"
	reset   = "\033[0m"
)

type GinMsg struct {
	Status int
	Proto  string
	Host   string
	Method string
	Path   string
}

// StatusCodeColor code状态的颜色
func (msg GinMsg) StatusCodeColor() string {
	code := msg.Status
	var StatusColor string
	switch {
	case code >= http.StatusOK && code < http.StatusMultipleChoices:
		StatusColor = green
	case code >= http.StatusMultipleChoices && code < http.StatusBadRequest:
		StatusColor = white
	case code >= http.StatusBadRequest && code < http.StatusInternalServerError:
		StatusColor = yellow
	default:
		StatusColor = red
	}
	return fmt.Sprintf("%v %d %v", StatusColor, msg.Status, reset)
}

// MethodColor 请求方式的颜色
func (msg GinMsg) MethodColor() string {
	method := msg.Method
	var StatusColor string
	switch method {
	case http.MethodGet:
		StatusColor = blue
	case http.MethodPost:
		StatusColor = cyan
	case http.MethodPut:
		StatusColor = yellow
	case http.MethodDelete:
		StatusColor = red
	case http.MethodPatch:
		StatusColor = green
	case http.MethodHead:
		StatusColor = magenta
	case http.MethodOptions:
		StatusColor = white
	default:
		StatusColor = reset
	}
	return fmt.Sprintf("%v %v     %v", StatusColor, msg.Method, reset)
}

// GinLog GIN日记
func (log *Logger) GinLog(ms GinMsg) {
	// 发送完成后解锁
	defer log.lock.Unlock()
	color.NoColor = false
	// 发送的时候锁定
	log.lock.Lock()

	pretext := fmt.Sprintf(
		"%s%s %s ",
		colors["Gin"]("[Gin]"),
		spaces["Gin"],
		time.Now().Format("2006-01-02 15:04:05"),
	)
	msg := fmt.Sprintf(
		" |%v|\t\t %v |\t\t %v | %v \"%v\" \n",
		ms.StatusCodeColor(),
		ms.Proto,
		ms.Host,
		ms.MethodColor(),
		ms.Path,
	)
	fmt.Printf(pretext + msg)
}
