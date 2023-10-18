package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Marcus-hayes/media-lookup/constants"
)

var lotrUrl = "https://the-one-api.dev/v2"

type LOTRClient struct{}

func validateEnv() (*string, error) {
	// Look-up LOTR API Key in environment
	key, ok := os.LookupEnv("LOTR_API_KEY")
	if !ok {
		return nil, fmt.Errorf(MissingEnvironmentVarError, "LOTR_API_KEY")
	}
	log.Println("LOTR API key found, environment verified...")
	return &key, nil
}

func performHTTPRequest(url string, params map[string]string, method string, body []byte) ([]byte, error) {
	token, err := validateEnv()
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("limit", "100")
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", *token))
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBytes, nil
}

func (l *LOTRClient) GetBooks() (*constants.LOTRBookResponse, error) {
	url := fmt.Sprintf("%s/book", lotrUrl)
	respBytes, err := performHTTPRequest(url, nil, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	var responseObj constants.LOTRBookResponse
	err = json.Unmarshal(respBytes, &responseObj)
	if err != nil {
		return nil, err
	}
	return &responseObj, nil
}

func (l *LOTRClient) GetBookById(id string) (*constants.LOTRBookResponse, error) {
	url := fmt.Sprintf("%s/book/%s", lotrUrl, id)
	respBytes, err := performHTTPRequest(url, nil, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	var responseObj constants.LOTRBookResponse
	err = json.Unmarshal(respBytes, &responseObj)
	if err != nil {
		return nil, err
	}
	return &responseObj, nil
}

func (l *LOTRClient) GetBookChaptersById(id string) (*constants.LOTRBookResponse, error) {
	url := fmt.Sprintf("%s/book/%s/chapter", lotrUrl, id)
	respBytes, err := performHTTPRequest(url, nil, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	var responseObj constants.LOTRBookResponse
	err = json.Unmarshal(respBytes, &responseObj)
	if err != nil {
		return nil, err
	}
	return &responseObj, nil
}

func (l *LOTRClient) GetMovies() (*constants.LOTRMovieResponse, error) {
	url := fmt.Sprintf("%s/movie", lotrUrl)
	respBytes, err := performHTTPRequest(url, nil, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	var responseObj constants.LOTRMovieResponse
	err = json.Unmarshal(respBytes, &responseObj)
	if err != nil {
		return nil, err
	}
	return &responseObj, nil
}

func (l *LOTRClient) GetMovieById(id string) (*constants.LOTRMovieResponse, error) {
	url := fmt.Sprintf("%s/movie/%s", lotrUrl, id)
	respBytes, err := performHTTPRequest(url, nil, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	var responseObj constants.LOTRMovieResponse
	err = json.Unmarshal(respBytes, &responseObj)
	if err != nil {
		return nil, err
	}
	return &responseObj, nil
}

func (l *LOTRClient) GetMovieQuotesById(id string) (*constants.LOTRMovieQuoteResponse, error) {
	url := fmt.Sprintf("%s/movie/%s/quote", lotrUrl, id)
	respBytes, err := performHTTPRequest(url, nil, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	var responseObj constants.LOTRMovieQuoteResponse
	err = json.Unmarshal(respBytes, &responseObj)
	if err != nil {
		return nil, err
	}
	return &responseObj, nil
}
