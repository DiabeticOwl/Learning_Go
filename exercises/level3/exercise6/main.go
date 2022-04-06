// Hands-on exercise #6
// Create a program that shows the “if statement” in action.

package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	me := Person{
		name: "Johann Cruz",
		age:  23,
	}

	if me.age >= 18 {
		fmt.Printf("The person named '%v' is an adult.\n", me.name)
	}
}
