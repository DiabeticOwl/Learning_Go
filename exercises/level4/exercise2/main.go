// Hands-on exercise #2
// ● Using a COMPOSITE LITERAL:
// 	● create a SLICE of TYPE int
// 	● assign 10 VALUES
// ● Range over the slice and print the values out.
// ● Using format printing
// 	○ print out the TYPE of the slice

package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 7, 8, 9, 10}

	fmt.Printf("Type of the slice: %T\n", slice)

	for i, value := range slice {
		fmt.Printf("Value of index %v is: %v\n", i, value)
	}
}
