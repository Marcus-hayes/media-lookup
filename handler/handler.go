package handler

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/Marcus-hayes/media-lookup/client"
	"github.com/Marcus-hayes/media-lookup/constants"
)

var (
	templateParseErr = "error parsing TMDB client response to template: %s\n"
	templateExecErr  = "error executing template: %s\n"
	tmdbClientErr    = "error calling TMDB client: %s\n"
	tmdbClient       constants.TMDBClient
)

/*
	Init: Initializes TMDB client instance, logic as a separate function to allow mocking client
*/
func Init() {
	t, err := client.PrepareClient()
	if err != nil {
		log.Fatalln("Error initializing handler: ", err)
	}
	tmdbClient = t
}

/*
	parseTMDBResults: Templatizes input results based on their media type and logs the results to console. Returns error if one occurs, nil otherwise
*/
func parseTMDBResults(results []constants.TMDBResult) []error {
	var errSlc []error
	for _, result := range results {
		switch result.MediaType {
		case "person":
			t, err := template.New(result.MediaType).Parse(constants.TMDBPersonTemplate)
			if err != nil {
				errSlc = append(errSlc, fmt.Errorf(templateParseErr, err))
				continue
			}
			err = t.Execute(os.Stdout, result)
			if err != nil {
				errSlc = append(errSlc, fmt.Errorf(templateExecErr, err))
				continue
			}
		case "movie":
			t, err := template.New(result.MediaType).Parse(constants.TMDBMovieTemplate)
			if err != nil {
				errSlc = append(errSlc, fmt.Errorf(templateParseErr, err))
				continue
			}
			err = t.Execute(os.Stdout, result)
			if err != nil {
				errSlc = append(errSlc, fmt.Errorf(templateExecErr, err))
				continue
			}
		case "tv":
			t, err := template.New(result.MediaType).Parse(constants.TMDBShowTemplate)
			if err != nil {
				errSlc = append(errSlc, fmt.Errorf(templateParseErr, err))
				continue
			}
			err = t.Execute(os.Stdout, result)
			if err != nil {
				errSlc = append(errSlc, fmt.Errorf(templateExecErr, err))
				continue
			}
		default:
			err := fmt.Errorf("result media type does not match known types: type is %s", result.MediaType)
			errSlc = append(errSlc, err)
			continue
		}
	}
	return errSlc
}

/*
	PerformSearch: Performs multimedia search via TMDB API using the input query, nsfw, language, and page parameters. Results are templatized and logged to console.
	Returns error if one occurred, nil otherwise
*/
func PerformSearch(query string, nsfw bool, language string, page int32) error {
	urlOpts := map[string]string{
		"include_adult": fmt.Sprintf("%t", nsfw),
		"page":          fmt.Sprintf("%d", page),
		"language":      language,
	}
	results, err := tmdbClient.MultimediaSearch(query, urlOpts)
	if err != nil {
		return fmt.Errorf(tmdbClientErr, err)
	}
	errs := parseTMDBResults(results)
	var errStr string
	if len(errs) > 0 {
		log.Println("Errors occurred while parsing TMDB Results:")
		for i, err := range errs {
			log.Println(err)
			errStr += fmt.Sprintf("error #%d: %s\n", i, err.Error())
		}
	}
	err = fmt.Errorf(errStr)
	if len(err.Error()) > 0 {
		return err
	}
	return nil
}
