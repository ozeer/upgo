package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/PuerkitoBio/goquery"
)

type Version struct {
	Version string `json:"version"`
}

const DEFAULT_GOLANG_VERSION = "0"

// 方式一：通过解析Go官方网页获取最新稳定版本Golang编号
func GetLatestVersionFromHtml() string {
	res, err := http.Get(base_url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	latestVersion := doc.Find("div.collapsed").Find("span").Eq(0).Text()

	return latestVersion
}

// 方式二：通过Go官方的接口获取最新稳定版本Golang编号
func GetLatestVersionFromApi() string {
	resp, err := http.Get("https://go.dev/dl/?mode=json&include=stable")
	if err != nil {
		fmt.Println("Error fetching version:", err)
		return DEFAULT_GOLANG_VERSION
	}
	defer resp.Body.Close()

	var versions []Version
	err = json.NewDecoder(resp.Body).Decode(&versions)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return DEFAULT_GOLANG_VERSION
	}

	latestVersion := DEFAULT_GOLANG_VERSION
	if len(versions) > 0 {
		latestVersion = versions[0].Version
	} else {
		fmt.Println("No stable Go versions found.")
	}

	return latestVersion
}

// 获取当前机器安装的golang版本
func GetCurrentGoVersion() string {
	cmd := exec.Command("go", "version")
	output, err := cmd.Output()
	if err != nil {
		// 无法获取Go版本
		return "go0"
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
