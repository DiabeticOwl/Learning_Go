// Hands-on exercise #3
// Create a for loop using this syntax
// ● for condition { }
// Have it print out the years you have been alive

package main

import (
	"fmt"
	"time"
)

func main() {
	actualYear := time.Now().Year()
	birthYear := 1998

	fmt.Println("--- Years I have been alive ----")

	for actualYear >= birthYear {
		fmt.Println(actualYear)
		actualYear--
	}

	fmt.Println("--- Years I have been alive ----")
}
