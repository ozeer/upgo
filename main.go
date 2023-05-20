package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"upgo/version"

	"github.com/PuerkitoBio/goquery"
)

var base_url = "https://go.dev/dl/"
var fileName string
var goInstallDir = "/usr/local"
var ch = make(chan struct{})

func main() {
	latestVersionGo := getLatestVersion()

	latestVersion := latestVersionGo[2:]
	currentVersion := runtime.Version()[2:]
	compareVersion := version.CompareVersionWithCache(latestVersion, currentVersion)

	if compareVersion < 1 {
		fmt.Printf("当前已经是最新版本!当前版本：%s，最新版本：%s\n", currentVersion, latestVersion)
		return
	}

	fmt.Printf("有新版本可更新!当前版本：%s，可更新版本：%s\n", currentVersion, latestVersion)

	// https://go.dev/dl/go1.20.4.darwin-amd64.tar.gz
	fileName = "go" + latestVersion + ".darwin-amd64.tar.gz"

	file, _ := PathExists(fileName)
	if !file {
		downloadUrl := base_url + fileName
		go func() {
			download := downloaded(downloadUrl)
			if download {
				ch <- struct{}{}
			}
		}()

		<-ch
	}

	install(fileName)
}

// 安装Golang
func install(fileName string) bool {
	fmt.Println("最新golang安装中...")

	// 删除老的golang
	deleteGoShell := "sudo rm -rf /usr/local/go"
	Command(deleteGoShell)

	shell := "sudo tar -C " + goInstallDir + " -xzf " + fileName
	result := Command(shell)
	if result {
		Command("rm " + fileName)
	}

	return result
}

// 获取最新Golang的版本号
func getLatestVersion() string {
	// Request the HTML page.
	res, err := http.Get(base_url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	latestVersion := doc.Find("div.collapsed").Find("span").Eq(0).Text()

	return latestVersion
}

// 根据给定url下载文件
func downloaded(fullURLFile string) bool {
	// Build fileName from fullPath
	fileURL, err := url.Parse(fullURLFile)
	if err != nil {
		log.Fatal(err)
	}
	path := fileURL.Path
	segments := strings.Split(path, "/")
	fileName := segments[len(segments)-1]

	// Create blank file
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	// Put content on file
	resp, err := client.Get(fullURLFile)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	fileSize := fmt.Sprintf("%.2fMB", float64(size)/float64(1024*1024))
	log := fmt.Sprintf("%s下载完成！大小%s", fileName, fileSize)
	fmt.Println(log)

	return true
}

// 判断文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
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
	fmt.Println("命令执行成功: ", cmd)
	return true
}
