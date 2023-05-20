package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// 根据给定url下载文件
func Downloaded(fullUrlFile string) bool {
	// Build fileName from fullPath
	fileURL, err := url.Parse(fullUrlFile)
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
	resp, err := client.Get(fullUrlFile)
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
	fmt.Printf("==> Downloading %s（%s）\n", fullUrlFile, fileSize)

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
