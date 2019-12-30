package main

import (
	"C"
	"fmt"
)

//export SayHello
func SayHello(s *C.char) {
	fmt.Print(C.GoString(s))
}
