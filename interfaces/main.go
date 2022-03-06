// Interfaces describe behaviors unlike structs
// that describe data.
package main

import (
	"fmt"
)

func main() {
	fmt.Println("----- Interfaces -----")

	// The value that a interface brings is the
	// ability of receive any type object with
	// a behavior that was defined inside the interface.
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))

	fmt.Println()

	fmt.Println("----- Interfaces | With a primitive -----")

	myInt := IntCounter(0)
	var inc Incrementer = &myInt

	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}

	fmt.Println("----- Interfaces | With a primitive -----")

	fmt.Println()

	fmt.Println("----- Interfaces | Type Switches and Empty Interfaces -----")

	// An empty interface can work as an intermediate
	// between the varible that you are instantiating
	// to and the interfaces implemented in that variable.
	var i interface{} = "Hola"

	// Here we are extracting the data type of the variable i.
	switch i.(type) {
	case int:
		fmt.Println("i is an integer.")
	case string:
		fmt.Println("i is a string.")
	default:
		fmt.Println("I don't know what i is.")
	}

	fmt.Println("----- Interfaces | Type Switches and Empty Interfaces -----")

	fmt.Println()

	fmt.Println("----- Interfaces | Value and pointer receivers -----")

	// When we implement interfaces with value receivers
	// instead of pointer receivers the object to which
	// we implement must be casted as a value and not a
	// pointer.
	// On other hand, if one of the methods has a pointer
	// receiver, the object mush be casted as a pointer.
	var wc WriterCloser = &myWriterCloser{}
	fmt.Println(wc)

	fmt.Println("----- Interfaces | Value and pointer receivers -----")

	fmt.Println("----- Interfaces -----")
}

// Inside an interface we should write methods
// definition.
type Writer interface {
	Write([]byte) (int, error)
}

// The way to implement this interface is through
// a method that has the signature of our desired
// interface.
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

// Any type that has a method associated with it
// can have a interface implemented in it.
type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	// Since "ic" is a pointer we need to dereference it.
	*ic++
	return int(*ic)
}

// Using the "Writer" interface created before we can
// define another set of interfaces what works with
// each other.

type Closer interface {
	Close() error
}

// An interface that encapsules the behaviors
// of both Writer and Closer.
type WriterCloser interface {
	Writer
	Closer
}

// For simplicity sake let's create an empty struct...

type myWriterCloser struct{}

// ...And then it's methods (also technically empty).

func (mwc *myWriterCloser) Write(data []byte) (int, error) {
	return 0, nil
}

func (mwc myWriterCloser) Close() error {
	return nil
}
