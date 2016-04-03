package main

import (
	"encoding/json"
	"github.com/go-martini/martini"
	"net/http"
	"stepy/db"
)

type StatusJson struct {
	Status string `json:status`
	Error  string `json:error`
}

func main() {
	m := martini.Classic()

	m.Get("/", func() string {
		return "Hello world!"
	})

	m.Post("/register", register)

	m.Run()
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
