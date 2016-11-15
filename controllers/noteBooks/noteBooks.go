package NoteBooks

import (
	"net/http"
	"stepy/db"
	sHttp "stepy/http"
	"strconv"
	"regexp"
)

type Notebooks struct {
	sHttp.Protocol
	Id string
}

func Controller(wr http.ResponseWriter, req *http.Request) {
	regex := regexp.MustCompile(`/notebooks`)
	match := regex.FindSubmatch([]byte(req.URL.Path))

	notebook := Notebooks{sHttp.Protocol{Wr: wr, Req: req}, ""}

	matchLen := len(match)
	if matchLen == 0 {
		notebook.list()
		return
	}

	notebook.Id = string(match[1])
	switch matchLen {
	case 2:
		switch req.Method {
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
		wr.WriteHeader(404)
	}
}

func (n Notebooks) list() {

}

func (n Notebooks) create() {
	userUuid := n.Req.PostFormValue("uuid")
	title := n.Req.PostFormValue("title")
	list := db.CreateNoteBook(title, userUuid)
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
