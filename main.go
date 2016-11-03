package main

import (
	"net/http"
	"stepy/config"
	_ "stepy/controllers"
)

func main() {
	http.ListenAndServe(":"+config.App.Port, nil)
}
