package NoteBooks

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type Items struct {
	ItemId string
	Notebooks
}

func (notebooks Notebooks) ItemsController(paths []string) {
	items := Items{
		Notebooks: notebooks,
		ItemId:    "",
	}

	switch len(paths) {
	case 3:
		items.create()
	case 4:
		items.update()
	default:
		items.Wr.WriteHeader(404)
	}
}

func (i Items) create() {
	_, ok := i.Session.Values["device_id"].(uint)
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
