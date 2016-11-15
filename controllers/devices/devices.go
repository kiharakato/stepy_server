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

func Controller(protocol sHttp.Protocol) {
	regex := regexp.MustCompile(`/devices/(\d+/?)$`)
	match := regex.FindSubmatch([]byte(protocol.Req.URL.Path))

	devices := Devices{protocol, ""}

	matchLen := len(match)
	if matchLen == 0 {
		devices.list()
		return
	}

	devices.Id = string(match[1])
	switch matchLen {
	case 2:
		switch protocol.Req.Method {
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
		protocol.Wr.WriteHeader(404)
	}
}

func (d Devices) list() {
	d.Json([]byte(nil))
}

func (d Devices) get() {
	device := db.ReadDeviceById(d.Id)
	d.JsonWithInterface(device)
}

func (d Devices) create() {
	deviceUniqId := d.Req.PostFormValue("device_id")
	device := db.CreateDevice(deviceUniqId)

	d.Session.Values["device_id"] = device.ID
	d.SessionSave()

	d.JsonWithInterface(device)
}

func (d Devices) delete() {

}

func (d Devices) update() {

}
