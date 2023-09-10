package client

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Marcus-hayes/media-lookup/constants"
	tmdb "github.com/cyruzin/golang-tmdb"
)

var (
	MissingEnvironmentVarError = "Missing environment variable error: Please set the following environment variables and re-build: "
	MediaNotFoundErr           = "Media for ID #%s was not found. Please try another ID"
)

type tmdbClient struct {
	client *tmdb.Client
}

/*
	PrepareClient: Preparation and configuration function for the TMDB Client. Verifies & retrieves TMDB API key from local environment, initializes client
	for TMDB queries
	https://developer.themoviedb.org/docs
*/
func PrepareClient() (*tmdbClient, error) {
	// Look-up TMDB API Key in environment
	apiKey, ok := os.LookupEnv("TMDB_API_KEY")
	if !ok {
		return nil, fmt.Errorf(MissingEnvironmentVarError)
	}
	log.Println("API key found, initializing client...")
	baseClient, err := tmdb.Init(apiKey)
	if err != nil {
		return nil, err
	}

	// OPTIONAL (Recommended): Enabling auto retry functionality.
	// This option will retry if the previous request fail (429 TOO MANY REQUESTS).
	baseClient.SetClientAutoRetry()

	// OPTIONAL: Set an alternate base URL if you have problems with the default one.
	// Use https://api.tmdb.org/3 instead of https://api.themoviedb.org/3.
	baseClient.SetAlternateBaseURL()

	// OPTIONAL: Setting a custom config for the http.Client.
	// The default timeout is 10 seconds. Here you can set other
	// options like Timeout and Transport.
	customClient := http.Client{
		Timeout: time.Second * 5,
		Transport: &http.Transport{
			MaxIdleConns:    10,
			IdleConnTimeout: 15 * time.Second,
		},
	}

	baseClient.SetClientConfig(customClient)
	parentClient := tmdbClient{
		client: baseClient,
	}
	return &parentClient, nil
}

/*
	MultimediaSearch: Performs multi-media search via TMDB API. Takes query string and url parameter map, perform query, de-paginates results and
	returns in a slice, along with any errors that may have occurred
*/
func (t *tmdbClient) MultimediaSearch(query string, urlOpts map[string]string) ([]constants.TMDBResult, error) {
	var resultSlc []constants.TMDBResult
	log.Printf("Getting results for page #1...")
	resp, err := t.client.GetSearchMulti(query, urlOpts)
	if err != nil {
		return nil, err
	}
	for _, item := range resp.Results {
		resultSlc = append(resultSlc, item)
	}
	for i := 2; i < int(resp.TotalPages); i++ {
		log.Printf("Getting results for page #%d...", i)
		urlOpts["page"] = fmt.Sprintf("%d", i)
		resp, err := t.client.GetSearchMulti(query, urlOpts)
		if err != nil {
			return nil, err
		}
		for _, item := range resp.Results {
			resultSlc = append(resultSlc, item)
		}
	}
	return resultSlc, err
}

/*
	GetDetails: Performs multi-media search via TMDB API. Takes query string and url parameter map, perform query, de-paginates results and
	returns in a slice, along with any errors that may have occurred
*/
func (t *tmdbClient) GetDetails(idStr string, mediaType string) (*constants.TMDBDetailResult, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}
	urlOpts := map[string]string{
		"language": "en-US",
	}
	var result constants.TMDBDetailResult
	switch mediaType {
	case "person":
		resp, err := t.client.GetPersonDetails(id, urlOpts)
		if err != nil {
			return nil, err
		}
		result.PersonDetails = resp
	case "tv":
		resp, err := t.client.GetTVDetails(id, urlOpts)
		if err != nil {
			return nil, err
		}
		result.ShowDetails = resp
	case "movie":
		resp, err := t.client.GetMovieDetails(id, urlOpts)
		if err != nil {
			return nil, err
		}
		result.MovieDetails = resp
	default:
		return nil, fmt.Errorf(MediaNotFoundErr)
	}
	if err != nil {
		return nil, err
	}
	return &result, err
}
