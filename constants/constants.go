package constants

const (
	TMDBShowTemplate = `
	----------------------------------------
	{{.Title}}
	Media type: {{.MediaType}}
	ID: {{.ID}}
	Release Date: {{.ReleaseDate}}
	Original Language: {{.OriginalLanguage}}
	Genre IDs: {{if .GenreIDs}}{{.GenreIDs}}{{else}}None{{end}}
	Description: {{if .Overview}}{{.Overview}}{{else}}None{{end}}
	`
	TMDBMovieTemplate = `
	----------------------------------------
	{{.Title}}
	Media type: {{.MediaType}}
	ID: {{.ID}}
	Release Date: {{.ReleaseDate}}
	Original Language: {{.OriginalLanguage}}
	Genre IDs: {{if .GenreIDs}}{{.GenreIDs}}{{else}}None{{end}}
	Description: {{if .Overview}}{{.Overview}}{{else}}None{{end}}
	`
	TMDBPersonTemplate = `
	----------------------------------------
	{{.Name}}
	Media type: {{.MediaType}}
	ID: {{.ID}}
	Known For: {{if .KnownFor}}{{.KnownFor}}{{else}}None{{end}}
	Genre IDs: {{if .GenreIDs}}{{.GenreIDs}}{{else}}None{{end}}
	Description: {{if .Overview}}{{.Overview}}{{else}}None{{end}}
	`
)

type TMDBClient interface {
	MultimediaSearch(query string, urlOpts map[string]string) ([]TMDBResult, error)
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
