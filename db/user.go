package db

import (
	"time"
)

func init() {
}

type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Uuid      string     `gorm:"size:32;not null;unique_index" json:"uuid"`
	Email     string     `json:"email"`
	Name      string     `gorm:"size:255" json:"name"`
	CreatedAt time.Time  `gorm:"not null default CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null default CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func CreateUser(email, name string) interface{} {
	user := &User{
		Email: email,
		Name:  name,
	}

	_db := open()
	_db.Create(user)
	return user
}

func ReadUserByUuid(uuid string) interface{} {
	var user User
	db := open()
	db.Where("uuid= ?", uuid).First(&user)
	return user
}
