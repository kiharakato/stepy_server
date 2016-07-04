package db

import (
	"time"
)

type User struct {
	Id        int64     `json: "id"`
	Email     string    `json: "email"`
	Name      string    `json: "name"`
	UUID      string    `json: "uuid"`
	CreatedAt time.Time `json: "created_at"`
}

func CreateUser(uuid, email, name string) {
	user := &User{
		Email:     email,
		Name:     name,
		UUID:      uuid,
		CreatedAt: time.Now(),
	}

	_db := open()
	_db.Create(user)
}
