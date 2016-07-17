package controllers

import (
	"net/http"
	"stepy/db"
	stepyHttp "stepy/http"
)

type users struct {
	stepyHttp.APIResourceBase
}

func init() {
	http.Handle("/user", stepyHttp.Chain(stepyHttp.APIResourceHandler(users{})))
}

func (u users) Get(req *http.Request) (stepyHttp.APIStatus, interface{}) {
	if id, _ := stepyHttp.RequestGetParam(req, "id"); len(id) != 0 {
		user := db.ReadUserByUuid(req.URL.Path)
		return stepyHttp.Success(http.StatusOK), user
	}

	return stepyHttp.Fail(http.StatusNotFound, "invalid user id"), nil
}

func (u users) Post(req *http.Request) (stepyHttp.APIStatus, interface{}) {
	email := req.PostFormValue("email")
	name := req.PostFormValue("name")
	user := db.CreateUser(email, name)
	return stepyHttp.Success(http.StatusOK), user
}
