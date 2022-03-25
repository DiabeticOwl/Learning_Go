// Hands-on exercise #5
// Create a variable of type string using a raw string literal. Print it.

package main

import (
	"fmt"
)

func main() {
	s := `
		So this is a
		raw string
		literal.
		Interesting.
		"Hola Mundo"
	`

	fmt.Println(s)
}
