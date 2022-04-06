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
	actalYear := time.Now().Year()
	birthYear := 1998

	fmt.Println("--- Years I have been alive ----")

	for actalYear >= birthYear {
		fmt.Println(actalYear)
		actalYear--
	}

	fmt.Println("--- Years I have been alive ----")
}
