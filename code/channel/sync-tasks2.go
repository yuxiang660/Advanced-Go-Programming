package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Sync Tasks with WaitGroup")

	var waitGroup sync.WaitGroup

	for i := 0; i < 10; i++ {
		waitGroup.Add(1)

		go func() {
			defer waitGroup.Done()
			fmt.Println("hello")
		}()
	}

	waitGroup.Wait()
}