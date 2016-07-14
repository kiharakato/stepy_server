package controllers

import (
	"net/http"
	"github.com/satori/go.uuid"
	stepyHttp "stepy/http"
	_ "stepy/db"
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

func (u users) Get(req *http.Request) (stepyHttp.APIStatus, interface{}) {
	user := User{Id: "test", Email: "test.com", address: "test ken"}
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