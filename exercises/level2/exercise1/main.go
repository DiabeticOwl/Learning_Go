// Hands-on exercise #1
// Write a program that prints a number in decimal, binary, and hex

package main

import (
	"fmt"
)

func main() {
	x := 86

	fmt.Printf("Decimal: %d\n", x)
	fmt.Printf("Binary: %b\n", x)
	fmt.Printf("Hexadecimal: %#X\n", x)
}
