/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/

package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	pb "github.com/zhou-en/idev/pkg/gopher"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
)

const (
	port         = ":9000"
	KuteGoAPIURL = "http://localhost:8080"
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
		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()

		// Register services
		pb.RegisterGopherServer(grpcServer, &Server{})
		log.Printf("GRPC server listening on %v", lis.Addr())

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	},
}

// GetGopher implements gopher.GopherServer
func (s *Server) GetGopher(ctx context.Context, req *pb.GopherRequest) (*pb.GopherReply, error) {
	res := &pb.GopherReply{}

	// Check request
	if req == nil {
		fmt.Println("Request must not be nil")
		return res, xerrors.Errorf("request must not return nil")
	}

	if req.Name == "" {
		msg := "name must not be empty n the request"
		fmt.Println(msg)
		return res, xerrors.Errorf(msg)
	}

	log.Printf("Received: %v", req.GetName())

	// Call KuteGo API in order to get Gopher's URL
	response, err := http.Get(KuteGoAPIURL + "/gophers?name=" + req.GetName())
	if err != nil {
		log.Fatalf("failed to call KuteGoAPI: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		// Transform our response to a []byte
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatalf("failed to read response body: %v", err)
		}

		// Put only needed informations of the JSON document in our array of Gopher
		var data []Gopher
		err = json.Unmarshal(body, &data)
		if err != nil {
			log.Fatalf("failed to unmarshal JSON: %v", err)
		}

		// Create a string with all of the Gopher's name and a blank line as septerator
		var gophers strings.Builder
		for _, gopher := range data {
			gophers.WriteString(gopher.URL + "\n")
		}

		res.Message = gophers.String()
	} else {
		log.Fatal("Can't get the Gopher :-(")
	}
	return res, nil
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
