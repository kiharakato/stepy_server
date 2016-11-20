package Devices

import (
	"errors"
	"fmt"
	"net/http"
	sHttp "stepy/http"
	"strings"
)

type Devices struct {
	sHttp.Protocol
	Id string
}

func Controller(protocol sHttp.Protocol) {
	url := strings.Replace(protocol.Req.URL.Path, "/", "", 1)
	paths := strings.Split(url, "/")

	devices := Devices{protocol, ""}
	method := protocol.Req.Method

	switch len(paths) {
	case 1:
		switch method {
		case http.MethodPost:
			devices.create()
		default:
			protocol.Wr.WriteHeader(404)
		}
	case 2:
		devices.Id = string(paths[1])
		switch method {
		case http.MethodPut:
			devices.update()
		default:
			protocol.Wr.WriteHeader(404)
		}
	default:
		protocol.Wr.WriteHeader(404)
	}
}

func (d Devices) list() {
	d.Json([]byte(nil))
}

func (d Devices) get() {
	device, err := d.DB.FindDeviceById(d.Id)
	if err != nil {
		fmt.Println(err.Error())
		d.Error(http.StatusNotFound, nil)
		return
	}

	d.JsonWithInterface(device)
}

func (d Devices) create() {
	uniqueId := d.Req.PostFormValue("unique_id")

	if uniqueId == "" {
		d.Error(http.StatusBadRequest, errors.New("invald arg."))
		return
	}

	device, err := d.DB.CreateDevice(uniqueId)
	if err != nil {
		fmt.Println(err.Error())
		d.Error(http.StatusBadRequest, nil)
		return
	}

	d.Session.Values["device_id"] = device.ID
	d.SessionSave()

	d.JsonWithInterface(device)
}

func (d Devices) delete() {

}

func (d Devices) update() {

}
