package main

import (
	"net"
	"log"
	"net/rpc"
)

// HelloService ...
type HelloService struct {}

// Hello ...
func (p *HelloService) Hello(request *String, reply *String) error {
	reply.Value = "hello:" + request.GetValue()
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	rpc.ServeConn(conn)
}
