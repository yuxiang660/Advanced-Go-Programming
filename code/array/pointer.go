package main

import "fmt"

func main() {
	fmt.Println("Pointer to An Array")

	var arr = [...]int{1, 2}
	var arrCopy = arr
	var pArr = &arr
	var pArrCopy = pArr

	fmt.Println("arr[0] =", arr[0], "arr[1] =", arr[1])
	fmt.Println("arrCopy[0] =", arrCopy[0], "arrCopy[1] =", arrCopy[1])
	fmt.Println("pArr[0] =", pArr[0], "pArr[1] =", pArr[1])
	fmt.Println("pArrCopy[0] =", pArrCopy[0], "pArrCopy[1] =", pArrCopy[1])

	pArr[0] = 5
	fmt.Println("\nAfter pArr[0] = 5:")

	for i, v := range arr {
		fmt.Printf("arr[%d]: %d\n", i, v)
	}
	for i, v := range arrCopy {
		fmt.Printf("arrCopy[%d]: %d\n", i, v)
	}
	for i, v := range pArr {
		fmt.Printf("pArr[%d]: %d\n", i, v)
	}
	for i, v := range pArrCopy {
		fmt.Printf("pArrCopy[%d]: %d\n", i, v)
	}
}