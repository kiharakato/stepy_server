package app

import "github.com/go-martini/martini"

func main() {
	m := martini.Classic()

	m.Get("/", func() string {
		return "Hello world!"
	})

	m.Run()
}