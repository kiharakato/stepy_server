package db

import (
	"time"
)

func init() {
}

type Notebook struct {
	ID    uint   `gorm:"primary_key" json:"id"`
	Title string `gorm:"size:255;not null" json:"title"`
	//UserID    uint       `json:"user_id"`
	DeviceID  string     `gorm:"not null" json:"device_id"`
	CreatedAt time.Time  `gorm:"not null; default: CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null; default: CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Items     []Item     `gorm:"-"`
}

func (db SDB) CreateNotebook(title string, deviceId string) (Notebook, error) {
	tx := db.Begin()

	noteBook := Notebook{
		Title:    title,
		DeviceID: deviceId,
	}

	if err := db.Create(noteBook).Error; err != nil {
		tx.Rollback()
		return noteBook, err
	}

	tx.Commit()
	return noteBook, nil
}

func (db SDB) FindListByIdWithItems(listId string) (interface{}, error) {
	var note Notebook
	note.Items = []Item{}
	if err := db.Where("id = ?", listId).First(&note).Model(&note).Related(&note.Items).Error; err != nil {
		return note, err
	}
	return note, nil
}
