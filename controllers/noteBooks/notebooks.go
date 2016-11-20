package NoteBooks

import (
	"errors"
	"fmt"
	"net/http"
	sHttp "stepy/http"
	"strings"
)

type Notebooks struct {
	Id string
	sHttp.Protocol
}

func Controller(protocol sHttp.Protocol) {
	url := strings.Replace(protocol.Req.URL.Path, "/", "", 1)
	paths := strings.Split(url, "/")
	notebook := Notebooks{Protocol: protocol, Id: ""}

	switch len(paths) {
	case 1:
		switch protocol.Req.Method {
		case http.MethodPost:
			notebook.create()
		case http.MethodGet:
			notebook.list()
		default:
			protocol.Wr.WriteHeader(404)
		}
	case 2:
		notebook.Id = paths[1]
		switch protocol.Req.Method {
		case http.MethodPut:
			notebook.update()
		default:
			protocol.Wr.WriteHeader(404)
		}
	case 3, 4:
		notebook.Id = paths[1]
		if paths[2] != "items" {
			protocol.Wr.WriteHeader(404)
			return
		}
		notebook.ItemsController(paths)
	default:
		protocol.Wr.WriteHeader(404)
	}
}

func (n Notebooks) list() {
	deviceId, ok := n.Session.Values["device_id"].(uint)
	if !ok {
		n.Error(http.StatusBadRequest, errors.New("invalid arg."))
		return
	}

	notebooks, err := n.DB.FindAllNotebooksByDeviceId(deviceId)
	if err != nil {
		n.Error(http.StatusBadRequest, nil)
	}

	n.JsonWithInterface(notebooks)
}

func (n Notebooks) create() {
	deviceId, ok := n.Session.Values["device_id"].(uint)
	if !ok {
		n.Error(http.StatusBadRequest, errors.New("invalid arg."))
		return
	}

	title := n.Req.PostFormValue("title")
	if title == "" {
		n.Error(http.StatusBadRequest, errors.New("invalid arg."))
		return
	}

	list, err := n.DB.CreateNotebook(title, deviceId)
	if err != nil {
		fmt.Println(err.Error())
		n.Error(http.StatusBadRequest, nil)
		return
	}

	n.JsonWithInterface(list)
}

func (n Notebooks) delete() {

}

func (n Notebooks) update() {

}

func (n Notebooks) get() {
	_, ok := n.Session.Values["device_id"].(uint)
	if !ok {
		n.Error(http.StatusBadRequest, errors.New("invalid arg."))
		return
	}

	notebook, err := n.DB.FindListByIdWithItems(n.Id)
	if err != nil {
		n.Error(http.StatusBadRequest, nil)
		return
	}

	n.JsonWithInterface(notebook)
}
