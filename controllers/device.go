package controllers

import (
	"net/http"
	"stepy/db"
	stepyHttp "stepy/http"
)

type Device struct {
	stepyHttp.APIResourceBase
}

func init() {
	http.Handle("/device", stepyHttp.Chain(stepyHttp.APIResourceHandler(users{})))
}

func (d Device) Get(req *http.Request) (stepyHttp.APIStatus, interface{}) {
	if id, _ := stepyHttp.RequestGetParam(req, "id"); len(id) != 0 {
		device := db.ReadDeviceByDeviceId(id)
		return stepyHttp.Success(http.StatusOK), device
	}

	return stepyHttp.Fail(http.StatusNotFound, "invalid user id"), nil
}

func (d Device) Post(req *http.Request) (stepyHttp.APIStatus, interface{}) {
	deviceId := req.PostFormValue("device_id")
	device := db.CreateDevice(deviceId)
	return stepyHttp.Success(http.StatusOK), device
}
