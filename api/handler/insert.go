package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mirito333/spotify-time-search-api/api/connection"
	"github.com/mirito333/spotify-time-search-api/api/model"
)

func SaveTrack(c *gin.Context) {
	db := connection.GetDB()
	defer db.Close()
	track_data := &model.Track{Name: "hoge",Key: "huga",Duration: 1223,Times: 999}
	db.Create(track_data)
	c.JSON(200, track_data)
}
