package db

import (
	"github.com/satori/go.uuid"
	"strings"
	"time"
)

func init() {
}

type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Uuid      string     `gorm:"size:32;not null;unique_index" json:"uuid"`
	Email     string     `json:"email"`
	Name      string     `gorm:"size:255" json:"name"`
	TodoLists []TodoList `json:"group"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt time.Time  `json:"deleted_at"`
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

func createUuid() string {
	return strings.Replace(uuid.NewV1().String(), "-", "", -1)
}

func (u *User) BeforeCreate() (err error) {
	u.Uuid = createUuid()
	return
}
