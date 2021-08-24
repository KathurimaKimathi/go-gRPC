package main

import (
	"context"
	"fmt"
	"log"

	"github.com/KathurimaKimathi/go-gRPC/hello/hellopb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello client ...")

	opt := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opt)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := hellopb.NewHelloServiceClient(cc)
	req := &hellopb.HelloRequest{
		Name: "Kathurima Kimathi",
	}

	resp, _ := client.Hello(context.Background(), req)
	fmt.Printf("Receive response --> [%v]\n", resp.Name)
}
