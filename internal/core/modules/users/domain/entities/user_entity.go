package entities

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	FullName string `gorm:"varchar(50)"`
	Email    string `gorm:"varchar(50)"`
	PasswordHash string `gorm:"varchar(255)"`
}
