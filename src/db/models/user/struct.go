package user

import (
	"chat-websocket/src/db"
	"database/sql"
)

type User struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Password string `db:"password"`
	Token    string `db:"token"`
}

func (u *User) SaveNew() (sql.Result, error) {
	return db.GetDB().Exec("INSERT INTO `users` (`name`, `password`, `token`) VALUES (?, ?, ?)", u.Name, u.Password, u.Token)
}

func (u *User) Update() (sql.Result, error) {
	return db.GetDB().Exec("UPDATE `users` SET `name` = ?, `password` = ?, `token` = ? WHERE `id` = ?", u.Name, u.Password, u.Token, u.ID)
}

func (u *User) Delete() (sql.Result, error) {
	return db.GetDB().Exec("DELETE FROM `users` WHERE `id` = ?", u.ID)
}

func FindById(id int) (User, error) {
	var user User
	err := db.GetDB().QueryRow("SELECT * FROM `users` WHERE `id` = ?", id).Scan(&user.ID, &user.Name, &user.Password, &user.Token)
	return user, err
}

func FindByNamePass(name string, pass string) (*User, error) {
	var user User
	err := db.GetDB().QueryRow("SELECT * FROM `users` WHERE `name` = ? AND `password = ?`", name, pass).Scan(user)
	return &user, err
}
