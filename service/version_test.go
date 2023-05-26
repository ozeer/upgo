package service

import "testing"

func TestGetLatestVersion(t *testing.T) {
	t.Logf("Latest version: %s", GetLatestVersionFromApiSimple())
}

func TestGetCurrentVersion(t *testing.T) {
	t.Logf("Current version: %s", GetCurrentGoVersion())
}

func TestTopStableVersion(t *testing.T) {
	TopStableVersion()
}
func TestGetUpGoLatestVersionTag(t *testing.T) {
	t.Logf("Latest tag: %s", GetUpGoLatestVersionTag())
}
