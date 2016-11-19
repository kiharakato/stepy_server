package NoteBooks

import (
	"errors"
	"fmt"
	"net/http"
	sHttp "stepy/http"
	"strings"
)

type Notebooks struct {
	sHttp.Protocol
	Id string
	Items
}

func Controller(protocol sHttp.Protocol) {
	paths := strings.Split(protocol.Req.URL.Path, "/")
	notebook := Notebooks{Protocol: protocol, Id: "", Items: Items{}}

	switch len(paths) {
	case 0, 1:
		notebook.list()
	case 2:
		notebook.Id = paths[1]
		switch protocol.Req.Method {
		case http.MethodPost:
			notebook.create()
		case http.MethodDelete:
			notebook.delete()
		case http.MethodPut:
			notebook.update()
		default:
			notebook.get()
		}
	case 4:
		if paths[2] == "items" {
			protocol.Wr.WriteHeader(404)
			return
		}
		notebook.ItemId = paths[3]
		ItemsController(notebook)
	default:
		protocol.Wr.WriteHeader(404)
	}
}

func (n Notebooks) list() {

}

func (n Notebooks) create() {
	deviceId, ok := n.Session.Values["device_id"].(string)
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
	_, ok := n.Session.Values["device_id"].(string)
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
