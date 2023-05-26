package service

import (
	"testing"

	"github.com/ozeer/upgo/conf"
	"github.com/ozeer/upgo/global"
)

func TestGetLatestVersion(t *testing.T) {
	t.Logf("Latest version: %s", GetLatestVersionFromApiSimple())
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
	defer global.Logger.Sync()

	t.Logf("Latest tag: %s", GetUpGoLatestVersionTag())
}
