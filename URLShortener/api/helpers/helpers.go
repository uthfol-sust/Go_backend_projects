package helpers

import (
	"os"
	"strings"
)

func EnforceHTTP(url string) string {
	if !strings.HasPrefix(url, "http") {
		return "http://" + url
	}
	return url
}

func RemoveDomainError(url string) bool {
	if url == os.Getenv("DOMAIN") {
		return false
	}

	u := strings.Replace(url, "http://", "", 1)
	u = strings.Replace(u, "https://", "", 1)
	u = strings.Replace(u, "www.", "", 1)
	u = strings.Split(u, "/")[0]

	if u == os.Getenv("DOMAIN") {
		return false
	}
	return true
}
