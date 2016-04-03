package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Database gorm.DB

func open() gorm.DB {
	// 接続するための情報文字列を作る
	d := Database{}
	connectionString := d.getConnectionString()

	// DB接続
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	db.DB()
	return db
}

func (d Database) getConnectionString() string {
	return "root:root@tcp([localhost]:3306)/stepy?parseTime=true"
}
