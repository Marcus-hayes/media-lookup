package cmd

import (
	"github.com/Marcus-hayes/media-lookup/handler"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(detailCmd)
	detailCmd.Flags().String("id", "", "")
	detailCmd.MarkFlagRequired("id")
	detailCmd.Flags().String("media-type", "", "")
	detailCmd.MarkFlagRequired("media-type")
}

var detailCmd = &cobra.Command{
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
		handler.Init()
		err = handler.GetTMDBDetails(id, media)
		if err != nil {
			return err
		}
		return nil
	},
}
