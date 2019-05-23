package db

import (
	"github.com/jinzhu/gorm"
)

type Message struct {
	gorm.Model
	Text string `gorm:"type:text;not null"`
}
