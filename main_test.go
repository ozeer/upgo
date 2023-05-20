package main

import (
	"testing"
)

func TestGetVersion(t *testing.T) {
	t.Logf("Latest version: %s", GetLatestVersion())
}

func TestDownloadFile(t *testing.T) {
	Downloaded("https://go.dev/dl/go1.20.4.darwin-amd64.tar.gz")
}

func TestInstall(t *testing.T) {
	install("go1.20.4.darwin-amd64.tar.gz")
}

func TestCommand(t *testing.T) {
	// shell sudo tar -C /usr/local -xzf /Users/zhouyang/web3/auto-upgrade-go/go1.20.4.darwin-amd64.tar.gz
	command("tar -xzf go1.20.4.darwin-amd64.tar.gz")
}
