// Package cmd /*
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// parJsonCmd represents the parJson command
var parJSONCmd = &cobra.Command{
	Use:   "parJson",
	Short: "Parse stringfied JSON to pretty format and output to console.",
	Long:  `Example: idev --parJson '{\"a\":\"aa\", \"b\": {\"c\":\"cc\"}, \"d\": [1,2,3]}'`,
	Run: func(cmd *cobra.Command, args []string) {
		jsonStr := ""
		if len(args) >= 1 && args[0] != "" {
			jsonStr = args[0]
		}
		cmdStr := fmt.Sprintf("echo %s | jq", jsonStr)
		//fmt.Println(cmdStr)
		c := exec.Command("/bin/bash", "-c", cmdStr)
		stdout, _ := c.Output()

		fmt.Println(string(stdout))
	},
}

func init() {
	rootCmd.AddCommand(parJSONCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	parJSONCmd.PersistentFlags().String("file", "", "Parse a string field JSON from a text file.")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// parJsonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
