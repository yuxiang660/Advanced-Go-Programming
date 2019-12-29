package main

import "fmt"

func main() {
	fmt.Println("Slice Definiation")
	
	var sliceDef1 []int // equal to nil, not exist
	fmt.Printf("%#v, len = %d, cap = %d\n", sliceDef1, len(sliceDef1), cap(sliceDef1))

	sliceDef2 := []int{} // unequal to nil, empty
	fmt.Printf("%#v, len = %d, cap = %d\n", sliceDef2, len(sliceDef2), cap(sliceDef2))

	sliceDef3 := []int{1, 2, 3}
	fmt.Printf("%#v, len = %d, cap = %d\n", sliceDef3, len(sliceDef3), cap(sliceDef3))

	sliceDef4 := sliceDef3[:2]
	fmt.Printf("%#v, len = %d, cap = %d\n", sliceDef4, len(sliceDef4), cap(sliceDef4))

	sliceDef5 := sliceDef3[:1:2]
	fmt.Printf("%#v, len = %d, cap = %d\n", sliceDef5, len(sliceDef5), cap(sliceDef5))

	sliceDef6 := sliceDef3[:0]
	fmt.Printf("%#v, len = %d, cap = %d\n", sliceDef6, len(sliceDef6), cap(sliceDef6))

	sliceDef7 := make([]int, 3)
	fmt.Printf("%#v, len = %d, cap = %d\n", sliceDef7, len(sliceDef7), cap(sliceDef7))

	sliceDef8 := make([]int, 2, 3)
	fmt.Printf("%#v, len = %d, cap = %d\n", sliceDef8, len(sliceDef8), cap(sliceDef8))

	sliceDef9 := make([]int, 0, 3)
	fmt.Printf("%#v, len = %d, cap = %d\n", sliceDef9, len(sliceDef9), cap(sliceDef9))
}
