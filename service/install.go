package service

import (
	"fmt"
	"os"
	"os/exec"
	"upgo/global"

	"github.com/fatih/color"
)

var base_url = "https://go.dev/dl/"
var goInstallDir = "/usr/local"
var ch = make(chan struct{})

// 开始
func Start() {
	latestVersionGo := GetLatestVersionFromApiSimple()
	currentVersionGo := GetCurrentGoVersion()

	latestVersion := latestVersionGo[2:]
	currentVersion := currentVersionGo[2:]

	// 如果未安装golang，提示语
	if currentVersion == "0" {
		color.Cyan("==> Preparing for installation...")
	} else {
		if !HasNewVersion(latestVersion, currentVersion) {
			text := fmt.Sprintf("You are already using the latest available Golang version %s (stable).", latestVersion)
			color.Cyan(text)
			return
		}

		color.Cyan("==> Upgrading...")
		fmt.Printf("go  %s  ->   %s\n", currentVersion, latestVersion)
	}

	// go1.20.4.darwin-amd64.tar.gz
	fileName := "go" + latestVersion + ".darwin-amd64.tar.gz"
	Install(fileName)
}

// 安装Golang
func Install(fileName string) bool {
	PrintMagenta("==> Installing golang...")

	file, _ := PathExists(fileName)
	if !file {
		go func() {
			download := Download(fileName)
			if download {
				ch <- struct{}{}
			}
		}()

		<-ch
	}

	// 删除老的golang
	deleteGoShell := "sudo rm -rf /usr/local/go"
	Command(deleteGoShell)

	shell := "sudo tar -C " + goInstallDir + " -xzf " + fileName
	result := Command(shell)
	if result {
		Command("rm " + fileName)
		color.Cyan("==> Congratulations! " + fileName + " Installed.")
	}

	return result
}

// 执行命令方法
func Command(cmd string) bool {
	c := exec.Command("/bin/sh", "-c", cmd)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Run()
	if err != nil {
		global.Error(fmt.Sprintf("command exec error: %s", err.Error()))
		os.Exit(1)
		return false
	}
	PrintMagenta("==> Running: " + cmd)
	return true
}

// 获取Go的安装目录
func GetGoRootDir() string {
	goRoot := os.Getenv("GOROOT")
	if goRoot == "" {
		global.Error("Failed to get Go installation directory.")
	}

	return goRoot
}
