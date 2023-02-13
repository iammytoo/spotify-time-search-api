package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mirito333/spotify-time-search-api/api/model"
	spotifyauth "github.com/zmb3/spotify/v2/auth"

	"golang.org/x/oauth2/clientcredentials"

	"github.com/zmb3/spotify/v2"
)

type Session struct {
	client *spotify.Client
	ctx    context.Context
}

var session Session

func Init() {
	session.ctx = context.Background()
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL:     spotifyauth.TokenURL,
	}
	token, err := config.Token(session.ctx)
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	httpClient := spotifyauth.New().Client(session.ctx, token)
	session.client = spotify.New(httpClient)
}

func GetSession() Session {
	return session
}

func (s *Session) GetTrackInfo(id string) model.Track {
	track, err := s.client.GetTrack(s.ctx, spotify.ID(id))
	if err != nil {
		log.Fatal(err)
	}

	return model.Track{Name: track.Name, Key: id, Duration: track.Duration, Times: track.Popularity}
}

func (s *Session) GetPlaylist(id string) (model.PlayList, []model.Track) {
	playlist, err := s.client.GetPlaylist(s.ctx, spotify.ID(id))
	playlistInfo := model.PlayList{Name: playlist.Name, Key: playlist.ID.String(), IsFetched: true}
	if err != nil {
		log.Fatal(err)
	}
	tracks := []model.Track{}
	playlistTracks := playlist.Tracks
	for _, t := range playlistTracks.Tracks {
		track := model.Track{Name: t.Track.Name, Key: t.Track.ID.String(), Duration: t.Track.Duration, Times: t.Track.Popularity}
		tracks = append(tracks, track)
	}
	return playlistInfo, tracks
}
func main() {
	godotenv.Load(".env")
	Init()
	session := GetSession()
	list, tracks := session.GetPlaylist("37i9dQZF1DXayDMsJG9ZBv")
	fmt.Println(list.Name + "," + list.Key)
	for _, t := range tracks {
		fmt.Println(t.Name, t.Key, t.Duration, t.Times)
	}
}
