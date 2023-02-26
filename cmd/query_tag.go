/*
Copyright © 2022 En Zhou
This file is part of CLI application idev.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	selector  = "selector"
	attribute = "attribute"
)

// queryTagCmd represents the queryTag command
var queryTagCmd = &cobra.Command{
	Use:   "queryTag",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		s, _ := cmd.Flags().GetString(selector)
		a, _ := cmd.Flags().GetString(attribute)
		fmt.Println(s)
		fmt.Println(a)
	},
}

func init() {
	scraperCmd.AddCommand(queryTagCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// queryTagCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	queryTagCmd.Flags().StringP(selector, "s", "", "CSS query selector")
	_ = queryTagCmd.MarkFlagRequired(selector)
	queryTagCmd.Flags().StringP(attribute, "a", "", "Attribute name of the element")
}
