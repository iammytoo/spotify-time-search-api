package main

import (
	"github.com/joho/godotenv"
	"github.com/mirito333/spotify-time-search-api/api/connection"
	"github.com/mirito333/spotify-time-search-api/api/router"
	"github.com/mirito333/spotify-time-search-api/api/spotify"
)

func main() {
	godotenv.Load(".env")
	spotify.Init()
	connection.Init()
	router.Init()
}
