package controllers

import (
	"net/http"
	"stepy/controllers/devices"
	"stepy/controllers/noteBooks"
	resisSession "gopkg.in/boj/redistore.v1"
	"strings"
	"fmt"
	"time"
	sHttp "stepy/http"
	"stepy/config"
)

func init() {
	http.HandleFunc("/", dispatcher)
}

func dispatcher(wr http.ResponseWriter, req *http.Request) {
	url := req.URL.Path

	store, err := resisSession.NewRediStore(10, "tcp", ":6379", config.App.Session.Password, []byte(config.App.Session.SecretKey))
	if err != nil {
		panic(err)
	}
	defer store.Close()

	session, err := store.Get(req, config.App.Session.SessionKey)
	if err != nil {
		panic(err)
	}

	protocol := sHttp.Protocol{Wr: wr, Req: req, Session: session}

	if strings.Index(url, "/ping") > -1 {
		Ping(protocol)
	}

	if strings.Index(url, "/devices") > -1 {
		Devices.Controller(protocol)
	}

	if strings.Index(url, "/notebooks") > -1 {
		NoteBooks.Controller(protocol)
	}

}

func Ping(protocol sHttp.Protocol) {
	protocol.Wr.Header().Set("Content-Type", "application/json")
	protocol.Session.Values["test"] = "かみちゅ"
	protocol.SessionSave()
	fmt.Fprintf(protocol.Wr, `{"status": "ok", "date": %s }`, time.Now())
}