package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mirito333/spotify-time-search-api/api/connection"
	"github.com/mirito333/spotify-time-search-api/api/model"
)

func GetTrack(c *gin.Context) {
	time := c.Query("time")
	minTime := c.DefaultQuery("min_time", "0")
	maxTime := c.DefaultQuery("max_time", "100000000")
	category := c.DefaultQuery("category", "no")
	db := connection.GetDB()
	relation := db
	if time != "" {
		relation = relation.Where("`duration` = ?", time)
	}
	if minTime != "0" || maxTime != "100000000" {
		relation = relation.Where("`duration` between ? and ?", minTime, maxTime)
	}
	if category != "no" {
		genre := model.Genre{}
		db.Where("`genre` = ?" ,category).First(&genre)
		relation = relation.Table("tracks").Joins("join track_artists on `track_artists`.`track_id` = `tracks`.`id`")
		relation = relation.Joins("join artist_genres on `track_artists`.`artist_id` = `artist_genres`.`artist_id`").Where("`artist_genres`.`genre_id` = ?" , genre.ID)
	}
	result := relation.Order("times desc").First(&model.Track{})
	c.JSON(200, result)
}
