// Hands-on exercise #1
// ● Using a COMPOSITE LITERAL:
// 	● create an ARRAY which holds 5 VALUES of TYPE int
// 	● assign VALUES to each index position.
// ● Range over the array and print the values out.
// ● Using format printing
// ○ print out the TYPE of the array

package main

import "fmt"

func main() {
	var array [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Printf("Type of the array: %T\n", array)

	for i, value := range array {
		fmt.Printf("Value of index %v is: %v\n", i, value)
	}

}
