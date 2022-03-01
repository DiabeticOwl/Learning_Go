package main

import (
	"fmt"
)

func main() {
	// Booleans.
	bool := false

	fmt.Println("----- Booleans -----")
	fmt.Printf("%v, %T\n", bool, bool)
	fmt.Println("----- Booleans -----")
	fmt.Println()

	// Integers.
	var n uint8 = 64

	fmt.Println("----- Integers -----")
	fmt.Printf("%v, %T\n", n, n)

	// Integers. Bit operators.
	a := 10 // Binary representation: 1010
	b := 3  // Binary representation: 0011

	fmt.Println("----- Integers | Bit operators -----")

	// "AND" bit operator: check the bit in each position
	// of both numbers and if it's 1 return 1, else 0.
	// In this example: 0010.
	fmt.Printf("a & b = %v or %4b\n", a&b, a&b)

	// "OR" bit operator: check the bit in each position
	// of both numbers and if it's 1 in at least one of the numbers
	// return 1, else 0.
	// In this example: 1011.
	fmt.Printf("a | b = %v or %4b\n", a|b, a|b)

	// "XOR" bit operator: check the bit in each position
	// of both numbers and if it's 1 in only one of the numbers
	// return 1, else 0.
	// In this example: 1001.
	fmt.Printf("a ^ b = %v or %4b\n", a^b, a^b)

	// "AND NOT" bit operator: check the bit in each position
	// of both numbers and if it's 1 in the first one and 0
	// in the second one return 1, else 0.
	// In this example: 1000.
	fmt.Printf("a &^ b = %v or %4b\n", a&^b, a&^b)

	fmt.Println("----- Integers | Bit operators -----")

	fmt.Println("----- Integers | Bit shifting -----")

	a = 8 // 2^3

	fmt.Printf("a << 2 = %v\n", a<<2) // 2^3 * 2^2
	fmt.Printf("a >> 2 = %v\n", a>>2) // 2^3 / 2^2
	// Since we are working on the backend with bits then
	// there are no decimal points.
	fmt.Printf("a >> 4 = %v\n", a>>4) // 2^3 / 2^4

	fmt.Println("----- Integers | Bit shifting -----")
	fmt.Println("----- Integers -----")

	fmt.Println()

	fmt.Println("----- Floats -----")

	f := 3.14
	f = 18e42
	f = 1468e-2

	fmt.Printf("%v, %T\n", f, f)

	c := 3.5
	d := 0.962

	fmt.Printf("c + d = %v\n", c+d)
	fmt.Printf("c - d = %v\n", c-d)
	fmt.Printf("c * d = %v\n", c*d)
	fmt.Printf("c / d = %v\n", c/d)

	fmt.Println("----- Floats -----")
	fmt.Println()

	fmt.Println("----- Complex Numbers -----")

	var comp1 complex128 = complex(1, 2)
	comp2 := 3 + 8.7i

	fmt.Printf("1 + 2i = %v, %T\n", comp1, comp1)
	fmt.Printf("Real Part: %v, %T\n", real(comp1), real(comp1))
	fmt.Printf("Imaginary Part: %v, %T\n", imag(comp1), imag(comp1))

	fmt.Println()

	fmt.Printf("(1 + 2i) + (3 + 8.7i) = %v\n", comp1+comp2)
	fmt.Printf("(1 + 2i) - (3 + 8.7i) = %v\n", comp1-comp2)
	fmt.Printf("(1 + 2i) * (3 + 8.7i) = %v\n", comp1*comp2)
	fmt.Printf("(1 + 2i) / (3 + 8.7i) = %v\n", comp1/comp2)

	fmt.Println("----- Complex Numbers -----")

	fmt.Println()
	fmt.Println("----- Strings (Collection of utf-8 characters) -----")

	s := "this is a string"
	fmt.Printf("'%v', %T\n", s, s)
	// Strings in Go are aliases for bytes.
	fmt.Printf("'%v', %T\n", s[2], s[2])
	fmt.Printf("'%v', %T\n", string(s[2]), s[2])

	fmt.Println("----- Strings | Collection of bytes -----")

	bC := []byte(s)
	fmt.Printf("'%v', %T\n", bC, bC)

	fmt.Println("----- Strings | Collection of bytes -----")
	fmt.Println("----- Strings (Collection of utf-8 characters) -----")

	fmt.Println()

	fmt.Println("----- Runes (Collection of utf-32 characters) -----")

	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)

	fmt.Println("----- Runes (Collection of utf-32 characters) -----")
}
