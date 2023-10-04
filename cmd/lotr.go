package cmd

import (
	"github.com/Marcus-hayes/media-lookup/handler"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(lotrCmd)
	lotrCmd.AddCommand(lotrBookSubCmd)
	// Book - List Sub CMD
	lotrBookSubCmd.AddCommand(lotrBookListSubCmd)
	// Book - Detail Sub CMD
	lotrBookSubCmd.AddCommand(lotrBookDetailSubCmd)
	lotrBookDetailSubCmd.Flags().String("id", "", "id string, used to search for LoTR book")
	lotrBookDetailSubCmd.MarkFlagRequired("id")

	// Book - Chapters Sub CMD
	lotrBookSubCmd.AddCommand(lotrBookChaptersSubCmd)
	lotrBookChaptersSubCmd.Flags().String("id", "", "id string, used to search for LoTR book")
	lotrBookChaptersSubCmd.MarkFlagRequired("id")
}

var lotrCmd = &cobra.Command{
	Use:   "lotr",
	Short: "Access LoTR facts",
	Long:  "Get LoTR data via the The One API, defined here: https://the-one-api.dev/documentation",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var lotrBookSubCmd = &cobra.Command{
	Use:   "book",
	Short: "Access LoTR book resources",
	Long:  "Get LoTR data via the The One API, defined here: https://the-one-api.dev/documentation",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var lotrBookListSubCmd = &cobra.Command{
	Use:   "list",
	Short: "List LoTR books",
	Long:  "Get LoTR data via the The One API, defined here: https://the-one-api.dev/documentation",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := handler.LotrGetBooks()
		if err != nil {
			return err
		}
		return nil
	},
}

var lotrBookDetailSubCmd = &cobra.Command{
	Use:   "detail",
	Short: "Get details of a LoTR book by searching its ID",
	Long:  "Get LoTR data via the The One API, defined here: https://the-one-api.dev/documentation",
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}
		err = handler.LotrGetBookById(id)
		if err != nil {
			return err
		}
		return nil
	},
}

var lotrBookChaptersSubCmd = &cobra.Command{
	Use:   "chapters",
	Short: "Get chapters of a LoTR book by searching its ID",
	Long:  "Get LoTR data via the The One API, defined here: https://the-one-api.dev/documentation",
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}
		err = handler.LotrGetBookChaptersById(id)
		if err != nil {
			return err
		}
		return nil
	},
}
