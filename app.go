package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"stepy/db"
	"time"
)

type StatusJson struct {
	Status string `json:status`
	Error  string `json:error`
}

func main() {

	http.HandleFunc("/ping", func(wr http.ResponseWriter, req *http.Request) {
		wr.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(wr, `{"status": "ok, "date": %s }`, time.Now())
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"message": "hello stepy"}`)
		return
	})

	http.HandleFunc("/register", func(rw http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			register(rw, req)
			return
		} else {
			http.NotFound(rw, req)
			return
		}
	})

	http.ListenAndServe(":5000", nil)
}

func okJson() string {
	status := StatusJson{"ok", ""}

	bytes, err := json.Marshal(status)
	if err != nil {
		return ""
	}

	statusString := string(bytes)
	return statusString
}

func register(res http.ResponseWriter, req *http.Request) string {
	uuid := req.PostFormValue("uuid")

	db.CreateUser(uuid, "", "")
	return okJson()
}
