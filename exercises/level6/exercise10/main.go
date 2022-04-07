// Hands-on exercise #10
// Closure is when we have “enclosed” the scope of a variable in some code block. For this
// hands-on exercise, create a func which “encloses” the scope of a variable.

package main

import "fmt"

func xEncapsulated() int {
	var x int = 0

	for i := 0; i < 10; i++ {
		x++
		fmt.Println("The value of the x in this function is now:", x)
	}

	return x
}

func main() {
	var x int = 6

	fmt.Println("The value of the x in main is:", x)

	xEncapsulated()

	fmt.Println("The value of the x in main is:", x)
}
