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
	State      string     `gorm:"type: enum('done', 'delete', 'run'); default: 'run'; not null"`
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

func (db SDB) UpdateItems(itemId, title, state string) (Item, error) {
	var item Item

	option := map[string]interface{}{}
	if title != "" {
		option["title"] = title
	}
	if state != "" {
		option["state"] = state
	}

	err := db.Model(&item).Where("id=?", itemId).Updates(option).Error
	if err != nil {
		return item, err
	}

	return item, nil
}
