package main

import (
	"github.com/joho/godotenv"
	"github.com/mirito333/spotify-time-search-api/api/connection"
	"github.com/mirito333/spotify-time-search-api/api/router"
)

func main() {
	godotenv.Load(".env")
	connection.Init()
	router.Init()
}
