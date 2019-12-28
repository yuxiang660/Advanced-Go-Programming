package main

import "fmt"

// Enable Chinese in CMD before runing the program.
// Run the command in CMD: "chcp 936"
func main() {
	fmt.Println("String UTF8")
	
	s := "Hello, 世界"
	fmt.Println(s)
	for i, c := range s {
		fmt.Printf("index: %d, %c\n", i, c)
	}
	fmt.Println("Print String with index")
	for i := 0; i < len(s); i++ {
		fmt.Printf("index: %d, %c\n", i, s[i])
	}

	fmt.Printf("%#v\n", []byte("Hello, 世界"))
	fmt.Println("\xe4\xb8\x96")
	fmt.Println("\xe7\x95\x8c")
	for i, c := range []byte("Hello, 世界") {
		fmt.Printf("index: %d, %c\n", i, c)
	}

	const placeOfInterest = `⌘`

    fmt.Printf("plain string: ")
    fmt.Printf("%s", placeOfInterest)
    fmt.Printf("\n")

    fmt.Printf("quoted string: ")
    fmt.Printf("%+q", placeOfInterest)
    fmt.Printf("\n")

    fmt.Printf("hex bytes: ")
    for i := 0; i < len(placeOfInterest); i++ {
        fmt.Printf("%x ", placeOfInterest[i])
    }
    fmt.Printf("\n")
}