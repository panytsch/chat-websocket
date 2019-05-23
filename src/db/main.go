package db

import (
	"chat-websocket/src/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

func init() {
	log.Println("init db")
	var err error
	connectionString := ""
	login := config.AppConfig.DB.Login
	password := config.AppConfig.DB.Password
	dbName := config.AppConfig.DB.DBName
	connectionString += login + ":" + password + "@/" + dbName
	connectionString += "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open("mysql", connectionString)
	if err != nil {
		log.Fatalln(err.Error())
	}
	automigrate()
}

func automigrate() {
	GetDB().AutoMigrate(&User{})
	GetDB().AutoMigrate(&Message{})
}

func Close() {
	_ = db.Close()
}

func GetDB() *gorm.DB {
	return db
}
