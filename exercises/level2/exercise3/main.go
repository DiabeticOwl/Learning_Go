// Hands-on exercise #3
// Create TYPED and UNTYPED constants. Print the values of the constants.

package main

import (
	"fmt"
)

const (
	a int = 5
	b     = "hola"
)

func main() {
	fmt.Printf("a: %v, type: %T\n", a, a)
	fmt.Printf("b: %v, type: %T\n", b, b)
}
