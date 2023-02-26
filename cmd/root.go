/*
Copyright © 2022 En Zhou <zhoeun.nathan@gmail.com>

*/

package cmd

import (
	"context"
	"fmt"
	"github.com/google/go-github/v50/github"
	"github.com/spf13/cobra"
	"os"
)

const enURL = "enUrl"
const parJSON = "parJson"
const parJsonF = "parJsonF"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "idev",
	Short: "Collection of CLI tools for local development purposes.",
	Long: `idev is a CLI tool written in Go that provides commonly used command during local development, 
such as parse stringfied JSON to console, read a stringfied text file and parse the output to console, and encode URL strings.
This purpose of this application is to increase the local development productivity by the simple interface.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of idev",
	Long:  `All software has versions. This is idev's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(getVersion())
	},
}

func getVersion() string {
	client := github.NewClient(nil)
	tags, _, err := client.Repositories.ListTags(context.Background(), "zhou-en", "idev", nil)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	if len(tags) > 0 {
		return tags[0].GetName()
	}
	return "No tags yet"
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(0)
	}
}

func init() {
	rootCmd.Flags().StringP(enURL, "", "", "Encode a URL")
	rootCmd.Flags().StringP(parJSON, "", "", "Parse stringfied JSON")
	rootCmd.Flags().StringP(parJsonF, "", "", "Parse stringfied JSON from a given file")
	rootCmd.AddCommand(versionCmd)
}
