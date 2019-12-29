package main

import (
	"fmt"
	"io"
	"bytes"
	"os"
)

/*
"Fprintf" function signature:
func Fprintln(w io.Writer, a ...interface{}) (n int, err error)

"io.Writer" is an interface:
type io.Write interface {
	Write(p []byte) (n int, err error)
}
*/

type upperWriter struct {
	io.Writer
}

func (p *upperWriter) Write(data []byte) (n int, err error) {
	return p.Writer.Write(bytes.ToUpper(data))
}

func main() {
	fmt.Println("Upper Writer")
	fmt.Fprintln(&upperWriter{os.Stdout}, "hello, world")
}
