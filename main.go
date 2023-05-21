package main

import (
	"upgo/cmd"
	"upgo/conf"
	"upgo/global"
)

func main() {
	// 初始化配置
	conf.InitConfig()

	// 初始化日志
	global.Logger = conf.InitLogger()
	defer global.Logger.Sync()

	// 启动Cobra命令行
	cmd.Execute()
}
