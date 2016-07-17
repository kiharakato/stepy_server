package db

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"strings"
)

func init() {
}

type User struct {
	gorm.Model
	UUID  string `gorm:"suze:32;not null;unique_index", json:"uuid"`
	Email string `json:"email"`
	Name  string `gorm:"size:255", json:"name"`
}

func CreateUser(email, name string) interface{} {
	user := &User{
		Email: email,
		Name:  name,
		UUID:  CreateUuid(),
	}

	_db := open()
	_db.Create(user)
	return _db.Value
}

func CreateUuid() string {
	return strings.Replace(uuid.NewV1().String(), "-", "", -1)
}

func ReadUserByUuid(uuid string) interface{} {
	db := open()
	db.Select("SELECT * FROM user WHERE uuid = " + uuid)

	var user User
	db.First(user, "uuid = ?", uuid)

	return user
}
