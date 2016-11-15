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
	DeviceID  string     `gorm:"not null" json:"device_id"`
	CreatedAt time.Time  `gorm:"not null default CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null default CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func CreateNoteBook(title string, deviceId string) NoteBook {
	db := open()
	noteBook := NoteBook{
		Title:  title,
		DeviceID: deviceId,
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
