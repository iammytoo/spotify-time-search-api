package model

import "github.com/jinzhu/gorm"

type Track struct {
	gorm.Model

	Name     string
	Key      string `gorm:"unique"`
	Duration int
	Times    int
}
