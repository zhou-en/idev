/*
Copyright © 2022 En Zhou
This file is part of CLI application idev.
*/
package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
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
		//schemaPath, _ := cmd.Flags().GetString(schema)
		u, _ := cmd.Flags().GetString(URL)
		resp, err := http.Get(u) //nolint:gosec
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			log.Panicf("status code error: %d %s", resp.StatusCode, resp.Status)
		}

		// scrape a site based on the given rule
		//if schemaPath != "" {
		//	outputPath, _ := cmd.Flags().GetString(output)
		//	querySite(schemaPath, outputPath, resp)
		//	fmt.Println("data was saved in " + outputPath)
		//}

		// scrape a element
		if s != "" {
			a, _ := cmd.Flags().GetString(attribute)
			doc, err := goquery.NewDocumentFromReader(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			result := doc.Find(s)
			if a != "" {
				val, exists := result.Attr(a)
				if exists {
					fmt.Println(val)
				}
				fmt.Println("attribute not found")
			} else {
				fmt.Println(result.Text())
			}
		}
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

func _(urlAddr string) (*http.Response, error) {
	res, err := http.Get(urlAddr) //nolint:gosec
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Panicf("status code error: %d %s", res.StatusCode, res.Status)
	}
	return res, nil
}

//func _(filePath string) string {
//	return ""
//}
//
//func writeFile(filePath string) {
//
//}
//
//func querySite(schemaPath, outputPath string, resp *http.Response) {
//
//}

func _(s string, resp *http.Response) {
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	result := doc.Find(s)
	fmt.Println(result.Text())
}
