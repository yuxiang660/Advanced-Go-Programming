package main

import "fmt"

func generateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func primeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()

	return out
}

func main() {
	fmt.Println("Prime Filter")

	ch := generateNatural()
	// Find out 100 Prime Number
	for i := 0; i < 100; i++ {
		prime := <-ch
		fmt.Printf("No.%v: %v\n", i+1, prime)
		ch = primeFilter(ch, prime)
	}
}
