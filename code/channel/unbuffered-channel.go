package main

import "fmt"

var done = make(chan bool)
var msg string

func foo() {
	msg = "hello"
	// blocked by "<- done"
	done <- true
}

func main() {
	fmt.Println("Unbuffered Channel")

	go foo()
	// blocked by "done <- true"
	<- done
	println(msg)
}
