package main

import "fmt"

func main() {
	fmt.Println("Array Defination")

	var arrDef1 [3]int
	fmt.Println("Defination: \"[3]int, value\":", arrDef1)

	var arrDef2 = [...]int{1, 2, 3}
	fmt.Println("Defination: \"[...]int{1, 2, 3}, value:\"", arrDef2)

	var arrDef3 = [...]int{2: 3, 1: 2}
	fmt.Println("Defination: \"[...]int{2: 3, 1: 2}, value:\"", arrDef3)

	var arrDef4 = [...]int{1, 2, 4: 5, 6}
	fmt.Println("Defination: \"[...]int{1, 2, 4: 5, 6}, value:\"", arrDef4)
}

