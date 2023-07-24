package lavasearch

import (
	"github.com/disgoorg/disgolink/v3/lavalink"
)

type SearchType string

const (
	SearchTypeTrack    SearchType = "track"
	SearchTypeAlbum    SearchType = "album"
	SearchTypeArtist   SearchType = "artist"
	SearchTypePlaylist SearchType = "playlist"
	SearchTypeText     SearchType = "text"
)

type SearchResult struct {
	Albums    []Album    `json:"albums"`
	Artists   []Artist   `json:"artists"`
	Playlists []Playlist `json:"playlists"`
	Tracks    []Track    `json:"tracks"`
	Texts     []Text     `json:"texts"`
}

type Album struct {
	ID         string  `json:"identifier"`
	Name       string  `json:"name"`
	Artist     string  `json:"artist"`
	URL        string  `json:"url"`
	TrackCount int     `json:"trackCount"`
	ArtworkURL *string `json:"artworkUrl"`
	ISRC       *string `json:"isrc"`
}

type Artist struct {
	ID         string  `json:"identifier"`
	Name       string  `json:"name"`
	URL        string  `json:"url"`
	ArtworkURL *string `json:"artworkUrl"`
}

type Playlist struct {
	ID         string  `json:"identifier"`
	Name       string  `json:"name"`
	URL        string  `json:"url"`
	TrackCount int     `json:"trackCount"`
	ArtworkURL *string `json:"artworkUrl"`
}

type Track struct {
	Title      string            `json:"title"`
	Author     string            `json:"author"`
	Length     lavalink.Duration `json:"length"`
	ID         string            `json:"identifier"`
	IsStream   bool              `json:"isStream"`
	URI        string            `json:"uri"`
	ArtworkURL *string           `json:"artworkUrl"`
	ISRC       *string           `json:"isrc"`
}

type Text struct {
	Text string `json:"text"`
}
