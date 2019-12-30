package main

import (
	"fmt"
	"context"
)

func generateNatural(task context.Context) chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			select {
			case <- task.Done():
				return
			case ch <- i:
			}
		}
	}()
	return ch
}

func primeFilter(task context.Context, in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			select {
			case <- task.Done():
				return
			case i := <-in:
				if i % prime != 0 {
					out <- i
				}
			}
		}
	}()

	return out
}

func main() {
	fmt.Println("Prime Filter")
	
	task, quit := context.WithCancel(context.Background())

	ch := generateNatural(task)
	// Find out 100 Prime Number
	for i := 0; i < 100; i++ {
		prime := <-ch
		fmt.Printf("No.%v: %v\n", i+1, prime)
		ch = primeFilter(task, ch, prime)
	}

	quit()
}
