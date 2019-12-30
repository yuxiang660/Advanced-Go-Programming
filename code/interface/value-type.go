package main

import "fmt"

type fooInterface interface {
	fooMethod()
}

type foo struct {
	value string
}

func (f *foo) fooMethod() {
	if f == nil {
		fmt.Println("f is <nil>")
		return
	}
	fmt.Println(f.value)
}

func main() {
	fmt.Println("Interface Value and Type")

	var myFooInterface fooInterface
	var helloFoo = &foo{"Hello"}
	var emptyFoo foo
	var emptyFooPtr *foo

	myFooInterface = helloFoo
	fmt.Printf("Concrete Struct: %v, %T\n", helloFoo, helloFoo)
	fmt.Printf("Not Empty: %v, %T\n", myFooInterface, myFooInterface)
	myFooInterface.fooMethod()

	myFooInterface = &emptyFoo
	fmt.Printf("Concrete Struct: %v, %T\n", &emptyFoo, &emptyFoo)
	fmt.Printf("Empty Foo: %v, %T\n", myFooInterface, myFooInterface)
	myFooInterface.fooMethod()

	myFooInterface = emptyFooPtr
	fmt.Printf("Concrete Struct: %v, %T\n", emptyFooPtr, emptyFooPtr)
	fmt.Printf("Empty Foo Pointer: %v, %T\n", myFooInterface, myFooInterface)
	myFooInterface.fooMethod()
}
