package main

import (
	"fmt"
	"strings"
	"os"
)

/*
"Fprintf" function signature:
func Fprintln(w io.Writer, a ...interface{}) (n int, err error)

"a" should implement "Stringer" interface, which describe the type as a string.
type Stringer interface {
    String() string
}
*/

type upperString string

func (s upperString) String() string {
	return strings.ToUpper(string(s))
}

func main() {
	fmt.Println("Upper String")
	fmt.Fprintln(os.Stdout, upperString("hello, world"))
}
