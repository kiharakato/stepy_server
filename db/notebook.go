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
	DeviceID  uint       `gorm:"not null" json:"device_id"`
	CreatedAt time.Time  `gorm:"not null; default: CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null; default: CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Items     []Item
}

func (db SDB) CreateNotebook(title string, deviceId uint) (Notebook, error) {
	tx := db.Begin()

	noteBook := Notebook{
		Title:    title,
		DeviceID: deviceId,
	}

	if err := db.Create(&noteBook).Error; err != nil {
		tx.Rollback()
		return noteBook, err
	}

	tx.Commit()
	return noteBook, nil
}

func (db SDB) FindListByIdWithItems(listId string) (interface{}, error) {
	var note Notebook
	note.Items = []Item{}
	if err := db.Where("id=?", listId).First(&note).Model(&note).Related(&note.Items).Error; err != nil {
		return note, err
	}
	return note, nil
}

func (db SDB) FindAllNotebooksByDeviceId(deviceId uint) ([]Notebook, error) {
	var notebooks []Notebook
	if err := db.Where("device_id = ?", deviceId).Find(&notebooks).Error; err != nil {
		return notebooks, err
	}

	for i := range notebooks {
		var items []Item
		if err := db.Where("notebook_id = ?", notebooks[i].ID).Find(&items).Error; err != nil {
			return notebooks, err
		}
		notebooks[i].Items = items
	}

	return notebooks, nil
}
