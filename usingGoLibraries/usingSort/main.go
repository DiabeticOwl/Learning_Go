package main

import (
	"fmt"
	"sort"
)

func main() {
	// Slice of ints.
	si := []int{3, 7, 4, 2, 6, 80}

	// Slice of strings.
	ss := []string{"Javier", "Albino", "Marrón"}

	fmt.Println("---- Printing the slices unordered. ----")
	fmt.Println(si)
	fmt.Println(ss)
	fmt.Println("---- Printing the slices unordered. ----")

	fmt.Println()

	fmt.Println("---- Printing the slices ordered. ----")
	// Sort's "Ints" and "Strings" methods are inplace,
	// and theorically work with the addresses of the
	// underlying arrays of the slices internally,
	// so they don't return anything.
	sort.Ints(si)
	sort.Strings(ss)

	fmt.Println(si)
	fmt.Println(ss)
	fmt.Println("---- Printing the slices ordered. ----")
}
