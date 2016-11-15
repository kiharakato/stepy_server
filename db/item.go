package db

import (
	"time"
)

func init() {
}

type Item struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Title     string     `gorm:"size:255;not null" json:"name"`
	ListID    uint       `json:"list_id"`
	State     string     `gorm:"size:32;not null"`
	CreatedAt time.Time  `gorm:"not null default CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null default CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func CreateItem(title string) interface{} {
	item := &Item{Title: title}

	db := open()
	db.Create(item)
	return item
}
