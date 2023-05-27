package main

import (
	"github.com/ozeer/upgo/cmd"
	"github.com/ozeer/upgo/conf"
	"github.com/ozeer/upgo/global"
)

func main() {
	// 初始化配置
	conf.InitConfig()

	// 初始化日志
	global.Logger = conf.InitLogger()
	defer func() {
		// nolint:errcheck
		_ = global.Logger.Sync()
	}()

	// 启动Cobra命令行
	cmd.Execute()
}
