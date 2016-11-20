package NoteBooks

import (
	"errors"
	"fmt"
	"net/http"
	sHttp "stepy/http"
	"strconv"
)

type Items struct {
	Id     string
	ItemId string
	sHttp.Protocol
}

func ItemsController(notebook Notebooks) {
	items := notebook.Items
	switch notebook.Req.Method {
	case http.MethodPost:
		items.create()
	case http.MethodPut:
		items.update()
	default:
		notebook.Wr.WriteHeader(404)
	}
}

func (i Items) create() {
	_, ok := i.Session.Values["device_id"].(string)
	if !ok {
		i.Error(http.StatusBadRequest, errors.New("invalid arg."))
		return
	}

	title := i.Req.PostFormValue("title")
	if title == "" {
		i.Error(http.StatusBadRequest, errors.New("invalid arg."))
		return
	}

	notebookId, _ := strconv.ParseUint(i.Id, 10, 10)
	item, err := i.DB.AddItemToNoteBook(uint(notebookId), title)
	if err != nil {
		fmt.Println(err.Error())
		i.Error(http.StatusBadRequest, nil)
		return
	}

	i.JsonWithInterface(item)
}

func (i Items) update() {
	title := i.Req.PostFormValue("title")
	state := i.Req.PostFormValue("state")

	item, err := i.DB.UpdateItems(i.ItemId, title, state)
	if err != nil {
		i.Error(http.StatusInternalServerError, nil)
		return
	}

	i.JsonWithInterface(item)
}
