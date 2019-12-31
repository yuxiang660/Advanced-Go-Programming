package main

import (
	"./pubsub"
	"time"
	"context"
	"strings"
	"google.golang.org/grpc"
	"log"
	"net"
)

// PubsubService is the RPC service for remote calls of pubsub methods.
type PubsubService struct {
	pub *pubsub.Publisher
}

// NewPubsubService is used to create the pubsub service.
func NewPubsubService() *PubsubService {
	return &PubsubService{
		pub: pubsub.NewPublisher(100 * time.Millisecond, 10),
	}
}

// Publish is an RPC method of pubsub service, to publish msg to subscribers.
func (p *PubsubService) Publish(ctx context.Context, arg *String) (*String, error) {
	p.pub.Publish(arg.GetValue())
	return &String{}, nil
}

// Subscribe is an RPC method of pubsub service, to subscibe a subscirber to the publisher on the service.
// And send a "chan" through gRPC stream channel. Through the "chan", client can get the msg from publisher.
func (p *PubsubService) Subscribe(arg *String, stream PubsubService_SubscribeServer) error {
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, arg.GetValue()) {
				return true
			}
		}
		return false
	})

	for v := range ch {
		if err := stream.Send(&String{Value: v.(string)}); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	grpcServer := grpc.NewServer()
	RegisterPubsubServiceServer(grpcServer, NewPubsubService())

	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(listen)
}
