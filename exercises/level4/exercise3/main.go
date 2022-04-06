// Hands-on exercise #3
// Using the code from the previous example, use SLICING to create the following new slices
// which are then printed:
// 	● [42 43 44 45 46]
// 	● [47 48 49 50 51]
// 	● [44 45 46 47 48]
// 	● [43 44 45 46 47]

package main

import "fmt"

func main() {
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Printf("Slice #1 of the slice variable is: %v\n", slice[:5])
	fmt.Printf("Slice #2 of the slice variable is: %v\n", slice[5:])
	fmt.Printf("Slice #3 of the slice variable is: %v\n", slice[2:7])
	fmt.Printf("Slice #4 of the slice variable is: %v\n", slice[1:6])
}
