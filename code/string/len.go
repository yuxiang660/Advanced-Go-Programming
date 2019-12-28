package main

import (
	"unsafe"
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Lenght of a String")

	s := "hello, world"
	s2 := s[:5]
	s3 := s[7:]

	fmt.Println("s:", s, "s2:", s2, "s3:", s3)
	fmt.Println("len(s)", (*reflect.StringHeader)(unsafe.Pointer(&s)).Len)
	fmt.Println("len(s2)", len(s2))
	fmt.Println("len(s3)", len(s3))
}
