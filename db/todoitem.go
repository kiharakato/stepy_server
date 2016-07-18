package db

import (
	"time"
)

func init() {
}

type TodoItem struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Title     string    `gorm:"size:255;not null" json:"name"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}

func CreateTodoItem(title string) interface{} {
	item := &TodoItem{Title: title}

	db := open()
	db.Create(item)
	return item
}
