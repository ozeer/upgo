package main

import (
	"testing"

	"github.com/ozeer/upgo/service"
)

const TEST_GO_FILE = "go1.20.4.darwin-amd64.tar.gz"

func TestGetLatestVersion(t *testing.T) {
	t.Logf("Latest version: %s", service.GetLatestVersionFromApiSimple())
}

func TestGetCurrentVersion(t *testing.T) {
	t.Logf("Current version: %s", service.GetCurrentGoVersion())
}

func TestDownloadFile(t *testing.T) {
	service.Download(service.GO_DOWNLOAD_WEB+TEST_GO_FILE, TEST_GO_FILE)
}

func TestCommand(t *testing.T) {
	service.Command("go version")
}

func TestTopStableVersion(t *testing.T) {
	service.TopStableVersion()
}

func TestInitUpGo(t *testing.T) {
	service.InitUpGo()
}
