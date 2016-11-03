package db

import (
	"time"
)

func init() {
}

type NoteBook struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Title     string     `gorm:"size:255;not null" json:"title"`
	UserID    uint       `gorm:"not null" json:"user_id"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func CreateNoteBook(title string, uuid string) interface{} {
	var user User
	db := open()
	db.Where("uuid = ?", uuid).First(&user)
	noteBook := &NoteBook{
		Title:  title,
		UserID: user.ID,
	}
	db.Create(noteBook)
	return noteBook
}

func AddItemToNoteBook(listId uint, item Item) interface{} {
	var noteBook NoteBook
	db := open()
	db.Where("id = ?", listId).First(&noteBook)
	db.Model(&noteBook).Update("TodoItems", item)

	return noteBook
}

func ReadListByIdWithItems(listId uint) interface{} {
	var list NoteBook
	var items []Item
	db := open()
	db.Where("id = ?", listId).First(&list).Model(&list).Related(&items)
	return list
}
