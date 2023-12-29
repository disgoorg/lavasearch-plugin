package lavasearch

import (
	"context"
	"errors"
	"net/http"
	"net/url"

	"github.com/disgoorg/disgolink/v3/disgolink"
	"github.com/disgoorg/disgolink/v3/lavalink"
	"github.com/disgoorg/json"
)

var (
	ErrEmptySearchResult = errors.New("empty search result")
)

// LoadSearch loads a search result from Lavalink & returns a *SearchResult or ErrEmptySearchResult if no results were found or lavalink.Error if something went wrong with Lavalink or any other error.
func LoadSearch(ctx context.Context, client disgolink.RestClient, query string, types []SearchType) (*SearchResult, error) {
	values := url.Values{}
	values.Set("query", query)
	for _, t := range types {
		values.Add("types", string(t))
	}

	rq, err := http.NewRequestWithContext(ctx, http.MethodGet, "/v4/loadsearch?"+values.Encode(), nil)
	if err != nil {
		return nil, err
	}
	rs, err := client.Do(rq)
	if err != nil {
		return nil, err
	}

	if rs.StatusCode == http.StatusNoContent {
		return nil, ErrEmptySearchResult
	} else if rs.StatusCode != http.StatusOK {
		var lavalinkError lavalink.Error
		if err = json.NewDecoder(rs.Body).Decode(&lavalinkError); err != nil {
			return nil, err
		}
		return nil, lavalinkError
	}

	defer rs.Body.Close()
	var result SearchResult
	if err = json.NewDecoder(rs.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
