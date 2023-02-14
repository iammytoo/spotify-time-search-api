package model

type TrackData struct {
	Track      Track
	Artists    []Artist
	Genres     []Genre
	TrackKey   string
	ArtistKeys []string
	AGKeys     []ArtistGenreKey
}

type ArtistGenreKey struct {
	ArtistKey string
	GenreKeys []string
}
