package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"net/url"
	"os"
	"stepy/config"
)

type Database gorm.DB

type Model struct {
}

var db gorm.DB

func init() {
	db := open()

	if (!db.HasTable(&User{})) {
		db.CreateTable(&User{})
	}
	if (!db.HasTable(&NoteBook{})) {
		db.CreateTable(&NoteBook{})
	}
	if (!db.HasTable(&Item{})) {
		db.CreateTable(&Item{})
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&NoteBook{})
	db.AutoMigrate(&Item{})
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
	var _config string
	var format = "%s@%s([%s]:%d)/%s?parseTime=true"

	if os.Getenv("CLEARDB_DATABASE_URL") != "" {
		url, _ := url.Parse(os.Getenv("CLEARDB_DATABASE_URL"))
		_config = fmt.Sprintf(format,
			url.User.String(),
			config.App.Db.Protocol,
			url.Host,
			config.App.Db.Port,
			url.Path)
	} else {
		_config = fmt.Sprintf(format,
			config.App.Db.User,
			config.App.Db.Protocol,
			config.App.Db.Host,
			config.App.Db.Port,
			config.App.Db.Name)
	}

	fmt.Printf(_config)
	return _config
}
