package main

import (
	"fmt"
	"sync"
)

var total struct {
	sync.Mutex
	value int
}

func worker(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	for i := 0; i <= 100; i++ {
		total.Lock()
		total.value += i
		total.Unlock()
	}
}

func main() {
	fmt.Println("Lock Operation with Mutex")

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go worker(&waitGroup)
	go worker(&waitGroup)
	waitGroup.Wait()

	fmt.Println(total.value)
}
