package main

import (
	"net/rpc"
	"log"
	"fmt"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply String
	request := String{Value: "user"}
	err = client.Call("HelloService.Hello", &request, &reply)
	if err!= nil {
		log.Fatal(err)
	}

	fmt.Println(reply.String())
}
