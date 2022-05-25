package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("----- Ifs -----")

	if true {
		fmt.Println("This is true!")
	}

	statePopulations := map[string]int{
		"California":   39237836,
		"Texas":        29527941,
		"Florida":      21781128,
		"New York":     19835913,
		"Pennsylvania": 12964056,
		"Illinois":     12671469,
		"Ohio":         11780017,
	}

	if state, ok := statePopulations["California"]; ok {
		fmt.Printf("The state of California has a population of %v. \n", state)
	}

	// Boolean operators are the same as Python.
	// Logical operators are || for or, && for and.

	// Go has a concept called short circuiting in
	// which will stop evaluating parts of an if
	// bool statement on the moment that the entire
	// bool would be logically speaking, true.
	// An example of this is an if statement with ors.

	// Ifs and elses clauses are very similar to C#.

	// Working with floating numbers and ifs clauses
	// are not always exact so an extra logical step
	// is required.
	myNum := 0.19498

	if myNum == math.Pow(math.Sqrt(myNum), 2) {
		fmt.Println("They are the same number!")
	} else {
		fmt.Println("They are different!")
	}

	// Here it is saying that if the division between
	// the number passed and itself to the square root
	// and the to the power of 2 is a number that,
	// after subtracted 1 (as equal numbers division
	// will result in 1), its absolute value is less
	// than a threshold then the statement is true.
	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.0001 {
		fmt.Println("They are the same number!")
	} else {
		fmt.Println("They are different!")
	}

	fmt.Println("----- Ifs -----")

	fmt.Println()

	fmt.Println("----- Switches -----")

	// Switches in Go are almost the same as in C# but
	// they allow multiple options in a single clause by
	// separating them with commas.
	// Cases can not overlap (be duplicate).
	switch 6 {
	case 1, 5, 10:
		fmt.Println("Numbers 1, 5 or 10")
	case 2, 3, 8:
		fmt.Println("Numbers 2, 3 or 8")
	case 6, 12, 25:
		fmt.Println("Numbers 6, 12 or 25")
	default:
		fmt.Println("Another number.")
	}

	// You can also use an initializer.
	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("Numbers 1, 5 or 10")
	case 2, 3, 8:
		fmt.Println("Numbers 2, 3 or 8")
	case 6, 12, 25:
		fmt.Println("Numbers 6, 12 or 25")
	default:
		fmt.Println("Another number.")
	}

	// There is also a "tag-less" switch syntax.
	// Here, since each case is a boolean operation
	// then they can overlap.
	i := 10

	switch {
	case i <= 10:
		fmt.Println("The number is less than or equal to 10.")
	case i <= 20:
		fmt.Println("The number is less than or equal to 20.")
	default:
		fmt.Println("The number is greater than 20.")
	}

	// If we want to pass to the other case statement
	// after something in the previous block happened,
	// we can use the "fallthrough" keyword.
	// Warning: this keyword is "logic-less", meaning
	// that even if the next boolean operation is false
	// it will run that block.
	j := 15

	switch {
	case j <= 10:
		fmt.Println("The number is less than or equal to 10.")
		fallthrough
	case j >= 20:
		fmt.Println("The number is less than or equal to 20.")
	default:
		fmt.Println("The number is greater than 20.")
	}

	// Using an interface (a type of variable that
	// can take any object, from a primitive to a
	// struct) we can switch by its generated type.
	var k interface{} = [3]int{}

	switch k.(type) {
	case int:
		fmt.Println("k is an int.")
	case float64:
		fmt.Println("k is a float64.")
	case string:
		fmt.Println("k is a string.")
	case [2]int:
		fmt.Println("k is a [2]int array.")
	case [3]int:
		fmt.Println("k is a [3]int array.")
	default:
		fmt.Println("k is another type.")
	}

	// We can use the "break" keyword to terminate a
	// switch case block early.
	// This keyword is implicit at the end of each block.
	var l interface{} = [8]int{}

	switch l.(type) {
	case int:
		fmt.Println("l is an int.")
	case float64:
		fmt.Println("l is a float64.")
	case string:
		fmt.Println("l is a string.")
	case [2]int:
		fmt.Println("l is a [2]int array.")
	case [3]int:
		fmt.Println("l is a [3]int array.")
	default:
		fmt.Println("l is another type.")
		break
		fmt.Println("Thanks for reading this!")
	}

	fmt.Println("----- Switches -----")
}
