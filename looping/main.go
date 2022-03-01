package main

import (
	"fmt"
)

func main() {
	fmt.Println("----- Loops -----")
	fmt.Println("----- Loops | Initial -----")

	for i := 0; i < 5; i = i + 2 {
		fmt.Println(i)
	}

	fmt.Println("----- Loops | Initial -----")

	fmt.Println()

	fmt.Println("----- Loops | With two variables -----")

	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	fmt.Println("----- Loops | With two variables -----")

	fmt.Println()

	fmt.Println("----- Loops | Existing initializer -----")

	i := 0
	for ; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println(i)

	fmt.Println("----- Loops | Existing initializer -----")

	fmt.Println()

	fmt.Println("----- Loops | While loop -----")

	j := 0
	for j < 10 {
		fmt.Println(j)
		j += 2
	}

	fmt.Println("----- Loops | While loop -----")

	fmt.Println()

	fmt.Println("----- Loops | Infinite loop -----")

	k := 0
	for {
		fmt.Println(k)
		k++

		if k == 5 {
			break
		}
	}

	fmt.Println("----- Loops | Infinite loop -----")

	fmt.Println()

	fmt.Println("----- Loops | continue keyword -----")

	for l := 0; l < 10; l++ {
		if l%2 != 0 {
			continue
		}

		fmt.Println(l)
	}

	fmt.Println("----- Loops | continue keyword -----")

	fmt.Println()

	fmt.Println("----- Loops | Nested Loop -----")

	// The word "Loop", followed by a colon, creates a label.
	// These things are very useful thanks to their ability
	// of being called in other blocks and used with breaks.
Loop:
	for l := 1; l < 3; l++ {
		for m := 1; m < 3; m++ {
			fmt.Println(l * m)

			if l*m >= 3 {
				break Loop
			}
		}
	}

	fmt.Println("----- Loops | Nested Loop -----")

	fmt.Println()

	fmt.Println("----- Loops | For Range Loop -----")

	// Just like a for loop with a dictionary in Python,
	// collections can be iterated with key and value pairs.
	s := []int{1, 2, 3, 4, 5, 6}
	for k, v := range s {
		fmt.Println(k, v)
	}

	fmt.Println()

	statePopulations := map[string]int{
		"California":   39237836,
		"Texas":        29527941,
		"Florida":      21781128,
		"New York":     19835913,
		"Pennsylvania": 12964056,
		"Illinois":     12671469,
		"Ohio":         11780017,
	}

	for k, v := range statePopulations {
		fmt.Println(k, "-", v)
	}

	fmt.Println()

	w := "Hello World!"
	for k, v := range w {
		fmt.Println(k, v, string(v))
	}

	fmt.Println("----- Loops | For Range Loop -----")

	fmt.Println("----- Loops -----")
}
