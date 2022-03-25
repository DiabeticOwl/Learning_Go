// Hands-on exercise #6
// Using iota, create 4 constants for the NEXT 4 years. Print the constant values.

package main

import (
	"fmt"
)

const (
	thisYear = 2022 + iota
	secondYear
	thirdYear
	fourthYear
)

func main() {
	fmt.Println(thisYear)
	fmt.Println(secondYear)
	fmt.Println(thirdYear)
	fmt.Println(fourthYear)
}
