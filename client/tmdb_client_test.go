package client

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Set required env vars
	os.Setenv("TMDB_API_KEY", "testAPIKey")
	code := m.Run()
	os.Exit(code)
}

func TestMultimediaSearch_Success(t *testing.T) {
	// Place-holder for tests
}

func TestMultimediaSearch_Failure(t *testing.T) {
	// Place-holder for tests
}
