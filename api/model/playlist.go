package model

import "github.com/jinzhu/gorm"

type PlayList struct {
	gorm.Model

	Name      string
	Key       string
	IsFetched bool
}
