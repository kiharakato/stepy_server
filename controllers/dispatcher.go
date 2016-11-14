package controllers

import (
	"net/http"
	"stepy/controllers/devices"
	"strings"
	"stepy/controllers/noteBooks"
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