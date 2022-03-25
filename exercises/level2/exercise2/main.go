// Hands-on exercise #2
// Using the following operators, write expressions and assign their values to variables:
// g. ==
// h. <=
// i. >=
// j. !=
// k. <
// l. >
// Now print each of the variables.

package main

import (
	"fmt"
)

func main() {
	a := 1
	b := 2

	aEQb := a == b
	aLEb := a <= b
	aGEb := a >= b
	aNEb := a != b
	aLTb := a < b
	aGTb := a > b

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("aEQb: %v\n", aEQb)
	fmt.Printf("aLEb: %v\n", aLEb)
	fmt.Printf("aGEb: %v\n", aGEb)
	fmt.Printf("aNEb: %v\n", aNEb)
	fmt.Printf("aLTb: %v\n", aLTb)
	fmt.Printf("aGTb: %v\n", aGTb)
}
