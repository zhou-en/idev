/*
Copyright © 2022 En Zhou
This file is part of CLI application idev.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// photoCmd represents the photo command
var photoCmd = &cobra.Command{
	Use:   "photo",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("photo called")
	},
}

func init() {
	rootCmd.AddCommand(photoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// photoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// photoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
