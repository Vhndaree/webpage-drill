package api

import (
	"fmt"
	"log"
	"net/http"
)

var green = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
var reset = string([]byte{27, 91, 48, 109})

func StartWeb() {
	port := ":8848"
	fmt.Printf("%sStarting API server on port%s.... %s \n", green, port, reset)
	log.Println("Press Ctrl+C to stop the server.")

	http.HandleFunc("/", ReadWebPage)

	http.ListenAndServe(port, nil)
}
