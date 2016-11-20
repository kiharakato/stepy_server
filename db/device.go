package db

import (
	"time"
)

func init() {
}

type Device struct {
	ID uint `gorm:"primary_key" json:"id"`
	//Os        string     `gorm:"size:255;not null" json:"os"`
	UniqueId  string     `gorm:"not null; text; unique;" json:"unique_id"`
	CreatedAt time.Time  `gorm:"not null; default: CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null; default: CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (db SDB) CreateDevice(uniqueId string) (Device, error) {
	var device Device

	db.Where("unique_id=?", uniqueId).First(&device)
	if device != (Device{}) {
		return device, nil
	}

	tx := db.Begin()
	device = Device{UniqueId: uniqueId}
	if err := db.Create(&device).Error; err != nil {
		tx.Rollback()
		return device, err
	}
	tx.Commit()

	return device, nil
}

func (db SDB) FindDeviceById(deviceId string) (Device, error) {
	var device Device
	if err := db.Where("id=?", deviceId).First(&device).Error; err != nil {
		return device, err
	}
	return device, nil
}
