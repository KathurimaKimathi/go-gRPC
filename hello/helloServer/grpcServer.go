package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/KathurimaKimathi/go-gRPC/hello/hellopb"
	"google.golang.org/grpc"
)

type Server struct{}

func (s *Server) Hello(ctx context.Context, request *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	name := request.Name

	response := &hellopb.HelloResponse{
		Name: "Hello " + name,
	}

	return response, nil
}

func main() {
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()
	hellopb.RegisterHelloServiceServer(s, &Server{})

	s.Serve(lis)

}
