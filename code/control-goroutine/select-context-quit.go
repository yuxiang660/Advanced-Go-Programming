package main

import (
	"fmt"
	"context"
	"sync"
	"time"
)

func worker(waitGroup *sync.WaitGroup, c context.Context) error {
	defer waitGroup.Done()

	for {
		select {
		default:
			fmt.Println("hello")
			time.Sleep(100 * time.Millisecond)
		case <-c.Done():
			return c.Err()
		}
	}
}

func main() {
	fmt.Println("Quit with Context")

	c, quit := context.WithTimeout(context.Background(), 10 * time.Second)

	var waitGroup sync.WaitGroup

	for i := 0; i < 3; i++ {
		waitGroup.Add(1)
		go worker(&waitGroup, c)
	}

	time.Sleep(time.Second)
	quit()

	waitGroup.Wait()
}
