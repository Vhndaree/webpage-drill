package util

import "strings"

func SanitizeURL(url string) string {
	if idx := strings.Index(url, "://"); idx != -1 {
		return url[(idx + 3):]
	}

	return url
}

func AddURLProtocol(url string) string {
	if !strings.Contains(url, "http") {
		return "https://" + url
	}

	return url
}
