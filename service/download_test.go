package service

import "testing"

const TEST_GO_FILE = "go1.20.4.darwin-amd64.tar.gz"

func TestDownload(t *testing.T) {
	Download(GO_DOWNLOAD_WEB+TEST_GO_FILE, TEST_GO_FILE)
}
