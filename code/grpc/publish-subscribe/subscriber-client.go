package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"fmt"
	"io"
	"os"
)

func usage() {
	fmt.Println("Usage for Subscriber Client:")
	fmt.Println("1. go run subsrciber-client.go pubsub.pb.go golang")
	fmt.Println("or")
	fmt.Println("2. go run subsrciber-client.go pubsub.pb.go docker")
}

func main() {
	args := os.Args
	if len(args) < 2 || args == nil {
		usage()
		return
	}

	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := NewPubsubServiceClient(conn)
	stream, err := client.Subscribe(context.Background(), &String{Value: args[1] + ":"})
	if err != nil {
		log.Fatal(err)
	}

	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(reply.GetValue())
	}
}