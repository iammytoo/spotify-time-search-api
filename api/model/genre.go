package model

import "github.com/jinzhu/gorm"

type Genre struct {
	gorm.Model

	Genre string `gorm:"unique"`
}
