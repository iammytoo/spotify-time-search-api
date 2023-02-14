package spotify

import (
	"context"
	"log"
	"os"

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

func (s *Session) GetTrack(id string) model.TrackData {
	track, err := s.client.GetTrack(s.ctx, spotify.ID(id))
	if err != nil {
		log.Fatal(err)
	}
	return s.getTrackInfo(track)
}

func (s *Session) GetPlaylist(id string) (model.PlayList, []model.TrackData) {
	playlist, err := s.client.GetPlaylist(s.ctx, spotify.ID(id))
	if err != nil {
		log.Fatal(err)
	}
	tracks := []model.TrackData{}
	playlistTracks := playlist.Tracks
	for _, t := range playlistTracks.Tracks {
		track := s.getTrackInfo(&t.Track)
		tracks = append(tracks, track)
	}
	return model.PlayList{Name: playlist.Name, Key: playlist.ID.String(), IsFetched: true}, tracks
}

func (s *Session) getTrackInfo(track *spotify.FullTrack) model.TrackData {
	result := model.TrackData{}
	result.Track = model.Track{Name: track.Name, Key: track.ID.String(), Duration: track.Duration, Times: track.Popularity}
	result.TrackKey = track.ID.String()
	for _, a := range track.Artists {
		artist, err := s.client.GetArtist(s.ctx, a.ID)
		if err != nil {
			log.Fatal(err)
		}
		result.Artists = append(result.Artists, model.Artist{Name: artist.Name,Key: artist.ID.String()})
		result.ArtistKeys = append(result.ArtistKeys, artist.ID.String())
		artgenekey := model.ArtistGenreKey{ArtistKey: artist.ID.String()}
		for _, g := range artist.Genres {
			result.Genres = append(result.Genres, model.Genre{Genre: g})
			artgenekey.GenreKeys = append(artgenekey.GenreKeys, g)
		}
		result.AGKeys = append(result.AGKeys, artgenekey)
	}
	return result
}
