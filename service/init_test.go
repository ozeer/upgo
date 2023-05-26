package service

import (
	"testing"

	"github.com/ozeer/upgo/conf"
)

// 初始化UpGo配置
func TestInitUpGo(t *testing.T) {
	conf.InitConfig()
	InitUpGo("/usr/local/bin")
}
