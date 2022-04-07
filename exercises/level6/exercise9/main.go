// Hands-on exercise #9
// A “callback” is when we pass a func into a func as an argument. For this exercise,
// ● pass a func into a func as an argument

package main

import "fmt"

func greeter(g func() string) string {
	return g()
}

func greeting() string {
	return fmt.Sprintf("Greetings and salutations.")
}

func main() {
	g := greeter(greeting)
	fmt.Println(g)
}
