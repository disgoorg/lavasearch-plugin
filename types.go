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
	Tracks    []lavalink.Track    `json:"tracks"`
	Albums    []lavalink.Playlist `json:"albums"`
	Artists   []lavalink.Playlist `json:"artists"`
	Playlists []lavalink.Playlist `json:"playlists"`
	Texts     []Text              `json:"texts"`
}

type Text struct {
	Text   string              `json:"text"`
	Plugin lavalink.PluginInfo `json:"pluginInfo"`
}
