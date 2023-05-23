package service

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"upgo/global"

	"github.com/cheggaaa/pb/v3"
)

type proxyWriter struct {
	bar       *pb.ProgressBar
	writer    io.Writer
	writeFunc func(p []byte) (n int, err error)
}

const DOWNLOAD_WEB = "https://go.dev/dl/"

func (pw *proxyWriter) Write(p []byte) (n int, err error) {
	n, err = pw.writeFunc(p)
	if err == nil {
		pw.bar.Add(n)
	}
	return
}

// 下载文件
func Download(fileName string) bool {
	// 文件完整下载地址
	fileURL := DOWNLOAD_WEB + fileName

	// 创建 HTTP 客户端
	client := http.Client{
		// Timeout: 15 * time.Second,
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	// 发起 GET 请求
	resp, err := client.Get(fileURL)
	if err != nil {
		fmt.Printf("无法下载文件：%s\n", err)
		return false
	}
	defer resp.Body.Close()

	// 创建输出文件
	outputFile, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("无法创建输出文件：%s\n", err)
		return false
	}
	defer outputFile.Close()

	// 获取响应体的大小
	fileSize := resp.ContentLength

	// 创建进度条
	bar := pb.Full.Start64(fileSize)
	bar.Set(pb.Bytes, true)

	// 创建代理写入器
	proxyWriter := &proxyWriter{
		bar:    bar,
		writer: outputFile,
		writeFunc: func(p []byte) (n int, err error) {
			n, err = outputFile.Write(p)
			return
		},
	}

	// 将响应体复制到代理写入器
	_, err = io.Copy(proxyWriter, resp.Body)
	if err != nil {
		fmt.Printf("下载过程中出现错误：%s\n", err)
		return false
	}

	// 完成进度条
	bar.Finish()

	// 打印下载文案
	fileMBSize := fmt.Sprintf("%.2fMB", float64(fileSize)/float64(1024*1024))
	PrintMagenta(fmt.Sprintf("==> Download complete！！！ %s（%s）", fileURL, fileMBSize))

	return true
}

// 从下载文件的 URL中解析文件名
func ParseFileNameFromUrl(fileURL string) string {
	uri, err := url.Parse(fileURL)
	if err != nil {
		global.Error(fmt.Sprintf("url parse error: %s", err.Error()))
		return ""
	}
	path := uri.Path
	segments := strings.Split(path, "/")
	fileName := segments[len(segments)-1]

	return fileName
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
