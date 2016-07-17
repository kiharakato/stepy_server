package main

import (
	"net/http"
	"os"
	_ "stepy/controllers"
	"strconv"
)

func main() {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	http.ListenAndServe(":"+port, nil)
}
