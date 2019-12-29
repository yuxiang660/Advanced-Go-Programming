package main

import (
	"unsafe"
	"fmt"
	"reflect"
	"unicode/utf8"
)

func str2bytes(s string) []byte {
	p := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		p[i] = s[i]
	}
	return p
}

func bytes2str(bytes []byte) (s string) {
	// Create data to make sure the string will be read-only.
	// If using bytes directly, compiler will use the bytes for the string underneath.
	data := make([]byte, len(bytes))
	for i, c := range bytes {
		data[i] = c
	}

	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	hdr.Data = uintptr(unsafe.Pointer(&data[0]))
	hdr.Len = len(bytes)
	return s
}

func str2runes(s string) []rune {
	var p []int32
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		p = append(p, int32(r))
		s = s[size:]
	}
	return []rune(p)
}

func runes2string(runes []int32) string {
	var bytes []byte
	buf := make([]byte, 3)
	for _, r := range runes {
		n := utf8.EncodeRune(buf, r)
		bytes = append(bytes, buf[:n]...)
	}
	return string(bytes)
}

func main() {
	fmt.Println("String to Bytes and Bytes to String")

	const placeOfInterest = `b⌘e`

	
	bytes := []byte(placeOfInterest)
	fmt.Printf("Range Bytes, len(bytes) = %d\n", len(bytes))
	for i, c := range bytes {
		fmt.Printf("index: %d, %c\n", i, c)
	}

	str := placeOfInterest
	fmt.Printf("Range String, len(str) = %d\n", len(str))
	for i, c := range str {
		fmt.Printf("index: %d, %c\n", i, c)
	}

	runes := []rune(placeOfInterest)
	fmt.Printf("Range Runes, len(runes) = %d\n", len(runes))
	for i, c := range runes {
		fmt.Printf("index: %d, %c\n", i, c)
	}

	// Bytes to String
	stringFromBytes := bytes2str(bytes)
	fmt.Println("Range String from Bytes:")
	for i, c := range stringFromBytes {
		fmt.Printf("index: %d, %c\n", i, c)
	}
	
	// String to Bytes
	bytesFromStr := str2bytes(str)
	fmt.Println("Range Bytes from String:")
	for i, c := range bytesFromStr {
		fmt.Printf("index: %d, %c\n", i, c)
	}

	// Runes to String
	stringFromRunes := runes2string(runes)
	fmt.Println("Range String from Runes:")
	for i, c := range stringFromRunes {
		fmt.Printf("index: %d, %c\n", i, c)
	}

	// String to Runes
	runesFromStr := str2runes(str)
	fmt.Println("Range Runes from String:")
	for i, c := range runesFromStr {
		fmt.Printf("index: %d. %c\n", i, c)
	}
}
