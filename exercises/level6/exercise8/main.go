// Hands-on exercise #8
// ● Create a func which returns a func
// ● assign the returned func to a variable
// ● call the returned func

package main

import "fmt"

func greeter() func() string {
	return greeting
}

func greeting() string {
	return fmt.Sprintf("Greetings and salutations.")
}

func main() {
	g := greeter()
	fmt.Println(g())
}
