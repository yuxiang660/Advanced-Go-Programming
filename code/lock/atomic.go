package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var total uint64

func worker(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	for i := uint64(0); i <= 100; i++ {
		atomic.AddUint64(&total, i)
	}
}

func main() {
	fmt.Println("Lock Operation with Atomic")

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go worker(&waitGroup)
	go worker(&waitGroup)
	waitGroup.Wait()

	fmt.Println(total)
}