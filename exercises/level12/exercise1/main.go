// Hands-on exercise #1
// Create a dog package. The dog package should have an exported func “Years” which takes
// human years and turns them into dog years (1 human year = 7 dog years). Document your
// code with comments. Use this code in func main.
// 	● run your program and make sure it works
// 	● run a local server with godoc and look at your documentation
package main

import (
	"Level12/Exercise1/dog"
	"fmt"
)

func main() {
	d := dog.Dog{
		Name:     "Doggy",
		YearsOld: 5,
	}

	fmt.Printf("The dog called '%v' and with %v years old has %v dog years old.\n", d.Name, d.YearsOld, d.Years())
}
