package main

import (
	"log"

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
		err := global.Logger.Sync()
		if err != nil {
			// 处理错误的逻辑
			log.Println("日志错误：", err.Error())
		}
	}()

	// 启动Cobra命令行
	cmd.Execute()
}
