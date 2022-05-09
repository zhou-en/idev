/*
Copyright © 2022 En Zhou <zhoeun.nathan@gmail.com>

*/

package cmd

import (
	"fmt"
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
	Use:   "go-gopher-grpc",
	Short: "gRPC app in Go",
	Long:  `gRPC application written in Go.`,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of idev",
	Long:  `All software has versions. This is idev's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
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
