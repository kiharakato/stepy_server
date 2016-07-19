package controllers

import (
	"net/http"
	"stepy/db"
	stepyHttp "stepy/http"
)

type lists struct {
	stepyHttp.APIResourceBase
}

type items struct {
	stepyHttp.APIResourceBase
}

func init() {
	http.Handle("/list/", stepyHttp.Chain(stepyHttp.APIResourceHandler(lists{})))
	http.Handle("/list/item/", stepyHttp.Chain(stepyHttp.APIResourceHandler(items{})))
}

func (l lists) Get(req *http.Request) (stepyHttp.APIStatus, interface{}) {
	if id, _ := stepyHttp.RequestGetParam(req, "id"); len(id) != 0 {
		user := db.ReadUserByUuid(id)
		return stepyHttp.Success(http.StatusOK), user
	}

	return stepyHttp.Fail(http.StatusNotFound, "invalid user id"), nil
}

func (l lists) Post(req *http.Request) (stepyHttp.APIStatus, interface{}) {
	title := req.PostFormValue("title")
	list := db.CreateTodoList(title)
	return stepyHttp.Success(http.StatusOK), list
}

func (i items) Get(req *http.Request) (stepyHttp.APIStatus, interface{}) {
	if id, _ := stepyHttp.RequestGetParam(req, "id"); len(id) != 0 {
		user := db.ReadUserByUuid(id)
		return stepyHttp.Success(http.StatusOK), user
	}

	return stepyHttp.Fail(http.StatusNotFound, "invalid user id"), nil
}

func (i items) Post(req *http.Request) (stepyHttp.APIStatus, interface{}) {
	title := req.PostFormValue("title")
	list := db.CreateTodoList(title)
	return stepyHttp.Success(http.StatusOK), list
}
