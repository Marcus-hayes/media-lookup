package handler

import (
	"fmt"

	"github.com/Marcus-hayes/media-lookup/client"
)

func PerformSearch(query string, nsfw string, language string) (interface{}, error) {
	tmdbClient, err := client.PrepareClient()
	if err != nil {
		return nil, fmt.Errorf("error preparing TMDB client: %s/n", err)
	}
	urlOpts := client.URLOptions{
		IncludeAdult: nsfw,
		Page:         1,
		Language:     language,
	}
	searchResult, err := tmdbClient.MultimediaSearch(query, urlOpts)
	return searchResult, nil
}
