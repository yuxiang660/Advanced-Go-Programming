package main

import (
	"fmt"
	"os"
	"io"
)

func main() {
	fmt.Println("Interface Implicit Conversion")

	f, _ := os.Open("./conversion.go")
	fmt.Printf("(*os.File) type: %T\n", f)
	
	var readCloser io.ReadCloser = f
	fmt.Printf("readCloser type: %T\n", readCloser)

	var reader io.Reader = readCloser
	fmt.Printf("reader type: %T\n", reader)

	var closer io.Closer = readCloser
	fmt.Printf("closer type: %T\n", closer)

	// Use Type Assertion to access to underlying concrete value
	var reader2 io.Reader = closer.(io.Reader)
	fmt.Printf("reader2 type: %T\n", reader2)

	f.Close()

	numInt := 1
	fmt.Printf("numInt type: %T\n", numInt)

	var numInt32 int32 = int32(numInt)
	fmt.Printf("type: %T\n", numInt32)

	var numInt64 int64 = int64(numInt)
	fmt.Printf("type: %T\n", numInt64)
}
