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

func init() {
	db := open()
	defer db.Close()

	if (!db.HasTable(&User{})) {
		db.CreateTable(&User{})
	}
	if (!db.HasTable(&Notebook{})) {
		db.CreateTable(&Notebook{})
	}
	if (!db.HasTable(&Item{})) {
		db.CreateTable(&Item{})
	}
	if (!db.HasTable(&Device{})) {
		db.CreateTable(&Device{})
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Notebook{})
	db.AutoMigrate(&Item{})
	db.AutoMigrate(&Device{})
}

type SDB struct {
	*gorm.DB
}

func NewDB() SDB {
	return open()
}

func open() SDB {
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

	db := SDB{_db}
	return db
}

func (d Database) getConnectionString() string {
	var _config string
	var format = "%s@%s([%s]:%d)/%s?parseTime=true"

	if os.Getenv("CLEARDB_DATABASE_URL") != "" {
		clearDBUrl, _ := url.Parse(os.Getenv("CLEARDB_DATABASE_URL"))
		_config = fmt.Sprintf(format,
			clearDBUrl.User.String(),
			config.App.Db.Protocol,
			clearDBUrl.Host,
			config.App.Db.Port,
			clearDBUrl.Path)
	} else {
		_config = fmt.Sprintf(format,
			config.App.Db.User,
			config.App.Db.Protocol,
			config.App.Db.Host,
			config.App.Db.Port,
			config.App.Db.Name)
	}

	return _config
}
