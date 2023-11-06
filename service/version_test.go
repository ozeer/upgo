package service

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/ozeer/upgo/conf"
	"github.com/ozeer/upgo/global"
)

func TestGetLatestVersion(t *testing.T) {
	begin1 := time.Now()
	t.Logf("Latest version: %s", GetLatestVersionFromHtml())
	fmt.Printf("\ncost: %dms\n", time.Since(begin1).Milliseconds())

	begin2 := time.Now()
	t.Logf("Latest version: %s", GetLatestVersionFromApi())
	fmt.Printf("\ncost: %dms\n", time.Since(begin2).Milliseconds())

	begin3 := time.Now()
	t.Logf("Latest version: %s", GetLatestVersionFromApiSimple())
	fmt.Printf("\ncost: %dms\n", time.Since(begin3).Milliseconds())
}

func TestGetCurrentVersion(t *testing.T) {
	t.Logf("Current version: %s", GetCurrentGoVersion())
}

func TestTopStableVersion(t *testing.T) {
	TopStableVersion()
}
func TestGetUpGoLatestVersionTag(t *testing.T) {
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

	t.Logf("Latest tag: %s", GetUpGoLatestVersionTag())
}
