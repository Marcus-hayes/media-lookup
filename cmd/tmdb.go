package cmd

import (
	"github.com/Marcus-hayes/media-lookup/handler"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(searchCmd)
	// Search CMD Flags
	searchCmd.Flags().String("query", "", "")
	searchCmd.MarkFlagRequired("query")
	searchCmd.Flags().Bool("nsfw", false, "")
	searchCmd.Flags().String("language", "en-US", "")
	searchCmd.Flags().Int32("page", 1, "")

	rootCmd.AddCommand(tmdbDetailCmd)
	// Detail CMD Flags
	tmdbDetailCmd.Flags().String("id", "", "")
	tmdbDetailCmd.MarkFlagRequired("id")
	tmdbDetailCmd.Flags().String("media-type", "", "")
	tmdbDetailCmd.MarkFlagRequired("media-type")
}

// Detail CMD
var tmdbDetailCmd = &cobra.Command{
	Use:   "details",
	Short: "Get metadata on TMDB object",
	Long:  "Get metadata for a movie, TV show, or person using the TMDB API ",
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}
		media, err := cmd.Flags().GetString("media-type")
		if err != nil {
			return err
		}
		handler.InitTMDB()
		err = handler.GetTMDBDetails(id, media)
		if err != nil {
			return err
		}
		return nil
	},
}

// Search CMD
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
		handler.InitTMDB()
		err = handler.PerformSearch(query, nsfw, lang, page)
		if err != nil {
			return err
		}
		return nil
	},
}
