package cmd

/*
Copyright © 2022 En Zhou
This file is part of CLI application idev.
*/

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	URL       = "url"
	output    = "output"
	schema    = "schema"
	selector  = "selector"
	attribute = "attribute"
)

// scraperCmd represents the scrape command
var scraperCmd = &cobra.Command{
	Use:   "scraper",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		u, _ := cmd.Flags().GetString(URL)
		s, _ := cmd.Flags().GetString(schema)
		o, _ := cmd.Flags().GetString(output)

		fmt.Println("scraper called with " + u + " " + s + " " + o)
	},
}

func init() {
	rootCmd.AddCommand(scraperCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	scraperCmd.PersistentFlags().StringP(filename, "f", "", "Target local file that is going to be scraped")
	_ = scraperCmd.MarkPersistentFlagRequired(filename)

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	scraperCmd.Flags().StringP(schema, "s", "", "Schema for scraped data structure in JSON format")
	// _ = scraperCmd.MarkFlagRequired(schema)
	scraperCmd.Flags().StringP(output, "o", "", "Output data file name with extension")
	// _ = scraperCmd.MarkFlagRequired(output)
}
