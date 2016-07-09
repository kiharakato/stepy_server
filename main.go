package main

import (
	"net/http"
	_ "stepy/controllers"
)

func main() {
	http.ListenAndServe(":5000", nil)
}
