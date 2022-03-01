package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// The "defer" keyword makes that an statement will
	// be run after all the statements in the function
	// are executed but before the function returns a value.
	// This control flow is very useful for variable cleansing
	// after a function is called.

	fmt.Println("---- defer | Robots Example ----")

	res, err := http.Get("http://www.google.com/robots.txt")

	if err != nil {
		log.Fatal(err)
	}

	// Here is used for telling go that after is done with
	// the main function it should close this resource.
	defer res.Body.Close()

	robots, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", robots)

	fmt.Println("---- defer | Robots Example ----")

	fmt.Println()

	fmt.Println("---- defer | Argument when its called ----")
	// "defer" takes the value of the arguments at the time
	// that the keyword is called.
	a := "start"
	defer fmt.Println(a)
	a = "end"

	fmt.Println("---- defer | Argument when its called ----")
	fmt.Println()

	// fmt.Println("---- panic ----")

	// // In go there are no exceptions like in other programming
	// // languages but there is "panic". Which, as it reads,
	// // is a keywords that allows us to say "hey, if something
	// // happens panic and stop running."

	// // Use "panic" only when the application can no longer
	// // continue, like an zero division.
	// // "panic" will stop the current function and therefore
	// // execute all the deffered functions that were called
	// // before the "panic".

	// // If one of the deffered functions is a handler (recover)
	// // then the rest of the aplication will continue running.

	// fmt.Println("I'm going to panic...")
	// panic("I don't feel good...")
	// fmt.Println("Its over now.")

	// fmt.Println("---- panic ----")

}
