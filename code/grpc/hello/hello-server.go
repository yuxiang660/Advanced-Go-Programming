package main

import (
	"context"
	"google.golang.org/grpc"
	"net"
	"log"
)

// HelloServiceImpl is the implemenation of gRPC HelloService.
type HelloServiceImpl struct{}

// Hello is the RPC method of HelloServiceImpl.
func (p *HelloServiceImpl) Hello(ctx context.Context, args *String) (*String, error) {
	reply := &String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func main() {
	grpcServer := grpc.NewServer()
	RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(listen)
}
