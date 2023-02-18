package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mirito333/spotify-time-search-api/api/connection"
	"github.com/mirito333/spotify-time-search-api/api/model"
)

func GetTrack(c *gin.Context) {
	time := c.Query("time")
	minTime := c.DefaultQuery("min_time", "0")
	maxTime := c.DefaultQuery("max_time", "100000000")
	category := c.QueryArray("category")
	around := c.Query("around")
	db := connection.GetDB()
	relation := db
	if time != "" {
		relation = relation.Where("`duration` = ?", time)
	}
	if minTime != "0" || maxTime != "100000000" {
		relation = relation.Where("`duration` between ? and ?", minTime, maxTime)
	}
	if len(category) != 0 {
		genreIDs := []uint{}
		genres := []model.Genre{}
		db.Where("`genre` IN (?)", category).Find(&genres)
		for _, g := range genres {
			genreIDs = append(genreIDs, g.ID)
		}
		relation = relation.Table("tracks").Joins("join track_artists on `track_artists`.`track_id` = `tracks`.`id`")
		relation = relation.Joins("join artist_genres on `track_artists`.`artist_id` = `artist_genres`.`artist_id`").Where("`artist_genres`.`genre_id` IN (?)", genreIDs)
	}
	if around != "" {
		relation = relation.Order(gorm.Expr("abs(`duration` - ?) asc", around)).Order("times desc")
	} else {
		relation = relation.Order("times desc")
	}
	result := relation.First(&model.Track{})
	c.JSON(200, result)
}
