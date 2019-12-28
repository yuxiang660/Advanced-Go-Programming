package main

import "fmt"

func main() {
	fmt.Println("Single Quote, Double Quote, Back Quote")

	s := string([]rune{'h', 'e', 'l', 'l', 'o', '\t', 'w', 'o', 'r', 'l', 'd'})
	s2 := "hello\tworld"
	s3 := `hello\tworld`

	fmt.Println(s)
	fmt.Println(s2)
	fmt.Println(s3)
}
