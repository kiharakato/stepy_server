package controllers

import (
	"fmt"
	"net/http"
	"time"
)

func init() {
	println("load controller / ping")
	http.HandleFunc("/ping", request)
}

func request(wr http.ResponseWriter, req *http.Request) {
	wr.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(wr, `{"status": "ok, "date": %s }`, time.Now())
}
