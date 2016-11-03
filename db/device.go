package db

import (
	"time"
)

func init() {
}

type Device struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Os        string     `gorm:"size:255;not null" json:"os"`
	DeviceID  string     `gorm:"text not null" json:"device_id"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func CreateDevice(deviceId string) interface{} {
	device := &Device{DeviceID: deviceId}

	db := open()
	db.Create(device)
	return device
}

func ReadDeviceByDeviceId(deviceId string) interface{} {
	var device Device
	db := open()
	db.Where("device_id= ?", deviceId).First(&device)
	return device
}
