package main

import (
	"fmt"
	"strconv"
)

// Varibles declaration clustering!
var (
	name    string = "Johann"
	surname string = "Cruz"
	age     int    = 23
)

var (
	c int = 0
	i int = 2
)

func main() {
	fmt.Printf("%v, %T\n", i, i)
	i := 3

	fmt.Printf("My name is %v %v\n", name, surname)

	fmt.Printf("%v, %T\n", name, name)
	fmt.Printf("%v, %T\n", surname, surname)
	fmt.Printf("%v, %T\n", i, i)

	var cSt string = strconv.Itoa(c)

	fmt.Printf("%v, %T\n", cSt, cSt)
}
