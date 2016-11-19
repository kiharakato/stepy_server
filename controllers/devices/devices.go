package Devices

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
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
		switch protocol.Req.Method {
		case http.MethodPost:
			devices.create()
		default:
			devices.list()
		}
		return
	}

	devices.Id = string(match[1])
	switch matchLen {
	case 2:
		switch protocol.Req.Method {
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

	d.Session.Values["unique_id"] = device.ID
	d.SessionSave()

	d.JsonWithInterface(device)
}

func (d Devices) delete() {

}

func (d Devices) update() {

}
