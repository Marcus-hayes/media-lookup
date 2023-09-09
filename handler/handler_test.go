package handler

import (
	"os"
	"testing"

	"github.com/Marcus-hayes/media-lookup/client"
)

func TestMain(m *testing.M) {
	// Set required env var
	os.Setenv("TMDB_API_KEY", "testAPIKey")
	// Mock out TMDB Client
	tmdbClient = &client.MockTMDBClient{}
	code := m.Run()
	os.Exit(code)
}

func TestPerformSearch_Success(t *testing.T) {
	//Place-holder text
}

func TestPerformSearch_ParseFailure(t *testing.T) {
	//Place-holder text
}

func TestPerformSearch_SearchFailure(t *testing.T) {
	//Place-holder text
}
