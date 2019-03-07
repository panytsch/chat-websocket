package db

import (
	"database/sql"
	"log"
	_ "src/github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	d, err := sql.Open("mysql", "root:Test-5949@/chat-go")
	if err != nil {
		log.Fatal(err)
	}
	db = d
}

func GetDB() *sql.DB {
	return db
}
