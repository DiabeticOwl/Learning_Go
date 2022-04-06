// Hands-on exercise #7
// Building on the previous hands-on exercise, create a program that uses “else if” and “else”.

package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	me := Person{
		name: "Johann Cruz",
		age:  16,
	}

	if me.age >= 18 {
		fmt.Printf("The person named '%v' is an adult.\n", me.name)
	} else if me.age < 18 {
		fmt.Printf("The person named '%v' is not an adult.\n", me.name)
	} else {
		fmt.Println("this won't ever run.")
	}
}
