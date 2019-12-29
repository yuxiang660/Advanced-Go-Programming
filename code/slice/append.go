package main

import "fmt"

func main() {
	fmt.Println("Slice Append")

	var slice []int
	fmt.Printf("%#v, len = %d, cap = %d\n", slice, len(slice), cap(slice))

	slice = []int{}
	fmt.Printf("%#v, len = %d, cap = %d\n", slice, len(slice), cap(slice))

	slice = append(slice, 1)
	fmt.Printf("%#v, len = %d, cap = %d\n", slice, len(slice), cap(slice))

	slice = slice[:0] //empty slice
	fmt.Printf("%#v, len = %d, cap = %d\n", slice, len(slice), cap(slice))

	slice = append(slice, 1, 2, 3, 4)
	fmt.Printf("%#v, len = %d, cap = %d\n", slice, len(slice), cap(slice))

	slice = append(slice, []int{5, 6, 7}...)
	fmt.Printf("%#v, len = %d, cap = %d\n", slice, len(slice), cap(slice))

	slice = append([]int{0}, slice...)
	fmt.Printf("%#v, len = %d, cap = %d\n", slice, len(slice), cap(slice))

	slice = append([]int{-2, -1}, slice...)
	fmt.Printf("%#v, len = %d, cap = %d\n", slice, len(slice), cap(slice))

	slice = append(slice[:2], append([]int{100}, slice[2:]...)...) // insert 100 to index 2 (0-base)
	fmt.Printf("%#v, len = %d, cap = %d\n", slice, len(slice), cap(slice))

	slice = append(slice[:3], append([]int{101, 102}, slice[3:]...)...) // insert slice{101, 102} to index 3
	fmt.Printf("%#v, len = %d, cap = %d\n", slice, len(slice), cap(slice))

	// insert 99 to index 2, with better performance
	slice = append(slice, 0)
	copy(slice[3:], slice[2:])
	slice[2] = 99
	fmt.Printf("%#v, len = %d, cap = %d\n", slice, len(slice), cap(slice))

	// insert slice {97, 98} to index 2, with better performance
	insertSlice := []int{97, 98}
	slice = append(slice, insertSlice...)
	copy(slice[2 + len(insertSlice):], slice[2:])
	copy(slice[2:], insertSlice)
	fmt.Printf("%#v, len = %d, cap = %d\n", slice, len(slice), cap(slice))

	// delete elements from index 2 to index 7
	slice = append(slice[:2], slice[8:]...)
	fmt.Printf("%#v, len = %d, cap = %d\n", slice, len(slice), cap(slice))

	// delete the first element
	slice = slice[:copy(slice, slice[1:])]
	fmt.Printf("%#v, len = %d, cap = %d\n", slice, len(slice), cap(slice))
}