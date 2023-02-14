package utils

import (
	"github.com/mirito333/spotify-time-search-api/api/connection"
	"github.com/mirito333/spotify-time-search-api/api/model"
)

func CreateArtistTrack(trackKey string, artists []string) {
	db := connection.GetDB()
	var track model.Track
	db.Where("`key` = ?", trackKey).First(&track)
	for _, a := range artists {
		var artist model.Artist
		db.Where("`key` = ?", a).First(&artist)
		row := model.TrackArtist{TrackID: track.ID, ArtistID: artist.ID}
		db.Save(&row)
	}
}

func CreateArtistGenre(artistKey string, genre []string) {
	db := connection.GetDB()
	var artist model.Artist
	db.Where("`key` = ?", artistKey).First(&artist)
	for _, g := range genre {
		var genre model.Genre
		db.Where("`genre` = ?", g).First(&genre)
		row := model.ArtistGenre{ArtistID: artist.ID, GenreID: genre.ID}
		db.Save(&row)
	}
}
