package main

import (
	"testing"
	"upgo/service"
)

func TestGetVersion(t *testing.T) {
	t.Logf("Latest version: %s", service.GetLatestVersionFromApi())
}

func TestDownloadFile(t *testing.T) {
	service.Downloaded("https://go.dev/dl/go1.20.4.darwin-amd64.tar.gz")
}

func TestInstall(t *testing.T) {
	service.Install("go1.20.4.darwin-amd64.tar.gz")
}

func TestCommand(t *testing.T) {
	// shell sudo tar -C /usr/local -xzf /Users/zhouyang/web3/auto-upgrade-go/go1.20.4.darwin-amd64.tar.gz
	service.Command("tar -xzf go1.20.4.darwin-amd64.tar.gz")
}
