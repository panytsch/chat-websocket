package main

import (
	"chat-websocket/src/db"
	"chat-websocket/src/db/models/user"
	"fmt"
)

func main() {
	d := db.GetDB()
	defer d.Close()

	us := user.User{}
	us.Name = "Roma"
	us.Password = "Test"
	_, err := us.SaveNew()
	fmt.Println(err)

	fmt.Println(user.FindById(5))
}
