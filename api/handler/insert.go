package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mirito333/spotify-time-search-api/api/connection"
	"github.com/mirito333/spotify-time-search-api/api/spotify"
)

func SaveTrack(c *gin.Context) {
	fmt.Println("hoge")
	trackId := c.Query("id")
	fmt.Println(trackId)
	db := connection.GetDB()
	session := spotify.GetSession()
	track := session.GetTrackInfo(trackId)
	db.Create(&track)
	c.JSON(200, track)
}

func SavePlayList(c *gin.Context) {
	trackId := c.Query("id")
	db := connection.GetDB()
	session := spotify.GetSession()
	playlist,track := session.GetPlaylist(trackId)
	db.Create(&playlist)
	for _, t := range track{
		db.Create(&t)
	}
	c.JSON(200, track)
}
