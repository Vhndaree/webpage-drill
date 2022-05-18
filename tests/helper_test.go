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
}
