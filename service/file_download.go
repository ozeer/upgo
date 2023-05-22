package service

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"upgo/global"
)

const DOWNLOAD_WEB = "https://go.dev/dl/"

// 根据给定url下载文件
func Downloaded(stFileName string) bool {
	// Build fileName from fullPath
	fullUrlFile := DOWNLOAD_WEB + stFileName
	fileURL, err := url.Parse(fullUrlFile)
	if err != nil {
		global.Error(fmt.Sprintf("url parse error: %s", err.Error()))
	}
	path := fileURL.Path
	segments := strings.Split(path, "/")
	fileName := segments[len(segments)-1]

	// Create blank file
	file, err := os.Create(fileName)
	if err != nil {
		global.Error(fmt.Sprintf("create file error: %s", err.Error()))
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	// Put content on file
	resp, err := client.Get(fullUrlFile)
	if err != nil {
		global.Error(fmt.Sprintf("get file error: %s", err.Error()))
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	if err != nil {
		global.Error(fmt.Sprintf("copy data error: %s", err.Error()))
	}

	defer file.Close()

	fileSize := fmt.Sprintf("%.2fMB", float64(size)/float64(1024*1024))
	PrintMagenta(fmt.Sprintf("==> Downloading %s（%s）", fullUrlFile, fileSize))

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
