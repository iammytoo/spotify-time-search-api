package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mirito333/spotify-time-search-api/api/connection"
	"github.com/mirito333/spotify-time-search-api/api/model"
	"github.com/mirito333/spotify-time-search-api/api/spotify"
	"github.com/mirito333/spotify-time-search-api/api/utils"
)

func SaveTrack(c *gin.Context) {
	fmt.Println("hoge")
	trackId := c.Query("id")
	fmt.Println(trackId)
	db := connection.GetDB()
	session := spotify.GetSession()
	trackData := session.GetTrack(trackId)
	saveTrackData(trackData, db)
	c.JSON(200, trackData)
}

func SavePlayList(c *gin.Context) {
	trackId := c.Query("id")
	db := connection.GetDB()
	session := spotify.GetSession()
	playlist, trackDatas := session.GetPlaylist(trackId)
	db.Save(&playlist)
	for _, t := range trackDatas {
		saveTrackData(t,db)
	}
	c.JSON(200, trackDatas)
}

func saveTrackData(t model.TrackData, db *gorm.DB) {
	db.Save(&t.Track)
	for _, a := range t.Artists {
		db.Save(&a)
	}
	for _, g := range t.Genres {
		db.Save(&g)
	}
	utils.CreateArtistTrack(t.TrackKey,t.ArtistKeys)
	for _, agk := range t.AGKeys {
		utils.CreateArtistGenre(agk.ArtistKey,agk.GenreKeys)
	}
}
