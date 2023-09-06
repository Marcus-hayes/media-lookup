package client

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	tmdb "github.com/cyruzin/golang-tmdb"
)

var MissingEnvironmentVarError = "Missing environment variable error: Please set the following environment variables and re-build: "

type Result struct {
	PosterPath       string   `json:"poster_path,omitempty"`
	Popularity       float32  `json:"popularity"`
	ID               int64    `json:"id"`
	Overview         string   `json:"overview,omitempty"`
	BackdropPath     string   `json:"backdrop_path,omitempty"`
	VoteAverage      float32  `json:"vote_average,omitempty"`
	MediaType        string   `json:"media_type"`
	FirstAirDate     string   `json:"first_air_date,omitempty"`
	OriginCountry    []string `json:"origin_country,omitempty"`
	GenreIDs         []int64  `json:"genre_ids,omitempty"`
	OriginalLanguage string   `json:"original_language,omitempty"`
	VoteCount        int64    `json:"vote_count,omitempty"`
	Name             string   `json:"name,omitempty"`
	OriginalName     string   `json:"original_name,omitempty"`
	Adult            bool     `json:"adult,omitempty"`
	ReleaseDate      string   `json:"release_date,omitempty"`
	OriginalTitle    string   `json:"original_title,omitempty"`
	Title            string   `json:"title,omitempty"`
	Video            bool     `json:"video,omitempty"`
	ProfilePath      string   `json:"profile_path,omitempty"`
	KnownFor         []struct {
		PosterPath       string  `json:"poster_path"`
		Adult            bool    `json:"adult"`
		Overview         string  `json:"overview"`
		ReleaseDate      string  `json:"release_date"`
		OriginalTitle    string  `json:"original_title"`
		GenreIDs         []int64 `json:"genre_ids"`
		ID               int64   `json:"id"`
		MediaType        string  `json:"media_type"`
		OriginalLanguage string  `json:"original_language"`
		Title            string  `json:"title"`
		BackdropPath     string  `json:"backdrop_path"`
		Popularity       float32 `json:"popularity"`
		VoteCount        int64   `json:"vote_count"`
		Video            bool    `json:"video"`
		VoteAverage      float32 `json:"vote_average"`
	} `json:"known_for,omitempty"`
}

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

func (t *tmdbClient) MultimediaSearch(query string, urlOpts map[string]string) ([]Result, error) {
	var resultSlc []Result
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
