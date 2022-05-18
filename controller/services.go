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
			return
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
		d := model.Data{
			HTMLVersion:  util.Version(content),
			Title:        util.Title(content),
			HeadingCount: util.Headings(content),
			Link:         util.Links(util.SanitizeURL(url), content),
			HasLoginForm: util.HasLoginForm(content),
		}

		res, err := json.Marshal(&d)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
			return
		}

		w.Write(res)
	}
}
