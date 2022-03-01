package models

import "github.com/jinzhu/gorm"

type Users struct {
	gorm.Model
	UID        int64  `gorm:"unique" `
	Gender     string ``
	UserName   string `gorm:"unique"`
	Password   string ``
	RePassword string ``
	Email      string `gorm:"unique"`
}
