package main

import (
	"log"
	"fmt"
	"./hello"
)

func main() {
	client, err := hello.DialService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Hello("userABC", &reply)
	if err!= nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}