package main

import (
	"fmt"
	"testing"
	"upgo/version"
)

func TestGetVersion(t *testing.T) {
	t.Logf("Latest version: %s", getLatestVersion())
}

func TestDownloadFile(t *testing.T) {
	downloaded("https://go.dev/dl/go1.20.4.darwin-amd64.tar.gz")
}

func TestInstall(t *testing.T) {
	install("go1.20.4.darwin-amd64.tar.gz")
}

func TestCommand(t *testing.T) {
	// shell sudo tar -C /usr/local -xzf /Users/zhouyang/web3/auto-upgrade-go/go1.20.4.darwin-amd64.tar.gz
	Command("tar -xzf go1.20.4.darwin-amd64.tar.gz")
}

func TestVersion(t *testing.T) {
	versions := []struct{ a, b string }{
		{"1.05.00.0156", "1.0.221.9289"},
		// Go versions
		{"1", "1.0.1"},
		{"1.0.1", "1.0.2"},
		{"1.0.2", "1.0.3"},
		{"1.0.3", "1.1"},
		{"1.1", "1.1.1"},
		{"1.1.1", "1.1.2"},
		{"1.1.2", "1.2"},
	}
	for _, stVersion := range versions {
		res := version.CompareVersionWithCache(stVersion.a, stVersion.b)
		switch {
		case res == 1:
			fmt.Println(stVersion.a, ">", stVersion.b)
		case res == -1:
			fmt.Println(stVersion.a, "<", stVersion.b)
		case res == 0:
			fmt.Println(stVersion.a, "=", stVersion.b)
		}
	}
}
