package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mirito333/spotify-time-search-api/api/handler"
)

func Init() {
  r := router()
 
  r.Run(":3000")
}
 

func router() *gin.Engine {
  r := gin.Default()
  r.Use(CORS())

  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "this is api for time search spotify",
    })
  })

  r.GET("/track", handler.SaveTrack)
  r.GET("/playlist", handler.SavePlayList)
  r.GET("/search", handler.GetTrack)
  return r
}


func CORS() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
    c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
    c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
 
    if c.Request.Method == "OPTIONS" {
      c.AbortWithStatus(http.StatusNoContent)
      return
    }
 
    c.Next()
  }
}