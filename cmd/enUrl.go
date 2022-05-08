/*
Copyright © 2022 En Zhou <zhouen.nathan@gmail.com>

*/

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/url"
)

// enUrlCmd represents the enUrl command
var enUrlCmd = &cobra.Command{
	Use:   "enUrl",
	Short: "Ecode a URL string",
	//	Long: `A longer description that spans multiple lines and likely contains examples
	//and usage of using your command. For example:
	//
	//Cobra is a CLI library for Go that empowers applications.
	//This application is a tool to generate the needed files
	//to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		urlStr := ""
		if len(args) >= 1 && args[0] != "" {
			urlStr = args[0]
		}

		fmt.Println(url.PathEscape(urlStr))

	},
}

func init() {
	rootCmd.AddCommand(enUrlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// enUrlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// enUrlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
