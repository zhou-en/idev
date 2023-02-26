/*
Copyright © 2022 En Zhou <zhoeun.nathan@gmail.com>

*/

package cmd

import (
	"context"
	"fmt"
	"github.com/google/go-github/v50/github"
	"github.com/spf13/cobra"
	"net/url"
	"os"
	"os/exec"
)

const enUrl = "enUrl"
const parJson = "parJson"
const parJsonf = "parJsonf"

func encodeUrl(urlStr string) {
	fmt.Println(url.PathEscape(urlStr))
}

func parseJson(jsonStr string) {
	cmdStr := fmt.Sprintf("echo %s | jq", jsonStr)
	c := exec.Command("/bin/bash", "-c", cmdStr)
	stdout, _ := c.Output()
	fmt.Println(string(stdout))
}

func parseJsonFile(filePath string) {
	cmdStr := fmt.Sprintf("jq 'fromjson | .' %s", filePath)
	c := exec.Command("/bin/bash", "-c", cmdStr)
	stdout, _ := c.Output()
	fmt.Println(string(stdout))
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "idev",
	Short: "Collection of CLI tools for local development purposes.",
	Long: `idev is a CLI tool written in Go that provides commonly used command during local development, 
such as parse stringfied JSON to console, read a stringfied text file and parse the output to console, and encode URL strings.
This purpose of this application is to increase the local development productivity by the simple interface.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		enUrlValue, _ := cmd.Flags().GetString(enUrl)
		parJsonValue, _ := cmd.Flags().GetString(parJson)
		parJsonfValue, _ := cmd.Flags().GetString(parJsonf)

		switch {
		case enUrlValue != "":
			encodeUrl(enUrlValue)
		case parJsonValue != "":
			parseJson(parJsonValue)
		case parJsonfValue != "":
			parseJsonFile(parJsonfValue)
		default:
			fmt.Println("No Value was Supplied.")
		}

	},
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
	//var Verbose bool

	//var Source string
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	//rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringP(enUrl, "", "", "Encode a URL")
	rootCmd.Flags().StringP(parJson, "", "", "Parse stringfied JSON")
	rootCmd.Flags().StringP(parJsonf, "", "", "Parse stringfied JSON from a given file")

	//viper.BindPFlag("test", rootCmd.PersistentFlags().Lookup("test"))
	//rootCmd.ParseFlags()
	//testStr := viper.GetString("test")
	//fmt.Println(testStr)
	rootCmd.AddCommand(versionCmd)
}
