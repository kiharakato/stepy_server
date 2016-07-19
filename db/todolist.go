package db

import (
	"time"
)

func init() {
}

type TodoList struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Title     string     `gorm:"size:255;not null" json:"title"`
	TodoItems []TodoItem `json:"todo_items"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt time.Time  `json:"deleted_at"`
}

func CreateTodoList(title string) interface{} {
	list := &TodoList{Title: title}
	db := open()
	db.Create(list)
	return list
}

func AddItemToList(listId uint, item TodoItem) interface{} {
	var list TodoList

	db := open()

	db.Where("id = ?", listId).First(&list)
	items := append(list.TodoItems, item)
	db.Model(&list).Update("TodoItems", items)

	return list
}