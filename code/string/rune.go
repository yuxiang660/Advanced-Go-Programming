package main

import "fmt"

func main() {
	fmt.Println("rune type")

	fmt.Printf("%#v\n", []rune("世界"))
	fmt.Printf("%#v\n", []rune{'世', '界'})
	fmt.Printf("%#v\n", string([]rune("世界")))
	fmt.Printf("%#v\n", string([]rune{'世', '界'}))
}