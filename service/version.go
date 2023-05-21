package service

import (
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/PuerkitoBio/goquery"
)

// 获取最新Golang的版本号
func GetLatestVersion() string {
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
func AllStableVersion() {

}
