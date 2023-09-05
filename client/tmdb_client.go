package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	tmdb "github.com/cyruzin/golang-tmdb"
)

var MissingEnvironmentVarError = "Missing environment variable error: Please set the following environment variables and re-build: "

/*
	Init: Preparation and configuration function for the TMDB Client. Verifies & retrieves TMDB API key from local environment, initializes client
	for TMDB queries
	https://developer.themoviedb.org/docs
*/

type tmdbClient struct {
	client *tmdb.Client
}

type URLOptions struct {
	IncludeAdult string `json:"include_adult"`
	Language     string `json:"language"`
	Page         int    `json:"page"`
}

func PrepareClient() (*tmdbClient, error) {
	// Look-up TMDB API Key in environment
	apiKey, ok := os.LookupEnv("TMDB_API_KEY")
	if !ok {
		return nil, fmt.Errorf(MissingEnvironmentVarError)
	}

	baseClient, err := tmdb.Init(apiKey)
	if err != nil {
		log.Println(err)
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

func (t *tmdbClient) MultimediaSearch(query string, urlOpts URLOptions) (*tmdb.SearchMulti, error) {
	var urlMap map[string]string
	data, err := json.Marshal(urlOpts)
	err = json.Unmarshal(data, &urlMap)
	resp, err := t.client.GetSearchMulti(query, urlMap)
	return resp, err
}
