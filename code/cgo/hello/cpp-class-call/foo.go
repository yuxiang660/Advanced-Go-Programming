package main

// #cgo LDFLAGS: -L. -lfoo
// #include "cFoo.h"
import "C"

type goFoo struct {
	f C.cFoo
}

func openFoo() goFoo {
	var foo goFoo
	foo.f = C.openCFoo()
	return foo
}

func (foo goFoo) close() {
	C.closeCFoo(foo.f)
}

func (foo goFoo) foo() {
	C.cfoo(foo.f)
}

func main() {
	foo := openFoo()
	foo.foo()
	foo.close()
}
