/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"github.com/spf13/cobra"
	pb "github.com/zhou-en/idev/pkg/gopher"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

const (
	address     = "localhost:9000"
	defaultName = "dr-who"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Query the gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		var conn *grpc.ClientConn
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Did not connect: %s", err)
		}
		defer conn.Close()

		client := pb.NewGopherClient(conn)

		var name string

		// Contact the server and print out its response
		// name := defaultName
		if len(os.Args) > 2 {
			name = os.Args[2]
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := client.GetGopher(ctx, &pb.GopherRequest{Name: name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("URL: %s", r.GetMessage())
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
