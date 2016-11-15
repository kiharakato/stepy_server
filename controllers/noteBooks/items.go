package NoteBooks

import (
)


type item struct {
	Notebooks
	ItemId string
}


//func (i item) Get(req *http.Request) (stepyHttp.APIStatus, interface{}) {
//	if id, _ := stepyHttp.RequestGetParam(req, "id"); len(id) != 0 {
//		user := db.ReadUserByUuid(id)
//		return stepyHttp.Success(http.StatusOK), user
//	}
//
//	return stepyHttp.Fail(http.StatusNotFound, "invalid user id"), nil
//}
//
//func (i item) Post(req *http.Request) (stepyHttp.APIStatus, interface{}) {
//	return stepyHttp.Success(http.StatusOK), nil
//}
