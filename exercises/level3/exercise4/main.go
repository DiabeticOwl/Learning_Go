// Hands-on exercise #4
// Create a for loop using this syntax
// ● for { }
// Have it print out the years you have been alive.

package main

import (
	"fmt"
	"time"
)

func main() {
	actualYear := time.Now().Year()
	birthYear := 1998

	fmt.Println("--- Years I have been alive ----")

	for {
		fmt.Println(actualYear)
		actualYear--

		if actualYear < birthYear {
			break
		}
	}

	fmt.Println("--- Years I have been alive ----")
}
