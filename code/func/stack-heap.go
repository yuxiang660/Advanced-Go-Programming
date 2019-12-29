package main

import "fmt"

func foo(x int) *int {
	return &x
}

func bar() *int {
	x := 6
	return &x;
}

func bar2() *int {
	var x = new(int)
	return x;
}

func main() {
	fmt.Println("Stack and Heap in Go")

	x := 5
	fmt.Printf("*(%#v) = %d\n", foo(x), *(foo(x)))
	fmt.Printf("*(%#v) = %d\n", bar(), *(bar()))
	fmt.Printf("*(%#v) = %d\n", bar2(), *(bar2()))
}
