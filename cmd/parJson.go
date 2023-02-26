// Package cmd /*
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

const (
	path = "path"
)

// parJsonCmd represents the parJson command
var parJSONCmd = &cobra.Command{
	Use:   "parJson",
	Short: "Parse stringfied JSON to pretty format and output to console.",
	Long:  `Example: idev --parJson '{\"a\":\"aa\", \"b\": {\"c\":\"cc\"}, \"d\": [1,2,3]}'`,
	Run: func(cmd *cobra.Command, args []string) {
		p, _ := cmd.Flags().GetString(path)
		j, _ := cmd.Flags().GetString(parJSON)
		if p == "" && j != "" {
			cmdStr := fmt.Sprintf("echo %s | jq", j)
			c := exec.Command("/bin/bash", "-c", cmdStr)
			stdout, _ := c.Output()
			fmt.Println(string(stdout))
		}
		cmdStr := fmt.Sprintf("jq . %s", p)
		fmt.Println(cmdStr)
		c := exec.Command("/bin/bash", "-c", cmdStr)
		stdout, _ := c.Output()
		fmt.Println(string(stdout))
	},
}

func init() {
	rootCmd.AddCommand(parJSONCmd)
	parJSONCmd.Flags().StringP(path, "p", "", "Pretty print JSON file.")
}
