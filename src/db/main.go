package db

import (
	"chat-websocket/src/config"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

func init() {
	log.Println("init db")
	var err error
	connectionString := ""
	conf := config.GetConfig().GetDBConfig()
	connectionString += conf.GetLogin() + ":" + conf.GetPassword() + "/" + conf.GetDBName()
	connectionString += "?charset=utf8mb4&parseTime=True"
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
