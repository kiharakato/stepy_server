package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"net/url"
	"os"
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
	var config = "b0b9f816ac0c4a:cd9775e9@tcp([localhost]:3306)/stepy?parseTime=true"

	if os.Getenv("CLEARDB_DATABASE_URL") != "" {
		url, _ := url.Parse(os.Getenv("CLEARDB_DATABASE_URL"))
		config = fmt.Sprintf("%s@tcp(%s:3306)%s", url.User.String(), url.Host, url.Path)
	}

	return config
}
