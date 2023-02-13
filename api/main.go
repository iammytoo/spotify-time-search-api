package main

import (
	"github.com/mirito333/spotify-time-search-api/api/connection"
	"github.com/mirito333/spotify-time-search-api/api/router"
)

func main() {
	connection.Init()
	router.Init()
}
