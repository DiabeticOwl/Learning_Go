// Hands-on exercise #4
// Write a program that
// * assigns an int to a variable
// * prints that int in decimal, binary, and hex
// * shifts the bits of that int over 1 position to the left, and assigns that to a variable
// * prints that variable in decimal, binary, and hex

package main

import (
	"fmt"
)

func main() {
	a := 7

	fmt.Println("--------------------")

	fmt.Printf("Decimal: %d\n", a)
	fmt.Printf("Binary: %b\n", a)
	fmt.Printf("Hexadecimal: %#X\n", a)

	fmt.Println("--------------------")

	aShifted := a << 1

	fmt.Printf("Decimal: %d\n", aShifted)
	fmt.Printf("Binary: %b\n", aShifted)
	fmt.Printf("Hexadecimal: %#X\n", aShifted)

	fmt.Println("--------------------")
}
