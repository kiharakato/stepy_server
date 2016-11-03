package controllers

import (
	"net/http"
	"stepy/db"
	stepyHttp "stepy/http"
	"strconv"
)

type noteBook struct {
	stepyHttp.APIResourceBase
}

type item struct {
	stepyHttp.APIResourceBase
}

func init() {
	http.Handle("/note_book/", stepyHttp.Chain(stepyHttp.APIResourceHandler(noteBook{})))
	http.Handle("/note_book/item/", stepyHttp.Chain(stepyHttp.APIResourceHandler(item{})))
}

func (l noteBook) Get(req *http.Request) (stepyHttp.APIStatus, interface{}) {
	if id, _ := stepyHttp.RequestGetParam(req, "id"); len(id) != 0 {
		if _id, err := strconv.ParseUint(id, 10, 32); err == nil {
			list := db.ReadListByIdWithItems(uint(_id))
			return stepyHttp.Success(http.StatusOK), list
		}
	}

	return stepyHttp.Fail(http.StatusNotFound, "invalid user id"), nil
}

func (l noteBook) Post(req *http.Request) (stepyHttp.APIStatus, interface{}) {
	userUuid := req.PostFormValue("uuid")
	title := req.PostFormValue("title")
	list := db.CreateNoteBook(title, userUuid)
	return stepyHttp.Success(http.StatusOK), list
}

func (i item) Get(req *http.Request) (stepyHttp.APIStatus, interface{}) {
	if id, _ := stepyHttp.RequestGetParam(req, "id"); len(id) != 0 {
		user := db.ReadUserByUuid(id)
		return stepyHttp.Success(http.StatusOK), user
	}

	return stepyHttp.Fail(http.StatusNotFound, "invalid user id"), nil
}

func (i item) Post(req *http.Request) (stepyHttp.APIStatus, interface{}) {
	return stepyHttp.Success(http.StatusOK), nil
}
