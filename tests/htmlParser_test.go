package tests

import (
	"testing"
	"webpageinfo/model"
	"webpageinfo/util"
)

func TestHTMLParser(t *testing.T) {
	t.Log("HTML Version Reader")
	{
		t.Run("Version", func(t *testing.T) {
			ExpectEqual(t, util.Version(dummyHTML1), "XHTML 1.0 Strict")
			ExpectEqual(t, util.Version(dummyHTML2), "HTML 3.2 Final")
			ExpectEqual(t, util.Version(dummyHTML3), "HTML 5")
		})
	}

	t.Log("HTML title Reader")
	{
		t.Run("Title", func(t *testing.T) {
			ExpectEqual(t, util.Title(dummyHTML1), "W3C QA - Recommended list of Doctype declarations you can use in your Web document")
			ExpectEqual(t, util.Title(dummyHTML2), "Home")
			ExpectEqual(t, util.Title(dummyHTML3), "Old site in the web")
		})
	}

	t.Log("HTML Count Headers")
	{
		t.Run("Headings", func(t *testing.T) {
			ExpectEqual(t, util.Headings(dummyHTML1), model.Headings{2, 9, 1, 1, 1, 1})
			ExpectEqual(t, util.Headings(dummyHTML2), model.Headings{2, 1, 1, 1, 1, 1})
			ExpectEqual(t, util.Headings(dummyHTML3), model.Headings{2, 1, 1, 1, 1, 1})
		})
	}

	t.Log("HTML Links Counter")
	{
		t.Run("Link", func(t *testing.T) {
			ExpectEqual(t, util.Links("https://www.w3.org/", dummyHTML1), model.Link{60, 6, 0})
			ExpectEqual(t, util.Links("https://www.tic.com/", dummyHTML2), model.Link{13, 9, 0})
			ExpectEqual(t, util.Links("https://www.tic.com/", dummyHTML3), model.Link{13, 9, 0})
		})
	}

	t.Log("Determine if login form is exist in document")
	{
		t.Run("HasLoginForm", func(t *testing.T) {
			ExpectEqual(t, util.HasLoginForm(dummyHTML1), false)
			ExpectEqual(t, util.HasLoginForm(dummyHTML2), false)
			ExpectEqual(t, util.HasLoginForm(dummyHTML3), false)
			ExpectEqual(t, util.HasLoginForm(dummyHTML4), true)
		})
	}
}
