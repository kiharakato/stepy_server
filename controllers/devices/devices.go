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

func Controller(wr http.ResponseWriter, req *http.Request) {
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
	deviceId := d.Req.PostFormValue("device_id")
	device := db.CreateDevice(deviceId)
	d.Json(device)
}

func (d Devices) delete() {

}

func (d Devices) update() {

}
