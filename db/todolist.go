package db

import (
	"time"
)

func init() {
}

type TodoList struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Title     string     `gorm:"size:255;not null" json:"title"`
	Items     []TodoItem `json:"todo_items"`
	UserID    uint       `gorm:"not null" json:"user_id"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func CreateTodoList(title string, uuid string) interface{} {
	var user User
	db := open()
	db.Where("uuid = ?", uuid).First(&user)
	list := &TodoList{
		Title:  title,
		UserID: user.ID,
	}
	db.Create(list)
	return list
}

func AddItemToList(listId uint, item TodoItem) interface{} {
	var list TodoList
	db := open()
	db.Where("id = ?", listId).First(&list)
	items := append(list.Items, item)
	db.Model(&list).Update("TodoItems", items)

	return list
}

func ReadListByIdWithItems(listId uint) interface{} {
	var list TodoList
	var items []TodoItem
	db := open()
	db.Where("id = ?", listId).First(&list).Model(&list).Related(&items)
	list.Items = items
	return list
}
