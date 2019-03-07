package user

import (
	"chat-websocket/src/db"
	"database/sql"
)

type User struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Password string `db:"password"`
}

func (u *User) SaveNew() (sql.Result, error) {
	return db.GetDB().Exec("INSERT INTO `users` (`name`, `password`) VALUES (?, ?);", u.Name, u.Password)
}

func (u *User) Update() (sql.Result, error) {
	return db.GetDB().Exec("UPDATE `users` SET `name` = ?, `password` = ? WHERE `id` = ?", u.Name, u.Password, u.ID)
}

func (u *User) Delete() (sql.Result, error) {
	return db.GetDB().Exec("DELETE FROM `users` WHERE `id` = ?", u.ID)
}

func FindById(id int) (User, error) {
	var user User
	err := db.GetDB().QueryRow("SELECT * FROM `users` WHERE `id` = ?", id).Scan(&user.ID, &user.Name, &user.Password)
	return user, err
}
