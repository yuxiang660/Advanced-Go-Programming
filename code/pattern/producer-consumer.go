package main

import (
	"fmt"
	"time"
	"os"
	"syscall"
	"os/signal"
)

func producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		time.Sleep(500 * time.Millisecond)
		out <- i * factor
	}
}

func consumer(in <-chan int) {
	for v := range in {
		time.Sleep(250 * time.Millisecond)
		fmt.Println(v)
	}
}

func main() {
	fmt.Println("Produce and Consumer Pattern")

	ch := make(chan int, 64)

	go producer(3, ch)
	go producer(5, ch)
	go consumer(ch)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)
	fmt.Printf("quit (%v)\n", <-sig)
}