package controllers

import (
	"io"
	"net/http"
	"net/url"
	stepyHttp "stepy/http"
)

type users struct {
	stepyHttp.APIResourceBase
}

func init() {
	println("load controller / user")
	http.Handle("/user", stepyHttp.Chain(stepyHttp.APIResourceHandler(users{})))
}

type User struct {
	Id      string `json: id`
	Email   string `email: id`
	address string `address: id`
}

func (u users) Get(url string, queries url.Values, body io.Reader) (stepyHttp.APIStatus, interface{}) {
	user := User{Id: "test", Email: "test.com", address: "test ken"}
	return stepyHttp.Success(http.StatusOK), user
}
