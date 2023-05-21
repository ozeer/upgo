package service

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
)

var base_url = "https://go.dev/dl/"
var fileName string
var goInstallDir = "/usr/local"
var ch = make(chan struct{})

// 开始
func Start() {
	latestVersionGo := GetLatestVersionFromApi()
	currentVersionGo := GetCurrentGoVersion()

	latestVersion := latestVersionGo[2:]
	currentVersion := currentVersionGo[2:]

	// 如果未安装golang，提示语
	if currentVersion == DEFAULT_GOLANG_VERSION {
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

	// https://go.dev/dl/go1.20.4.darwin-amd64.tar.gz
	fileName = "go" + latestVersion + ".darwin-amd64.tar.gz"

	file, _ := PathExists(fileName)
	if !file {
		downloadUrl := base_url + fileName
		go func() {
			download := Downloaded(downloadUrl)
			if download {
				ch <- struct{}{}
			}
		}()

		<-ch
	}

	Install(fileName)
}

// 安装Golang
func Install(fileName string) bool {
	fmt.Println("==> Installing golang...")

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
		fmt.Println(err)
		os.Exit(1)
		return false
	}
	fmt.Println("==> Running: ", cmd)
	return true
}