package cmd

import (
	"github.com/Marcus-hayes/media-lookup/handler"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(guessAgeCmd)
	guessAgeCmd.Flags().String("name", "", "Name used to guess age")
	guessAgeCmd.MarkFlagRequired("name")
}

var guessAgeCmd = &cobra.Command{
	Use:   "guess-age-by-name",
	Short: "Guess your age based on your name",
	Long:  "Guess your age based on your name via Agify.io",
	RunE: func(cmd *cobra.Command, args []string) error {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		err = handler.GuessAgeByName(name)
		if err != nil {
			return err
		}
		return nil
	},
}
