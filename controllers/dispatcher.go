package controllers

import (
	"net/http"
	"stepy/controllers/devices"
	"strings"
	"stepy/controllers/noteBooks"
	"fmt"
	"time"
)

func init() {
	http.HandleFunc("/", dispatcher)
}

func dispatcher(wr http.ResponseWriter, req *http.Request) {
	url := req.URL.Path

	if strings.Index(url, "/ping") > -1 {
		Ping(wr, req)
	}

	if strings.Index(url, "/devices") > -1 {
		Devices.Controller(wr, req)
	}

	if strings.Index(url, "/notebooks") > -1 {
		NoteBooks.Controller(wr, req)
	}

}

func Ping(wr http.ResponseWriter, req *http.Request) {
	wr.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(wr, `{"status": "ok", "date": %s }`, time.Now())
}