package main

import "github.com/go-martini/martini"

func main() {
	m := martini.Classic()

	m.Get("/", func() string {
		return "Hello world!"
	})

	m.Post("/", func() string {
		return "Hello post!"
	})

	m.Run()
}
