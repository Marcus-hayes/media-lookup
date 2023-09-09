package handler

import (
	"fmt"
	"os"
	"text/template"

	"github.com/Marcus-hayes/media-lookup/client"
	"github.com/Marcus-hayes/media-lookup/constants"
)

var (
	templateParseErr = "error parsing TMDB client response to template: %s\n"
	templateExecErr  = "error executing template: %s\n"
)

/*
	parseTMDBResults: Templatizes input results based on their media type and logs the results to console. Returns error if one occurs, nil otherwise
*/
func parseTMDBResults(results []constants.TMDBResult) error {
	for _, result := range results {
		switch result.MediaType {
		case "person":
			t, err := template.New(result.MediaType).Parse(constants.TMDBPersonTemplate)
			if err != nil {
				return fmt.Errorf(templateParseErr, err)
			}
			err = t.Execute(os.Stdout, result)
			if err != nil {
				return fmt.Errorf(templateExecErr, err)
			}
		case "movie":
			t, err := template.New(result.MediaType).Parse(constants.TMDBMovieTemplate)
			if err != nil {
				return fmt.Errorf(templateParseErr, err)
			}
			err = t.Execute(os.Stdout, result)
			if err != nil {
				return fmt.Errorf(templateExecErr, err)
			}
		case "tv":
			t, err := template.New(result.MediaType).Parse(constants.TMDBShowTemplate)
			if err != nil {
				return fmt.Errorf(templateParseErr, err)
			}
			err = t.Execute(os.Stdout, result)
			if err != nil {
				return fmt.Errorf(templateExecErr, err)
			}
		default:
			return fmt.Errorf("Result media type does not match known types: Type is %s", result.MediaType)
		}
	}
	return nil
}

/*
	PerformSearch: Performs multimedia search via TMDB API using the input query, nsfw, language, and page parameters. Results are templatized and logged to console.
	Returns error if one occurred, nil otherwise
*/
func PerformSearch(query string, nsfw bool, language string, page int32) error {
	tmdbClient, err := client.PrepareClient()
	if err != nil {
		return fmt.Errorf("error preparing TMDB client: %s/n", err)
	}
	urlOpts := map[string]string{
		"include_adult": fmt.Sprintf("%t", nsfw),
		"page":          fmt.Sprintf("%d", page),
		"language":      language,
	}
	results, err := tmdbClient.MultimediaSearch(query, urlOpts)
	if err != nil {
		return fmt.Errorf("error calling client in handler: %s/n", err)
	}
	err = parseTMDBResults(results)
	if err != nil {
		return fmt.Errorf("error calling client in handler: %s/n", err)
	}
	return nil
}
