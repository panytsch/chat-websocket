package db

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Login string `gorm:"type:varchar(30);unique_index;not null"`
}
