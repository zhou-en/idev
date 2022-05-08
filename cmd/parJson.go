/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// parJsonCmd represents the parJson command
var parJsonCmd = &cobra.Command{
	Use:   "parJson",
	Short: "Parse stringfied JSON to pretty format and output to console.",
	//	Long: `A longer description that spans multiple lines and likely contains examples
	//and usage of using your command. For example:
	//
	//Cobra is a CLI library for Go that empowers applications.
	//This application is a tool to generate the needed files
	//to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		jsonStr := ""
		if len(args) >= 1 && args[0] != "" {
			jsonStr = args[0]
		}
		cmdStr := fmt.Sprintf("echo %s | jq", jsonStr)
		//fmt.Println(cmdStr)
		c := exec.Command("/bin/bash", "-c", cmdStr)
		//c := exec.Command("/bin/sh", "-c", "echo $GOPATH")
		stdout, _ := c.Output()

		fmt.Println(string(stdout))
	},
}

func init() {
	rootCmd.AddCommand(parJsonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	parJsonCmd.PersistentFlags().String("file", "", "Parse a stringfied JSON from a text file.")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// parJsonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
