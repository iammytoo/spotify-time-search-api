package model

import "github.com/jinzhu/gorm"

type Track struct {
	gorm.Model

	Name     string
	Key      string
	Duration int
	Times    int
}
