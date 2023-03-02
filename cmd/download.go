/*
Copyright © 2022 En Zhou
This file is part of CLI application idev.
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

const (
	filename = "filename"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("download called")

		f, _ := cmd.Flags().GetString(filename)
		u, _ := cmd.Flags().GetString(URL)
		fmt.Println("Downloading ", u, " to ", f)

		resp, err := http.Get(u) //nolint:gosec
		if err != nil {
			log.Fatalf("Erro open site: %s", err)
		}
		defer resp.Body.Close()

		file, err := os.Create(f)
		if err != nil {
			log.Fatalf("Error creating file: %s", err)
		}
		defer file.Close()

		_, err = io.Copy(file, resp.Body)
	},
}

func init() {
	scraperCmd.AddCommand(downloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	downloadCmd.PersistentFlags().StringP(URL, "u", "", "Target website url that is going to be scraped")
	_ = downloadCmd.MarkPersistentFlagRequired(URL)
	downloadCmd.Flags().StringP(filename, "f", "", "File name to store downloaded HTML page")
	_ = downloadCmd.MarkFlagRequired(filename)
}
