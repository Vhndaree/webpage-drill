package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"webpageinfo/model"
	"webpageinfo/util"
)

func ReadWebPage(url string) model.Data {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	dataInBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return model.Data{}
	}

	content := string(dataInBytes)
	return model.Data{
		HTMLVersion:  util.Version(content),
		Title:        util.Title(content),
		HeadingCount: util.Headings(content),
		Link:         util.Links(util.SanitizeURL(url), content),
		HasLoginForm: util.HasLoginForm(content),
	}
}
