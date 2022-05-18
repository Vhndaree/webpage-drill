package util

import (
	"net/http"
	"strings"
	"webpageinfo/model"
)

func Version(content string) string {
	title := strings.Split(content, "\n")[0]
	if !strings.Contains(title, "!DOCTYPE") {
		return "invalid html version"
	}

	if title == "<!DOCTYPE html>" {
		return "HTML 5"
	}

	if title == "<!DOCTYPE math SYSTEM \"http://www.w3.org/Math/DTD/mathml1/mathml.dtd\">" {
		return "MathML 1.01"
	}

	dtdStartIndex := strings.Index(title, "//DTD")
	if dtdStartIndex == -1 {
		return "ivalid html version"
	}

	dtdStartIndex += 6
	title = title[dtdStartIndex:]
	dtdEndIndex := strings.Index(title, "//")

	return title[:dtdEndIndex]
}

func Title(content string) string {
	titleStartIndex := strings.Index(content, "<title")
	if titleStartIndex == -1 {
		return "no title exist in the document"
	}

	titleStartIndex += 7
	titleEndIndex := strings.Index(content, "</title>")
	if titleEndIndex == -1 {
		return "valid title do not exist"
	}

	title := content[titleStartIndex:titleEndIndex]
	titleclosing := strings.Index(title, ">")
	return title[(titleclosing + 1):]
}

func Headings(content string) (h model.Headings) {
	h.H1 = func(c string) int {
		h1ClosingTags := strings.Split(c, "</h1>")
		return int(len(h1ClosingTags))
	}(content)

	h.H2 = func(c string) int {
		h2ClosingTags := strings.Split(c, "</h2>")
		return int(len(h2ClosingTags))
	}(content)

	h.H3 = func(c string) int {
		h3ClosingTags := strings.Split(c, "</h3>")
		return int(len(h3ClosingTags))
	}(content)

	h.H4 = func(c string) int {
		h4ClosingTags := strings.Split(c, "</h4>")
		return int(len(h4ClosingTags))
	}(content)

	h.H5 = func(c string) int {
		h5ClosingTags := strings.Split(c, "</h5>")
		return int(len(h5ClosingTags))
	}(content)

	h.H6 = func(c string) int {
		h6ClosingTags := strings.Split(c, "</h6>")
		return int(len(h6ClosingTags))
	}(content)

	return
}

func Links(url, content string) (l model.Link) {
	hrefStrings := strings.Split(content, "href=")
	for _, h := range hrefStrings {
		hrefStartIndex := strings.Index(h, "://")
		href := h
		if hrefStartIndex != -1 {
			href = h[(hrefStartIndex + 3):]
		}
		hrefQuote := h[:1]
		hrefEndIndex := strings.Index(href, hrefQuote)
		href = href[:hrefEndIndex]

		if !strings.Contains(href, ".") {
			l.Internal++
			continue
		}

		response, err := http.Get(AddURLProtocol(href))
		if err != nil {
			l.InaccessibleLinks++
			continue
		}
		defer response.Body.Close()

		if response.StatusCode < 200 || response.StatusCode > 299 {
			l.InaccessibleLinks++
			continue
		}

		domainEndIndex := strings.Index(url, "/")
		domain := url[:domainEndIndex]

		if strings.Contains(href, domain) {
			l.Internal++
		} else {
			l.External++
		}
	}

	return
}

func HasLoginForm(content string) bool {
	/**
	There is no any perfect criteria for a form to be a login form
	So I am building a ranking algorithm which gives rank to the page based on a content
	and keyword exists in the page.
	-------------------------------------------------------------------------------------
	Algorithm:
	1. Find if there is a keyword "form" exists in the document
	2. Find if there is a keyword "login" or "signin" exists in the document
	3. Find if there is a keyword "username" exists in the document
	4. Find if there is a keyword "user name" exists in the document
	5. Find if there is a keyword "password" exists in the document
	6. Find if there is a keyword "forgot password" exists in the document
	7. Find if there is a keyword "forget password" exists in the document
	8. Find if there is a keyword "email" exists in the document
	9. criteria - 60% and above score should be considered as true


	*/
	content = strings.ToLower(content)
	keywords := map[string]int{
		"form":            1,
		"login":           1,
		"signin":          1,
		"user name":       1,
		"username":        1,
		"password":        1,
		"forgot password": 5,
		"forget password": 5,
		"email":           1,
	}

	count := 0
	criteria := 0
	for key, val := range keywords {
		if strings.Contains(content, key) {
			criteria += val
		}

		count++
	}

	return (float32)(criteria)/float32(count) > 0.4
}
