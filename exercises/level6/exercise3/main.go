// Hands-on exercise #3
// Use the “defer” keyword to show that a deferred func runs after the func containing it exits.

package main

import "fmt"

func main() {
	bar()
}

func foo() {
	fmt.Println("Hola, corro al final gracias al keyword 'defer'.")
}

func bar() {
	defer foo()

	fmt.Println("De manera que yo corro primero. :D")
}
