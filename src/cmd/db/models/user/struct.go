package user

import (
	"chat-websocket/src/cmd/db"
	"chat-websocket/src/cmd/db/models"
)

func init() {
	db.GetDB().AutoMigrate(&User{})
}

type User struct {
	models.Model
	Name     string `gorm:"type:varchar(30);unique_index;not null"`
	Password string `gorm:"type:varchar(30); not null"`
	Token    string `gorm:"type:varchar(300);unique_index;not null"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) SaveNew() {
	db.GetDB().Create(u)
}

func (u *User) Update() {
	db.GetDB().Model(u).Updates(u)
}

func (u *User) Delete() {
	db.GetDB().Delete(u)
}

func FindById(id int) *User {
	var user *User
	db.GetDB().Where("id = ?", id).Find(user)
	return user
}

func FindByNamePass(name string, pass string) *User {
	var user = &User{}
	db.GetDB().Where("name = ? and password = ?", name, pass).Find(user)
	if user.ID != 0 {
		return user
	}
	return nil
}
