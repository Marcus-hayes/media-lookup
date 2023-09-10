package client

import (
	"fmt"

	"github.com/Marcus-hayes/media-lookup/constants"
)

// MockTMDBClient - Mock client for use in unit testing
type MockTMDBClient struct {
	MultimediaSearchErr bool
}

var MockMultimediaSearchResponse = []constants.TMDBResult{
	// Person
	{
		ID:        1,
		Name:      "Tony Montana",
		MediaType: "person",
	},
	// Movie
	{
		ID:        2,
		Title:     "Scarface",
		MediaType: "movie",
	},
	// TV Show
	{
		ID:        3,
		Title:     "Narcos",
		MediaType: "tv",
	},
}

func (m *MockTMDBClient) MultimediaSearch(query string, urlOpts map[string]string) ([]constants.TMDBResult, error) {
	if m.MultimediaSearchErr {
		return nil, fmt.Errorf("mock error during MultimediaSearch()")
	}
	return MockMultimediaSearchResponse, nil
}

func (m *MockTMDBClient) GetDetails(idStr string, mediaType string) (*constants.TMDBDetailResult, error) {
	return nil, nil
}
