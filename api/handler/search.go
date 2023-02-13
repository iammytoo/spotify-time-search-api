package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mirito333/spotify-time-search-api/api/connection"
	"github.com/mirito333/spotify-time-search-api/api/model"
)

func GetTrack(c *gin.Context){
	time := c.Query("time")
	db := connection.GetDB()
	result := db.Where("duration = ?", time).First(&model.Track{})
	c.JSON(200, result)
}