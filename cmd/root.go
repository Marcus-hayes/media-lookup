package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "media-lookup",
	Short: "media-lookup is a CLI tool for finding media metadata.",
	Long:  "Retrieves media information, such as creation date, creator, description, and other various information given an input media",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
	}
	os.Exit(1)
}
