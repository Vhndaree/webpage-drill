package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"webpageinfo/model"
	"webpageinfo/util"
)

// func ReadWebPage(url string) model.Data {
func ReadWebPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")

		url := util.AddURLProtocol(r.URL.Query().Get("url"))
		if url == "" {
			handleEmptyURL(w)
		}

		response, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()

		dataInBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
			return
		}

		content := string(dataInBytes)
		d := generateResponse(url, content)

		res, err := json.Marshal(&d)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
			return
		}

		w.Write(res)
		return
	}

	w.WriteHeader(200)
	var resp = make(map[string]string)
	resp["message"] = "There is no task at for the request"
	res, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
	}

	w.Write(res)
}

func handleEmptyURL(w http.ResponseWriter) {
	var resp = make(map[string]string)
	resp["message"] = "url is not valid"

	res, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("[ERROR] Marshalling data")
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(400)
	w.Write(res)
}

func generateResponse(url, content string) model.Data {
	return model.Data{
		HTMLVersion:  util.Version(content),
		Title:        util.Title(content),
		HeadingCount: util.Headings(content),
		Link:         util.Links(util.SanitizeURL(url), content),
		HasLoginForm: util.HasLoginForm(content),
	}
}
