package model

import (
	"github.com/jinzhu/gorm"
)

type TrackArtist struct {
	gorm.Model

	ArtistID uint
	TrackID  uint
}

type ArtistGenre struct {
	gorm.Model

	ArtistID uint
	GenreID  uint
}
