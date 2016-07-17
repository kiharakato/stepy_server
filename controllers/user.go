package controllers

import (
	"github.com/satori/go.uuid"
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

func (u users) Post(req *http.Request) (stepyHttp.APIStatus, interface{}) {
	uuid := uuid.NewV4()
	email := req.PostFormValue("email")
	name_1 := req.PostFormValue("name")
	user := User{Id: name_1, Email: email, address: uuid.String()}
	//stepyDb.CreateUser(uuid, email)
	return stepyHttp.Success(http.StatusOK), user
}
