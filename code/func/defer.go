package main

import "fmt"

func main() {
	fmt.Println("Defer Function")

	for i := 0; i < 3; i++ {
		defer func() {fmt.Println("closure:", i)} ()
	}

	for i := 0; i < 3; i++ {
		defer func(i int) {fmt.Println("pass-by-value", i)} (i)
	}
}
