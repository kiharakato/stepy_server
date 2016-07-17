package main

import (
	"net/http"
	"os"
	_ "stepy/controllers"
)

func main() {
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, nil)
}
