package models

import (
	"github.com/jinzhu/gorm"
)

// User is a struct representing a client
type User struct {
	gorm.Model

	Username  string `gorm:"column:username;unique;not null" json:"name"`
	Email     string `gorm:"column:email;unique;not null" json:"email"`
	Password  string `gorm:"column:password;not null" json:"address"`
	AuthLevel uint   `gorm:"column:auth_level; default:2" json:"note"`
}
