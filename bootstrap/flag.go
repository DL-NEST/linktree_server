package bootstrap

import (
	"flag"
	"linktree_server/utils/logger"
	"os"
)

var GlobalConsole Console

// Console /*
type Console struct {
	Conf string
	Log  string
}

func InitFlag() {
	sysFlag := flag.NewFlagSet("linkTree", flag.ExitOnError)
	sysFlag.StringVar(&GlobalConsole.Conf, "conf", "null", "配置文件地址")
	sysFlag.StringVar(&GlobalConsole.Log, "log", " null", "日记的输出路径")
	// 解析
	err := sysFlag.Parse(os.Args[1:])
	if err != nil {
		return
	}
}

func (co Console) OutFlag() {
	logger.Log().Flag("[\u001B[;34m -conf \u001B[0m%v,\u001B[;34m -log \u001B[0m %v]", co.Conf, co.Log)
}
