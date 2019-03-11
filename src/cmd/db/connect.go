package db

import (
	"log"
	_ "src/github.com/go-sql-driver/mysql"
	"src/github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	d, err := gorm.Open("mysql", "root:Test-5949@/chat-go?charset=utf8mb4&parseTime=True")
	if err != nil {
		log.Fatal(err.Error())
	}

	db = d
}

func Close() {
	db.Close()
}

func GetDB() *gorm.DB {
	return db
}
