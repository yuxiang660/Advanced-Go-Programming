package main

import (
	"net/rpc"
	"log"
	"fmt"
	"net/rpc/jsonrpc"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	err = client.Call("HelloService.Hello", "user", &reply)
	if err!= nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}