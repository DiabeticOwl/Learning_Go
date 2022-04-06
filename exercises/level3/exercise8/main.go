// Hands-on exercise #8
// Create a program that uses a switch statement with no switch expression specified.

package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("This case is the false one.")
	case true:
		fmt.Println("This case is the true one.")
	}
}
