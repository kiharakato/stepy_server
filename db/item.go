package db

import (
	"time"
)

func init() {
}

type Item struct {
	ID         uint       `gorm:"primary_key" json:"id"`
	Title      string     `gorm:"size:255;not null" json:"title"`
	NotebookId uint       `gorm:"not null" json:"notebook_id"`
	State      string     `gorm:"size ENUM('done', 'delete', 'run'); default: 'run'; not null"`
	CreatedAt  time.Time  `gorm:"not null; default: CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"not null; default: CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}

func (db SDB) AddItemToNoteBook(notebookId uint, itemTitle string) (Item, error) {
	if err := db.Select("id").Where("id=?", notebookId).First(&Notebook{}).Error; err != nil {
		return Item{}, err
	}

	item := Item{
		Title:      itemTitle,
		NotebookId: notebookId,
	}

	tx := db.Begin()
	if err := db.Create(&item).Error; err != nil {
		tx.Rollback()
		return item, err
	}

	tx.Commit()
	return item, nil
}
