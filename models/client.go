package models

import (
	"github.com/jinzhu/gorm"
)

// Client : A struct representing a client
type Client struct {
	gorm.Model

	Name    string `gorm:"column:name" json:"name"`
	Address string `gorm:"column:address" json:"address"`
	Note    string `gorm:"column:note" json:"note"`
}
