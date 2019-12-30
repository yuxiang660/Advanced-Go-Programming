package main

import (
	"net"
	"log"
	"net/rpc"
	"./hello"
)

// HelloService is the concrete hello service.
type HelloService struct {}

// Hello is the RPC method of the hello service.
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main() {
	hello.RegisterService(new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
	
		go rpc.ServeConn(conn)
	}
}