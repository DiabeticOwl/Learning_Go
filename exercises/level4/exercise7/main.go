// Hands-on exercise #7
// Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional
// slice:
// "James", "Bond", "Shaken, not stirred"
// "Miss", "Moneypenny", "Helloooooo, James."
// Range over the records, then range over the data in each record.

package main

import "fmt"

func main() {
	matrix := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."},
	}

	fmt.Printf("%v", matrix)

	for _, row := range matrix {
		fmt.Println("\nRow: ", row)
		for _, cell := range row {
			fmt.Println(cell)
		}
	}
}
