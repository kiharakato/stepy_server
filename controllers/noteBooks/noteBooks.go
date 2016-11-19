package NoteBooks

import (
	"net/http"
	"regexp"
	"stepy/db"
	sHttp "stepy/http"
	"strconv"
)

type Notebooks struct {
	sHttp.Protocol
	Id string
}

func Controller(protocol sHttp.Protocol) {
	regex := regexp.MustCompile(`/notebooks`)
	match := regex.FindSubmatch([]byte(protocol.Req.URL.Path))

	notebook := Notebooks{protocol, ""}

	matchLen := len(match)
	if matchLen == 0 {
		notebook.list()
		return
	}

	notebook.Id = string(match[1])
	switch matchLen {
	case 2:
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
	default:
		protocol.Wr.WriteHeader(404)
	}
}

func (n Notebooks) list() {

}

func (n Notebooks) create() {
	deviceId := n.Session.Values["device_id"]
	str, ok := deviceId.(string)
	if !ok {
		panic("cast fail")
	}

	title := n.Req.PostFormValue("title")
	list := db.CreateNoteBook(title, str)
	n.JsonWithInterface(list)
}

func (n Notebooks) delete() {

}

func (n Notebooks) update() {

}

func (n Notebooks) get() {
	id := n.Req.URL.Query().Get("id")

	if len(id) == 0 {
		n.JsonWithInterface(nil)
	}

	_id, err := strconv.ParseUint(id, 10, 32)
	if err == nil {
		n.JsonWithInterface(nil)
	}

	notebook := db.ReadListByIdWithItems(uint(_id))
	n.JsonWithInterface(notebook)
}
