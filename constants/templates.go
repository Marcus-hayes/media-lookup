package constants

var (
	TMDBShowTemplate = `
----------------------------------------
{{.Title}}
Media type: {{.MediaType}}
ID: {{.ID}}
Release Date: {{.ReleaseDate}}
Original Language: {{.OriginalLanguage}}
Genre IDs: {{if .GenreIDs}}{{.GenreIDs}}{{- else}}None{{end}}
Description: {{if .Overview}}{{.Overview}}{{- else}}None{{end}}
`
	TMDBMovieTemplate = `
----------------------------------------
{{.Title}}
Media type: {{.MediaType}}
ID: {{.ID}}
Release Date: {{.ReleaseDate}}
Original Language: {{.OriginalLanguage}}
Genre IDs: {{if .GenreIDs}}{{.GenreIDs}}{{- else}}Not Available{{end}}
Description: {{if .Overview}}{{.Overview}}{{- else}}Not Available{{end}}
`
	TMDBPersonTemplate = `
----------------------------------------
{{.Name}}
Media type: {{.MediaType}}
ID: {{.ID}}
Known For: {{if .KnownFor}}{{.KnownFor}}{{- else}}Not Available{{end}}
Genre IDs: {{if .GenreIDs}}{{.GenreIDs}}{{- else}}Not Available{{end}}
Description: {{if .Overview}}{{.Overview}}{{- else}}Not Available{{end}}
`
	// https://pkg.go.dev/github.com/cyruzin/golang-tmdb@v1.5.1#TVDetails
	TMDBShowDetailTemplate = `
----------------------------------------
{{.Name}}
ID: {{.ID}}
Overview: {{if .Overview}}{{.Overview}}{{- else}}Not Available{{end}}
Created By: {{if .CreatedBy}}{{.CreatedBy}}{{- else}}Not Available{{end}}
Production Companies: {{range .ProductionCompanies}}
	Name: {{if .Name}}{{.Name}}{{- else}}Not Available{{end}}
	ID: {{if .ID}}{{.ID}}{{- else}}Not Available{{end}}
	Country of Origin: {{.OriginCountry}}
	----------------------------------------
	{{- else}}Not Available{{end}}
Networks: {{range .Networks}}
	Name: {{if .Name}}{{.Name}}{{- else}}Not Available{{end}}
	ID: {{if .ID}}{{.ID}}{{- else}}Not Available{{end}}
	Country of Origin: {{.OriginCountry}}
	----------------------------------------
	{{- else}}Not Available{{end}}
Genres: {{if .Genres}}{{.Genres}}{{- else}}Not Available{{end}}
Description: {{if .Overview}}{{.Overview}}{{- else}}Not Available{{end}}
First Air Date: {{.FirstAirDate}}
In-Production : {{.InProduction}}
Run-time: {{if .EpisodeRunTime}}{{.EpisodeRunTime}}{{- else}}Not Available{{end}}
Number of Episodes: {{if .NumberOfEpisodes}}{{.NumberOfEpisodes}}{{- else}}Not Available{{end}}
Number of Seasons: {{if .NumberOfSeasons}}{{.NumberOfSeasons}}{{- else}}Not Available{{end}}
Next Episode to Air: {{if .NextEpisodeToAir}}
	Name: {{if .NextEpisodeToAir.Name}}{{.NextEpisodeToAir.Name.Name}}{{- else}}Not Available{{end}}
	ID: {{if .NextEpisodeToAir.ID}}{{.NextEpisodeToAir.ID}}{{- else}}Not Available{{end}}
	Season Number: {{if .NextEpisodeToAir.SeasonNumber}}{{.NextEpisodeToAir.SeasonNumber}}{{- else}}Not Available{{end}}
	Air Date: {{if .NextEpisodeToAir.AirDate}}{{.NextEpisodeToAir.AirDate}}{{- else}}Not Available{{end}}
	Overview: {{if .NextEpisodeToAir.Overview}}{{.NextEpisodeToAir.Overview}}{{- else}}Not Available{{end}}
	----------------------------------------
	{{- else}}Not Available{{end}}
Last Episode to Air: {{if .LastEpisodeToAir}}
	Name: {{if .LastEpisodeToAir.Name}}{{.LastEpisodeToAir.Name}}{{- else}}Not Available{{end}}
	ID: {{if .LastEpisodeToAir.ID}}{{.LastEpisodeToAir.ID}}{{- else}}Not Available{{end}}
	Season Number: {{.LastEpisodeToAir.AirDate}}
	Air Date: {{.LastEpisodeToAir.AirDate}}
	Overview: {{if .LastEpisodeToAir.Overview}}{{.LastEpisodeToAir.Overview}}{{- else}}Not Available{{end}}
	----------------------------------------
	{{- else}}Not Available{{end}}
Seasons: {{range .Seasons}}
	Name: {{if .Name}}{{.Name}}{{- else}}Not Available{{end}}
	ID: {{if .ID}}{{.ID}}{{- else}}Not Available{{end}}
	Season Number: {{if .SeasonNumber}}{{.SeasonNumber}}{{- else}}Not Available{{end}}
	Overview: {{if .Overview}}{{.Overview}}{{- else}}Not Available{{end}}
	{{- else}}Not Available{{end}}
	----------------------------------------
`
	// https://pkg.go.dev/github.com/cyruzin/golang-tmdb@v1.5.1#MovieDetails
	TMDBMovieDetailTemplate = `
----------------------------------------
{{.Title}}
ID: {{.ID}}
Release Date: {{.ReleaseDate}}
Original Language: {{.OriginalLanguage}}
Genres: {{if .Genres}}{{.Genres}}{{- else}}None{{end}}
Production Companies: {{range .ProductionCompanies}}
	Name: {{.Name}}
	ID: {{.ID}}
	Country of Origin: {{.OriginCountry}}
	----------------------------------------
	{{- else}}Not Available{{end}}
Budget: {{if .Budget}}{{.Budget}}{{- else}}Not Available{{end}}
Revenue: {{if .Revenue}}{{.Revenue}}{{- else}}Not Available{{end}}
Runtime: {{if .Runtime}}{{.Runtime}}{{- else}}Not Available{{end}}
Description: {{if .Overview}}{{.Overview}}{{- else}}Not Available{{end}}
`
	// https://pkg.go.dev/github.com/cyruzin/golang-tmdb@v1.5.1#PersonDetails
	TMDBPersonDetailTemplate = `
----------------------------------------
{{.Name}}
ID: {{.ID}}
Also Known As: {{if .AlsoKnownAs}}{{.AlsoKnownAs}}{{- else}}None{{end}}
Known For Department: {{if .KnownForDepartment}}{{.KnownForDepartment}}{{- else}}Not Available{{end}}
Place of Birth: {{if .PlaceOfBirth}}{{.PlaceOfBirth}}{{- else}}Not Available{{end}}
Birthday: {{if .Birthday}}{{.Birthday}}{{- else}}Not Available{{end}}
Date of Death: {{if .Deathday}}{{.Deathday}}{{- else}}Not Available{{end}}
Bio: {{if .Biography}}{{.Biography}}{{- else}}Not Available{{end}}
`
)
