package Devices

import (
	"net/http"
	"regexp"
	"stepy/db"
	sHttp "stepy/http"
)

type Devices struct {
	sHttp.Protocol
	Id string
}

func DevicesController(wr http.ResponseWriter, req *http.Request) {
	regex := regexp.MustCompile(`/devices/(\d+/?)$`)
	match := regex.FindSubmatch([]byte(req.URL.Path))

	devices := Devices{sHttp.Protocol{Wr: wr, Req: req}}

	matchLen := len(match)
	if matchLen == 0 {
		devices.list()
		return
	}

	devices.Id = string(match[1])
	switch matchLen {
	case 2:
		switch req.Method {
		case http.MethodPost:
			devices.create()
		case http.MethodDelete:
			devices.delete()
		case http.MethodPut:
			devices.update()
		default:
			devices.get()
		}
	default:
		wr.WriteHeader(404)
	}
}

func (d Devices) list() {
	d.Json([]byte(nil))
}

func (d Devices) get() {
	device := db.ReadDeviceById(d.Id)
	d.Json(device)
}

func (d Devices) create() {
	db.CreateDevice()
}

func (d Devices) delete() {

}

func (d Devices) update() {

}

//func (d Device) Get(req *http.Request) (stepyHttp.APIStatus, interface{}) {
//	if id, _ := stepyHttp.RequestGetParam(req, "id"); len(id) != 0 {
//		device := db.ReadDeviceByDeviceId(id)
//		return stepyHttp.Success(http.StatusOK), device
//	}
//
//	return stepyHttp.Fail(http.StatusNotFound, "invalid user id"), nil
//}
//
//func (d Device) Post(req *http.Request) (stepyHttp.APIStatus, interface{}) {
//	deviceId := req.PostFormValue("device_id")
//	device := db.CreateDevice(deviceId)
//	return stepyHttp.Success(http.StatusOK), device
//}
