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

func (db SDB) CreateDevice(uniqueId string) Device {
	device := Device{UniqueId: uniqueId}

	db.Create(&device)

	if db.NewRecord(&device) {
		println(db.Error.Error())
		return device
	}

	return device
}

func ReadDeviceById(deviceId string) Device {
	var device Device
	db := open()
	db.Where("device_id= ?", deviceId).First(&device)
	return device
}
