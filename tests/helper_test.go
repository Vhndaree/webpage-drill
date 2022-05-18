package tests

import (
	"testing"
	"webpageinfo/util"
)

func TestHelper(t *testing.T) {
	t.Log("URL Sanitizer")
	{
		t.Run("SanitizeURL", func(t *testing.T) {
			ExpectEqual(t, util.SanitizeURL("https://app.go"), "app.go")
			ExpectEqual(t, util.SanitizeURL("http://app.go"), "app.go")
			ExpectEqual(t, util.SanitizeURL("app.go"), "app.go")
			ExpectEqual(t, util.SanitizeURL(""), "")
		})
	}

	t.Log("Adds http protocol")
	{
		t.Run("AddURLProtocol", func(t *testing.T) {
			ExpectEqual(t, util.AddURLProtocol("https://app.go"), "https://app.go")
			ExpectEqual(t, util.AddURLProtocol("http://app.go"), "http://app.go")
			ExpectEqual(t, util.AddURLProtocol("app.go"), "https://app.go")
		})
	}
}
