// Hands-on exercise #9
// Create a program that uses a switch statement with the switch expression specified as a
// variable of TYPE string with the IDENTIFIER “favSport”.

package main

import "fmt"

func main() {
	var favSport interface{} = "Swimming"

	switch favSport.(type) {
	case string:
		fmt.Println("This case is the string one.")
	default:
		fmt.Println("This case is the one of all other types.")
	}
}
