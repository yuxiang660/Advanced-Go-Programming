package main

import "fmt"

type point struct {
	X int
	Y int
}

func foo(firstName string) int{
	fmt.Println("First Name:", firstName)
	return len(firstName)
}

func bar(lastName string) int{
	fmt.Println("Last Name:", lastName)
	return len(lastName)
}

func main() {
	fmt.Println("Array Types")

	var stringArr = [2]string{"hello", "world"}
	var stringArr2 = [...]string{"hello", "world"}
	var stringArr3 = [...]string{1:"world", 0:"world"}

	fmt.Printf("stringArr: %T, %#v\n", stringArr, stringArr)
	fmt.Printf("stringArr2: %T, %#v\n", stringArr2, stringArr2)
	fmt.Printf("stringArr3: %T, %#v\n", stringArr3, stringArr3)

	var structArr [2]point	
	var structArr2 = [...]point{point{X: 0, Y: 0}, point{X: 1, Y: 1}}
	var structArr3 = [...]point{{0, 0}, {1, 1}}

	fmt.Printf("structArr: %T, %#v\n", structArr, structArr)
	fmt.Printf("structArr2: %T, %#v\n", structArr2, structArr2)
	fmt.Printf("structArr3: %T, %#v\n", structArr3, structArr3)

	var funcArr [2]func(string) (int)
	var funcArr2 = [...]func(string) (int) {
		foo,
		bar,
	}

	fmt.Printf("funcArr: %T, %#v\n", funcArr, funcArr)
	fmt.Printf("funcArr2: %T, %#v\n", funcArr2, funcArr2)

	var interfaceArr [2]interface{}
	var interfaceArr2 = [...]interface{}{123, "hello"}

	fmt.Printf("interfaceArr: %T, %#v\n", interfaceArr, interfaceArr)
	fmt.Printf("interfaceArr2: %T, %#v\n", interfaceArr2, interfaceArr2)

	var chanArr = [2]chan int{}

	fmt.Printf("chanArr: %T, %#v\n", chanArr, chanArr)
}