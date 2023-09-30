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
	// Default call arguments
	err := PerformSearch("example_query", false, "en", int32(1))
	if err != nil {
		t.Error(err)
	}
}

func TestPerformSearch_ParseFailure(t *testing.T) {
	client.MockMultimediaSearchResponse[0].MediaType = "non-existent-type"
	expErr := "error #0: result media type does not match known types: type is non-existent-type\n"
	// Default call arguments
	err := PerformSearch("example_query", false, "en", int32(1))
	if err.Error() != expErr {
		t.Errorf("Unexpected error thrown: Got '%s', but expected '%s'\n", err.Error(), expErr)
	}
}

func TestPerformSearch_SearchFailure(t *testing.T) {
	tmdbClient = &client.MockTMDBClient{MultimediaSearchErr: true}
	expErr := "error calling TMDB client: mock error during MultimediaSearch()\n"
	// Default call arguments
	err := PerformSearch("example_query", false, "en", int32(1))
	if err.Error() != expErr {
		t.Errorf("Unexpected error thrown: Got '%s', but expected '%s'\n", err.Error(), expErr)
	}
	tmdbClient = &client.MockTMDBClient{}
}
