package constants

var (
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
	// https://pkg.go.dev/github.com/cyruzin/golang-tmdb@v1.5.1#TVDetails
	TMDBShowDetailTemplate = `
----------------------------------------
{{.Name}}
ID: {{.ID}}
Created By: {{if .CreatedBy}}{{.CreatedBy}}{{else}}Not Available{{end}}
Production Companies: {{if .ProductionCompanies}}{{.ProductionCompanies}}{{else}}Not Available{{end}}
Networks: {{if .Networks}}{{.Networks}}{{else}}Not Available{{end}}
Genres: {{if .Genres}}{{.Genres}}{{else}}None{{end}}
Description: {{if .Overview}}{{.Overview}}{{else}}None{{end}}
Run-time: {{if .EpisodeRunTime}}{{.EpisodeRunTime}}{{else}}Not Available{{end}}
Number of Episodes: {{if .NumberOfEpisodes}}{{.NumberOfEpisodes}}{{else}}Not Available{{end}}
Number of Seasons: {{if .NumberOfSeasons}}{{.NumberOfSeasons}}{{else}}Not Available{{end}}
Seasons: {{if .Seasons}}{{.Seasons}}{{else}}Not Available{{end}}
First Air Date: {{.FirstAirDate}}
In-Production : {{.InProduction}}
{{if .InProduction}}Next Episode to Air: {{if .NextEpisodeToAir.Name}}{{.NextEpisodeToAir}}{{else}}Not Available{{end}}{{else}}Last Air Date: {{.LastAirDate}}
Last Episode to Air: {{if .LastEpisodeToAir}}{{.LastEpisodeToAir}}{{else}}None{{end}}{{end}}
`
	// https://pkg.go.dev/github.com/cyruzin/golang-tmdb@v1.5.1#MovieDetails
	TMDBMovieDetailTemplate = `
----------------------------------------
{{.Title}}
ID: {{.ID}}
Release Date: {{.ReleaseDate}}
Original Language: {{.OriginalLanguage}}
Genres: {{if .Genres}}{{.Genres}}{{else}}None{{end}}
Production Companies: {{if .ProductionCompanies}}{{.ProductionCompanies}}{{else}}Not Available{{end}}
Budget: {{if .Budget}}{{.Budget}}{{else}}Not Available{{end}}
Revenue: {{if .Revenue}}{{.Revenue}}{{else}}Not Available{{end}}
Runtime: {{if .Runtime}}{{.Runtime}}{{else}}Not Available{{end}}
Description: {{if .Overview}}{{.Overview}}{{else}}None{{end}}
`
	// https://pkg.go.dev/github.com/cyruzin/golang-tmdb@v1.5.1#PersonDetails
	TMDBPersonDetailTemplate = `
----------------------------------------
{{.Name}}
ID: {{.ID}}
Also Known As: {{if .AlsoKnownAs}}{{.AlsoKnownAs}}{{else}}None{{end}}
Known For Department: {{if .KnownForDepartment}}{{.KnownForDepartment}}{{else}}None{{end}}
Place of Birth: {{if .PlaceOfBirth}}{{.PlaceOfBirth}}{{else}}Not Available{{end}}
Birthday: {{if .Birthday}}{{.Birthday}}{{else}}Not Available{{end}}
{{if .Deathday}}Date of Death: {{.Deathday}}{{else}}{{end}}
Bio: {{if .Biography}}{{.Biography}}{{else}}None{{end}}
`
)
