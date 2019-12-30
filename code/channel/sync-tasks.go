package main

import "fmt"

func foo(n int, done chan<- int) {
	for i := 0; i < n; i++ {
		fmt.Println("hello")
		done <- i
	}
	close(done)
}

func main() {
	fmt.Println("Sync Tasks with Channel")

	done := make(chan int, 10)

	go foo(cap(done), done)

	for i := range done {
		fmt.Println(i)
	}
}