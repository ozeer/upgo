package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"
	"upgo/global"

	"github.com/Masterminds/semver/v3"
	"github.com/PuerkitoBio/goquery"
)

type Version struct {
	Version string `json:"version"`
}

const DEFAULT_GOLANG_VERSION = "go0"

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

// 查询所有稳定版本的Golang
// https://go.dev/dl/?mode=json&include=all
func AllStableVersion() {

}
