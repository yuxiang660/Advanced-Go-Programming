package main

import "fmt"

func main() {
	fmt.Println("Two Dimension Array")

	var arr [3][2]int
	var arr2 = [3][2]int{{1,2},{3,4},{5,6}}
	var arr3 = [3][2]int{{},{3,4},{1:6}}

	for row, v := range arr {
		for column, value := range v {
			fmt.Printf("Arr index[%d][%d]: %d\n", row, column, value)
		}
	}

	for row, v := range arr2 {
		for column, value := range v {
			fmt.Printf("Arr2 index[%d][%d]: %d\n", row, column, value)
		}
	}

	for row, v := range arr3 {
		for column, value := range v {
			fmt.Printf("Arr3 index[%d][%d]: %d\n", row, column, value)
		}
	}
}