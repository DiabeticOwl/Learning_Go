package main

import (
	"fmt"
)

// Enumerated constants.
// iota is a special symbol that,
// alongside a const block, can
// initialize a counter.
// In this block of constants x is
// assign iota and the rest is
// inferred to be the same type,
// continuing the counting.
const (
	x = iota + 1
	y
	w
	z
)

const (
	o = iota
	p
	q
	r
)

const (
	_ = iota // Underscore means that it will be ignored.
	// This creates a block of constants that will return
	// 1 * 2^(10*iota). This formula is the description of
	// bytes in the binary system used by computers.
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Println("----- Typed Constants -----")
	const myConst int = 42
	f := 85

	fmt.Printf("%v, %T\n", myConst*f, myConst*f)
	fmt.Println("----- Typed Constants -----")

	fmt.Println()

	fmt.Println("----- Untyped Constants -----")
	// Untyped Constants.
	// It will infer the type... (int)
	const myConst2 = 74
	var a int8 = 6

	// ..and at an operation with another type it will
	// convert it into the needed type.
	fmt.Printf("%v, %T\n", myConst2+a, myConst2+a)
	fmt.Println("----- Untyped Constants -----")

	fmt.Println()
	fmt.Println("----- Enumerated Constants -----")

	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", y, y)
	fmt.Printf("%v, %T\n", w, w)
	fmt.Printf("%v, %T\n", z, z)

	fmt.Println()

	fmt.Printf("%v, %T\n", o, o)
	fmt.Printf("%v, %T\n", p, p)
	fmt.Printf("%v, %T\n", q, q)
	fmt.Printf("%v, %T\n", r, r)

	fmt.Println()

	fileSize := 89454165489.
	fmt.Printf("%.2f GBs\n", fileSize/GB)

	fmt.Println("----- Enumerated Constants -----")
}
