// Hands-on exercise #2
// This exercise will reinforce our understanding of method sets:
// 	● create a type person struct
// 	● attach a method speak to type person using a pointer receiver
// 		○ *person
// 	● create a type human interface
// 		○ to implicitly implement the interface, a human must have the speak method
// 	● create func “saySomething”
// 		○ have it take in a human as a parameter
// 		○ have it call the speak method
// 	● show the following in your code
// 		○ you CAN pass a value of type *person into saySomething
// 		○ you CANNOT pass a value of type person into saySomething
// 	● here is a hint if you need some help
// 		○ https://play.golang.org/p/FAwcQbNtMG

// Receivers Values
// -----------------------------------------------
// (t T) T and *T
// (t *T) *T

package main

import "fmt"

// ------------------------- person -------------------------

type person struct {
	first string
	last  string
	age   int
}

func (p *person) speak() {
	fmt.Printf("Hello, I am %v.\n", p.first+" "+p.last)
}

// ------------------------- person -------------------------

// ------------------------- human -------------------------

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

// ------------------------- human -------------------------

func main() {
	p1 := person{
		first: "Johann",
		last:  "Cruz",
		age:   28,
	}

	// As stated before I can't pass a VALUE of TYPE person to the "saySomething" function.
	// Go will say that TYPE person doesn't implement human since the method set that
	// we created attaches to this TYPE through a pointer.
	// saySomething(p1)

	// On other hand, if we pass an address to the VALUE of TYPE person Go will understand it.
	saySomething(&p1)
}
