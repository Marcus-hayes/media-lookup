package handler

import (
	"fmt"
	"strings"

	"github.com/Marcus-hayes/media-lookup/client"
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

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
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("ID", "Name/Title", "Original Name/Title", "Release Date", "Original Language", "Country of Origin", "")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for _, result := range results {
		title := strings.TrimSpace(fmt.Sprintf("%s %s", result.Title, result.Name))
		origTitle := strings.TrimSpace(fmt.Sprintf("%s %s", result.OriginalName, result.OriginalTitle))

		tbl.AddRow(result.ID, title, origTitle, result.ReleaseDate, result.OriginalLanguage, result.OriginCountry)
	}
	tbl.Print()
	return nil
}
