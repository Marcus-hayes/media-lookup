package cmd

import (
	"github.com/Marcus-hayes/media-lookup/handler"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().String("query", "", "")
	searchCmd.MarkFlagRequired("query")
	searchCmd.Flags().Bool("nsfw", false, "")
	searchCmd.Flags().String("language", "en-US", "")
	searchCmd.Flags().Int32("page", 1, "")
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Perform search using input query string",
	Long:  "Search for a movie, TV show, or person using the TMDB API ",
	RunE: func(cmd *cobra.Command, args []string) error {
		query, err := cmd.Flags().GetString("query")
		if err != nil {
			return err
		}
		nsfw, err := cmd.Flags().GetBool("nsfw")
		if err != nil {
			return err
		}
		lang, err := cmd.Flags().GetString("language")
		if err != nil {
			return err
		}
		page, err := cmd.Flags().GetInt32("page")
		if err != nil {
			return err
		}
		handler.Init()
		err = handler.PerformSearch(query, nsfw, lang, page)
		if err != nil {
			return err
		}
		return nil
	},
}
