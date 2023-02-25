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
	Short: "Encode a URL string",
	Long:  `Example usage: idev enUrl "https://github.com/zhou-en"`,
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
