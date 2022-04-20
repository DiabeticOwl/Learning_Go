// Hands-on exercise #3
// Create a struct “customErr” which implements the builtin.error interface. Create a func “foo” that
// has a value of type error as a parameter. Create a value of type “customErr” and pass it into
// “foo”. If you need a hint, here is one.

package main

import (
	"fmt"
	"log"
)

type customErr struct {
	myErr string
}

func (cE customErr) Error() string {
	return fmt.Sprintf(cE.myErr)
}

func main() {
	myCustomErr := customErr{
		myErr: "This is my custom error.",
	}

	foo(myCustomErr)
}

func foo(err error) {
	if err != nil {
		log.Println(err)
	}
}
