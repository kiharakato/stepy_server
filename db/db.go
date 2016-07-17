package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

type Database gorm.DB

type Model struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"not null", json:"created_at"`
	UpdatedAt time.Time `gorm:"not null", json:"updated_at"`
	DeletedAt time.Time
}

var db gorm.DB

func init() {
	db := open()

	if (!db.HasTable(&User{})) {
		db.CreateTable(&User{})
	}

	db.AutoMigrate(&User{})
}

func open() gorm.DB {
	// 接続するための情報文字列を作る
	d := Database{}
	connectionString := d.getConnectionString()

	// DB接続
	_db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	_db.DB()
	_db.LogMode(true)

	db = _db
	return db
}

func (d Database) getConnectionString() string {
	return "root@tcp([localhost]:3306)/stepy?parseTime=true"
}
