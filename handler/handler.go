package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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
	InitTMDB: Initializes TMDB client instance, logic as a separate function to allow mocking client
*/
func InitTMDB() {
	t, err := client.PrepareTMDBClient()
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

/*
	GetTMDBDetails: Performs multimedia search via TMDB API using the input id and page parameters. Results are templatized and logged to console.
	Returns error if one occurred, nil otherwise
*/
func GetTMDBDetails(id string, mediaType string) error {
	result, err := tmdbClient.GetDetails(id, mediaType)
	if err != nil {
		return fmt.Errorf(tmdbClientErr, err)
	}
	if result.MovieDetails != nil {
		t, err := template.New(mediaType).Parse(constants.TMDBMovieDetailTemplate)
		if err != nil {
			return err
		}
		err = t.Execute(os.Stdout, *result.MovieDetails)
		if err != nil {
			return err
		}
	} else if result.PersonDetails != nil {
		t, err := template.New(mediaType).Parse(constants.TMDBPersonDetailTemplate)
		if err != nil {
			return err
		}
		err = t.Execute(os.Stdout, *result.PersonDetails)
		if err != nil {
			return err
		}
	} else if result.ShowDetails != nil {
		t, err := template.New(mediaType).Parse(constants.TMDBShowDetailTemplate)
		if err != nil {
			return err
		}
		err = t.Execute(os.Stdout, *result.ShowDetails)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no metadata returned for that ID")
	}
	return nil
}

func GuessAgeByName(name string) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.agify.io", nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("name", name)
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBytes, _ := io.ReadAll(resp.Body)
	var resultTemplate constants.GuessAgeResult
	err = json.Unmarshal(respBytes, &resultTemplate)
	if err != nil {
		return err
	}
	t, err := template.New("guess-age-template").Parse(constants.GuessAgeByNameTemplate)
	if err != nil {
		return err
	}
	err = t.Execute(os.Stdout, resultTemplate)
	if err != nil {
		return err
	}
	return nil
}

func LotrGetBooks() error {
	lotrClient := &client.LOTRClient{}
	resp, err := lotrClient.GetBooks()
	if err != nil {
		return err
	}
	t, err := template.New("lotr-book-list-template").Parse(constants.LOTRBookTemplate)
	if err != nil {
		return err
	}
	err = t.Execute(os.Stdout, resp)
	if err != nil {
		return err
	}
	return nil
}

func LotrGetBookById(id string) error {
	lotrClient := &client.LOTRClient{}
	resp, err := lotrClient.GetBookById(id)
	if err != nil {
		return err
	}
	t, err := template.New("lotr-book-detail-template").Parse(constants.LOTRBookTemplate)
	if err != nil {
		return err
	}
	err = t.Execute(os.Stdout, resp)
	if err != nil {
		return err
	}
	return nil
}

func LotrGetBookChaptersById(id string) error {
	lotrClient := &client.LOTRClient{}
	resp, err := lotrClient.GetBookChaptersById(id)
	if err != nil {
		return err
	}
	t, err := template.New("lotr-book-detail-template").Parse(constants.LOTRBookTemplate)
	if err != nil {
		return err
	}
	err = t.Execute(os.Stdout, resp)
	if err != nil {
		return err
	}
	return nil
}

func LotrGetMovies() error {
	lotrClient := &client.LOTRClient{}
	resp, err := lotrClient.GetMovies()
	if err != nil {
		return err
	}
	t, err := template.New("lotr-movie-list-template").Parse(constants.LOTRMovieTemplate)
	if err != nil {
		return err
	}
	err = t.Execute(os.Stdout, resp)
	if err != nil {
		return err
	}
	return nil
}

func LotrGetMovieById(id string) error {
	lotrClient := &client.LOTRClient{}
	resp, err := lotrClient.GetMovieById(id)
	if err != nil {
		return err
	}
	t, err := template.New("lotr-movie-detail-template").Parse(constants.LOTRMovieTemplate)
	if err != nil {
		return err
	}
	err = t.Execute(os.Stdout, resp)
	if err != nil {
		return err
	}
	return nil
}

func LotrGetMovieQuotesById(id string) error {
	lotrClient := &client.LOTRClient{}
	resp, err := lotrClient.GetMovieQuotesById(id)
	if err != nil {
		return err
	}
	t, err := template.New("lotr-movie-quote-template").Parse(constants.LOTRMovieQuoteTemplate)
	if err != nil {
		return err
	}
	err = t.Execute(os.Stdout, resp)
	if err != nil {
		return err
	}
	return nil
}
