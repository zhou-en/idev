/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/

package cmd

import (
	"github.com/spf13/cobra"
	pb "github.com/zhou-en/idev/pkg/gopher"
)

const (
	port         = ":9000"
	KuteGoAPIURL = "https://kutego-api-xxxxx-ew.a.run.app"
)

// Server is used to implement gopher.GopherServer
type Server struct {
	pb.UnimplementedGopherServer
}

type Gopher struct {
	URL string `json: "url"`
}

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the schema gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		//lis, err := net.Listen("tcp", port)
		//if err != nil {
		//	log.Fatalf("failed to listen: %v", err)
		//}
		//
		//grpcServer := grpc.NewServer()
		//
		//// Register services
		//pb.Regis
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
