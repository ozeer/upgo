package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
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

	// 默认只选取固定的.darwin_amd64.tar.gz文件下载
	var downloadURL = ""
	var fileName = ""
	for _, v := range release.Assets {
		if strings.Contains(v.Name, "darwin_amd64.tar.gz") {
			downloadURL = v.BrowserDownloadURL
			fileName = v.Name
		}
	}

	// 下载文件
	finish := Download(downloadURL, fileName)

	if finish {
		// 解压压缩文件并进入文件夹
		// unzipFolder := strings.TrimSuffix(fileName, ".tar.gz")
		// Command("mkdir " + unzipFolder)
		Command("tar -xzf " + fileName)

		// 分配可执行权限
		Command("chmod +x upgo")

		// 移动下载的可执行文件到指定目录
		execRes := Command("mv upgo " + installDir)

		// 如果执行成功，删除可执行文件
		if execRes {
			Command("rm " + fileName)
			Command("rm LICENSE README.md")
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
