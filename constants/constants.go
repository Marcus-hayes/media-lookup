package constants

import (
	tmdb "github.com/cyruzin/golang-tmdb"
)

type TMDBClient interface {
	MultimediaSearch(query string, urlOpts map[string]string) ([]TMDBResult, error)
	GetDetails(idStr string, mediaType string) (*TMDBDetailResult, error)
}

type LOTRClient interface {
	GetBooks() error
}

type TMDBResult struct {
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

type TMDBDetailResult struct {
	PersonDetails *tmdb.PersonDetails
	MovieDetails  *tmdb.MovieDetails
	ShowDetails   *tmdb.TVDetails
}

type GuessAgeResult struct {
	Count int32  `json:"count"`
	Name  string `json:"name"`
	Age   int32  `json:"age"`
}

type LOTRBookResponse struct {
	Docs   []LOTRBookItem `json:"docs"`
	Total  int32          `json:"total"`
	Limit  int32          `json:"limit"`
	Offset int32          `json:"offset"`
	Page   int32          `json:"page"`
	Pages  int32          `json:"pages"`
}

type LOTRBookItem struct {
	ID          string `json:"_id"`
	Name        string `json:"name,omitempty"`
	ChapterName string `json:"chapterName,omitempty"`
}

type LOTRMovieResponse struct {
	Docs   []LOTRMovieItem `json:"docs"`
	Total  int32           `json:"total"`
	Limit  int32           `json:"limit"`
	Offset int32           `json:"offset"`
	Page   int32           `json:"page"`
	Pages  int32           `json:"pages"`
}

type LOTRMovieItem struct {
	ID                      string  `json:"_id"`
	Name                    string  `json:"name"`
	RottenTomatoesScore     float32 `json:"rottenTomatoesScore"`
	AcademyAwardNominations int     `json:"academyAwardNominations"`
	AcademyAwardWins        int     `json:"academyAwardWins"`
	Runtime                 int     `json:"runtimeInMinutes"`
	BoxOfficeRevenue        float32 `json:"boxOfficeRevenueInMillions"`
	Budget                  float32 `json:"BudgetInMillions"`
}

type LOTRMovieQuoteResponse struct {
	Docs   []LOTRMovieQuoteItem `json:"docs"`
	Total  int32                `json:"total"`
	Limit  int32                `json:"limit"`
	Offset int32                `json:"offset"`
	Page   int32                `json:"page"`
	Pages  int32                `json:"pages"`
}

type LOTRMovieQuoteItem struct {
	ID          string `json:"_id"`
	QuoteId     string `json:"id"`
	CharacterId string `json:"character"`
	Dialog      string `json:"dialog"`
	MovieId     string `json:"movie"`
}
