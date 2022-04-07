// Hands on exercise #1
// ● create a func with the identifier foo that returns an int
// ● create a func with the identifier bar that returns an int and a string
// ● call both funcs
// ● print out their results

package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	return 1
}

func bar() (int, string) {
	return 2, "hola"
}
