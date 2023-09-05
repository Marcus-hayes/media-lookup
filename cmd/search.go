package cmd

import (
	"fmt"
	"log"

	"github.com/Marcus-hayes/media-lookup/handler"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().String("query", "", "")
	searchCmd.MarkFlagRequired("query")
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Perform search using input query string",
	Long:  "Search for a movie, TV show, or person using the TMDB API ",
	RunE: func(cmd *cobra.Command, args []string) error {
		query, err := cmd.Flags().GetString("query")
		if err != nil {
			log.Println("Error retrieving query: ", err)
		}
		nsfw, err := cmd.Flags().GetString("nsfw")
		if err != nil {
			log.Println("Error retrieving nsfw flag: ", err)
		}
		query, err := cmd.Flags().GetString("language")
		if err != nil {
			log.Println("Error retrieving query: ", err)
		}
		result, err := handler.PerformSearch(query)
		fmt.Println(result)
		return nil
	},
}
