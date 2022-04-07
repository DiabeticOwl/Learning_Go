// Hands-on exercise #4
// ● Create a user defined struct with
// 	○ the identifier “person”
// 	○ the fields:
// 		■ first
// 		■ last
// 		■ age
// ● attach a method to type person with
// 	○ the identifier “speak”
// 	○ the method should have the person say their name and age
// ● create a value of type person
// ● call the method from the value of type person

package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	person1 := person{
		first: "Johann",
		last:  "Cruz",
		age:   28,
	}

	person1.speak()
}

func (p person) speak() {
	fullName := p.first + " " + p.last
	fmt.Printf("Hello, my name is %v and I'm %v years old.\n", fullName, p.age)
}
