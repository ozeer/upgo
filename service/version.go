package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"

	"github.com/ozeer/upgo/global"

	"github.com/Masterminds/semver/v3"
	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
)

type Version struct {
	Version string `json:"version"`
}

// 未安装golang时的标识
const DEFAULT_GOLANG_VERSION = "go0"
const UP_GO_INITIAL_VERSION = "1.0.0"

// 方式一：通过解析Go官方网页获取最新稳定版本Golang编号
func GetLatestVersionFromHtml() string {
	res, err := http.Get(base_url)
	if err != nil {
		global.Error(fmt.Sprintf("new document from reader error: %s", err.Error()))
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		global.Error(fmt.Sprintf("status code error: %d %s", res.StatusCode, res.Status))
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		global.Error(fmt.Sprintf("new document from reader error: %s", err.Error()))
	}

	latestVersion := doc.Find("div.collapsed").Find("span").Eq(0).Text()

	return latestVersion
}

// 方式二：通过Go官方的接口获取最新稳定版本Golang编号
func GetLatestVersionFromApi() string {
	resp, err := http.Get("https://go.dev/dl/?mode=json&include=stable")
	if err != nil {
		global.Error(fmt.Sprintf("error fetching version: %s", err.Error()))
	}
	defer resp.Body.Close()

	var versions []Version
	err = json.NewDecoder(resp.Body).Decode(&versions)
	if err != nil {
		global.Error(fmt.Sprintf("error decoding JSON: %s", err.Error()))
	}

	latestVersion := ""
	if len(versions) > 0 {
		latestVersion = versions[0].Version
	} else {
		global.Error("no stable Go versions found.")
	}

	return latestVersion
}

// 方式三：使用官方更精简的接口获取最新稳定版本Golang编号
func GetLatestVersionFromApiSimple() string {
	resp, err := http.Get("https://go.dev/VERSION?m=text")
	if err != nil {
		global.Error(fmt.Sprintf("error fetching version: %s", err.Error()))
	}
	defer resp.Body.Close()

	latestVersion := ""
	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			global.Error(fmt.Sprintf("error reading response: %s", err.Error()))
		}

		latestVersion = string(body)
	} else {
		global.Error(fmt.Sprintf("Error fetching version:  %s", resp.Status))
	}

	return latestVersion
}

// 获取当前机器安装的golang版本
func GetCurrentGoVersion() string {
	cmd := exec.Command("go", "version")
	output, err := cmd.Output()
	if err != nil {
		// 无法获取Go版本
		return DEFAULT_GOLANG_VERSION
	}

	version := strings.Split(string(output), " ")[2]
	return version
}

// 判断是否有新版本
func HasNewVersion(latest, current string) bool {
	latestVersion, _ := semver.NewVersion(latest)
	currentVersion, _ := semver.NewVersion(current)

	return latestVersion.GreaterThan(currentVersion)
}

// 查询最近15个最新稳定版本的Golang
// https://go.dev/dl/?mode=json&include=all
func TopStableVersion() {
	resp, err := http.Get("https://go.dev/dl/?mode=json&include=all")
	if err != nil {
		global.Error(fmt.Sprintf("error fetching version: %s", err.Error()))
	}
	defer resp.Body.Close()

	var versions []Version
	err = json.NewDecoder(resp.Body).Decode(&versions)
	if err != nil {
		global.Error(fmt.Sprintf("error decoding JSON: %s", err.Error()))
	}

	if len(versions) > 0 {
		yellow := color.New(color.FgYellow).SprintFunc()
		fmt.Println(yellow("Top 15 available stable versions: "))

		topTenVersions := versions[:15]
		PrintFixedColumnVersion(topTenVersions)
	} else {
		global.Error("no stable Go versions found.")
	}
}

// 检查当前最新版本的golang
func CheckNewestVersion() {
	latestVersionGo := GetLatestVersionFromApiSimple()
	currentVersionGo := GetCurrentGoVersion()

	latestVersion := latestVersionGo[2:]
	currentVersion := currentVersionGo[2:]

	// 如果未安装golang，提示语
	text := ""
	if currentVersion == "0" {
		red := color.New(color.FgRed).SprintFunc()
		text = red("Golang is not installed locally!")
	} else {
		if HasNewVersion(latestVersion, currentVersion) {
			magenta := color.New(color.FgMagenta).SprintFunc()
			text = fmt.Sprintf("New version available for update: %s  ->   %s\n", magenta(currentVersion), magenta(latestVersion))
		} else {
			text = fmt.Sprintf("You are already using the latest available Golang version %s (stable).", latestVersion)
		}
	}
	color.Cyan(text)
}

// 检查版本号格式是否有效
func IsValidVersion(version string) bool {
	_, err := semver.NewVersion(version)
	if err != nil {
		global.Error(fmt.Sprintf("Error parsing version: %s", err.Error()))
		return false
	}

	return true
}

// 升级UpGo到最新版本
func SelfUpdate() {
	magenta := color.New(color.FgMagenta).SprintFunc()
	fmt.Println(magenta("==> UpGo self updating..."))
	res := Command("go install github.com/ozeer/upgo@latest")

	if res {
		fmt.Println(magenta("==> Update succeeded!"))
	} else {
		red := color.New(color.FgRed).SprintFunc()
		text := red("==> Update fail!")
		fmt.Println(text)
	}
}

// 获取UpGo最新版本的标签值
func GetUpGoLatestVersionTag() string {
	cmd := exec.Command("git", "describe", "--tags", "--abbrev=0")
	output, err := cmd.Output()
	if err != nil {
		global.Error("GetUpGoLatestVersionTag error: " + err.Error())
		return UP_GO_INITIAL_VERSION
	}

	tag := strings.TrimSpace(string(output))

	return tag
}
