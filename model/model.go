package model

/**
- HTML Version
- Page Title
- Headings count by level
- Amount of internal, external and inaccessible links
- If a page contains a login form
*/
type Data struct {
	HTMLVersion  string   `json:"html_version"`
	Title        string   `json:"title"`
	HeadingCount Headings `json:"heading_count"`
	Link         Link     `json:"links"`
	HasLoginForm bool     `json:"has_login_form"`
}

type Headings struct {
	H1 int `json:"h1"`
	H2 int `json:"h2"`
	H3 int `json:"h3"`
	H4 int `json:"h4"`
	H5 int `json:"h5"`
	H6 int `json:"h6"`
}

type Link struct {
	External          int `json:"external_links"`
	Internal          int `json:"internal_links"`
	InaccessibleLinks int `json:"inaccessible_links"`
}
