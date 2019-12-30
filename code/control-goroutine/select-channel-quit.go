package main

import (
	"fmt"
	"time"
	"sync"
)

func worker(waitGroup *sync.WaitGroup, quit <-chan bool) {
	defer waitGroup.Done()

	for {
		select {
		default:
			fmt.Println("hello")
			time.Sleep(100 * time.Millisecond)
		case <-quit:
			return
		}
	}
}

func main() {
	fmt.Println("Select to Quit a Goroutine")

	quit := make(chan bool)
	var waitGroup sync.WaitGroup
	
	for i := 0; i < 3; i++ {
		waitGroup.Add(1)
		go worker(&waitGroup, quit)
	}

	time.Sleep(time.Second)
	close(quit)

	waitGroup.Wait()
}
