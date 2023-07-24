package lavasearch

import (
	"net/http"
	"net/url"

	"github.com/disgoorg/disgolink/v3/disgolink"
	"github.com/disgoorg/json"
)

func LoadSearch(client disgolink.RestClient, query string, types []SearchType) (*SearchResult, error) {
	values := url.Values{}
	values.Set("query", query)
	for _, t := range types {
		values.Add("types", string(t))
	}

	rq, err := http.NewRequest(http.MethodGet, "/v4/loadsearch?"+values.Encode(), nil)
	if err != nil {
		return nil, err
	}
	rs, err := client.Do(rq)
	if err != nil {
		return nil, err
	}

	defer rs.Body.Close()
	var result SearchResult
	if err = json.NewDecoder(rs.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
