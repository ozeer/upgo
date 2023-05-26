package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"sort"
	"strings"

	"github.com/ozeer/upgo/global"
	"github.com/spf13/viper"
)

// 初始化UpGo程序安装
func InitUpGo(installDir string) {
	PrintMagenta("==> Init UpGo")

	// 仓库所有者的用户名或组织名
	owner := viper.GetString("github.owner")
	// 仓库的名称
	repo := viper.GetString("github.repo")

	// 使用GitHub API获取仓库的最新发布信息
	releaseURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", owner, repo)
	resp, err := http.Get(releaseURL)
	if err != nil {
		global.Error(fmt.Sprintf("无法获取最新发布信息：%s", err.Error()))
		return
	}
	defer resp.Body.Close()

	// 解析API响应获取下载链接
	var release struct {
		Assets []struct {
			Name               string `json:"name"`
			BrowserDownloadURL string `json:"browser_download_url"`
		} `json:"assets"`
	}
	err = json.NewDecoder(resp.Body).Decode(&release)
	if err != nil {
		global.Error(fmt.Sprintf("无法解析最新发布信息：%s", err.Error()))
		return
	}

	// 根据文件名称排序，以便选择要下载的文件
	sort.Slice(release.Assets, func(i, j int) bool {
		return release.Assets[i].Name < release.Assets[j].Name
	})

	// 获取最新发布的文件下载链接
	latestAsset := release.Assets[len(release.Assets)-1]
	downloadURL := latestAsset.BrowserDownloadURL

	// 下载文件
	finish := Download(downloadURL, latestAsset.Name)

	if finish {
		// 分配可执行权限
		Command("chmod +x upgo")

		// 移动下载的可执行文件到指定目录
		execRes := Command("mv upgo " + installDir)

		// 如果执行成功，删除可执行文件
		if execRes {
			PrintMagenta("==> UpGo初始化安装成功!安装目录：" + installDir)
		}
	}
}

// 检查用户输入的目录是否有效
func CheckInputDirIsValid(dirPath string) bool {
	// 检查目录是否有效
	info, err := os.Stat(dirPath)
	if err != nil {
		if os.IsNotExist(err) {
			global.Error("目录不存在: " + dirPath)
		} else {
			global.Error(fmt.Sprintf("发生错误：%s", err.Error()))
		}
		return false
	}

	if !info.IsDir() {
		global.Error("输入的路径不是一个目录: " + dirPath)
		return false
	}

	return true
}

// 检查Golang环境是否正常配置
func IsGoEnvConfigured() bool {
	cmd := exec.Command("go", "version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		global.Error("检查Golang环境出现错误: " + err.Error())
		return false
	}

	goVersion := string(output)
	return strings.Contains(goVersion, "go")
}
