// My adventure with functions in Go.
package main

import (
	"fmt"
)

func main() {
	// A Go application needs an entry point, which is
	// this one and comes in the "package" main.

	fmt.Println("----- Functions -----")

	sayGreeting("Hello", "Johann")

	fmt.Println("----- Functions -----")

	fmt.Println()

	fmt.Println("----- Functions | With pointers -----")

	greeting, name := "Hello", "Johann"

	// After initializing the variables we can send
	// their pointers through the function by using
	// the address keyword (&).
	sayGreetingWithPointers(&greeting, &name)

	// Let's print the variable that was changed.
	fmt.Println(name)

	fmt.Println("----- Functions | With pointers -----")

	fmt.Println()

	fmt.Println("----- Functions | Variadic parameters -----")

	sum("The sum is:", 1, 2, 3, 4, 5, 6)

	fmt.Println("----- Functions | Variadic parameters -----")

	fmt.Println()

	fmt.Println("----- Functions | Return -----")

	s := sumThatReturns(1, 2, 3, 4, 5, 6)
	fmt.Println("The sum is:", s)

	fmt.Println("----- Functions | Return -----")

	fmt.Println()

	fmt.Println("----- Functions | Return pointer -----")

	sPointer := sumThatReturnsAPointer(1, 2, 3, 4, 5, 6)
	fmt.Println("The sum is:", sPointer)
	fmt.Println("The sum is:", *sPointer)

	fmt.Println("----- Functions | Return pointer -----")

	fmt.Println()

	fmt.Println("----- Functions | Return multiple values -----")

	d, err := divide(5.0, 4.36)

	// This is a good example of a controlflow that
	// doesn't "panics" and that is more used.
	if err != nil {
		fmt.Println(err)
		// Here the "return" keyword is normally
		// used so the application ends.
		return
	}
	fmt.Println(d)

	fmt.Println("----- Functions | Return multiple values -----")

	fmt.Println()

	fmt.Println("----- Functions | Anonymous functions -----")

	func() {
		fmt.Println("This is an anonymous function.")
	}()

	fmt.Println("----- Functions | Anonymous functions -----")

	fmt.Println()

	fmt.Println("----- Functions | Functions as variables -----")

	// Functions can also be used as variables for
	// mainly optimization usages.
	divideVar := func(a, b float64) (float64, error) {
		if b == 0 {
			return 0, fmt.Errorf("cannot divide by zero")
		}

		return a / b, nil
	}

	//	// To write the same function but a bit more verbose
	//	// we can do the following, where the first list
	//	// of comma separated names are the types of the
	//	// expected parameters and the second one are the
	//	// type of the expected types of values to return.
	// var divideVar2 func(float64, float64) (float64, error)
	// divideVar2 = func(a, b float64) (float64, error) {
	// 	if b == 0 {
	// 		return 0, fmt.Errorf("cannot divide by zero")
	// 	}

	// 	return a / b, nil
	// }

	e, err := divideVar(3, 8.12547)

	if err != nil {
		fmt.Println(err)
		// Here the "return" keyword is normally
		// used so the application ends.
		return
	}

	fmt.Println(e)

	// f, err := divideVar2(4.5, 8.12547)
	f, err := divideVar(4.5, 8.12547)

	if err != nil {
		fmt.Println(err)
		// Here the "return" keyword is normally
		// used so the application ends.
		return
	}

	fmt.Println(f)

	fmt.Println("----- Functions | Functions as variables -----")

	fmt.Println()

	fmt.Println("----- Functions | Methods -----")

	g := greeter{
		greeting: "Greetings and salutations my dear",
		name:     "Johann",
	}

	g.greet()

	g.salute()
	fmt.Println("The new name is:", g.name)

	h := intAdder(8)
	fmt.Println(h.add(6))
	fmt.Println(h.add(6, 6))

	fmt.Println("----- Functions | Methods -----")
	fmt.Println("----- Functions -----")
}

// Go allows us to have a comma separated list of
// paramaters.
// It also gives us the ability of list various
// parameters of the same type so the code would be
// cleaner.
func sayGreeting(greeting, name string) {
	fmt.Println(greeting, name)
}

func sayGreetingWithPointers(greeting, name *string) {
	// After receiving the pointer variables we can
	// use them as we are used to by derefering them.
	fmt.Println(*greeting, *name)

	// This change here will after the variable
	// created in the main function since, instead
	// of being a copy of the variable passed
	// it is a pointer to that variable.
	*name = "Cruz"
	fmt.Println(*name)
}

// In Go we can use variadic parameters that will
// work essentially like *args in Python.
// In other words a function with this type of
// parameter will encapsulate all variables with
// the same type in a slice.
// If we have multiple types of parameters the
// variadic paramters must be put at the end.
func sum(msg string, values ...int) {
	fmt.Println(values)

	result := 0

	for _, v := range values {
		result += v
	}

	fmt.Println(msg, result)
}

// This is the same function as "sum"
// (without the first parameter) but instead of directly
// print the result it returns it.
// The value that is returned is a copy of the
// already instantiated variable. Go then deletes
// everything created in the functions namespace.
func sumThatReturns(values ...int) int {
	fmt.Println(values)

	result := 0

	for _, v := range values {
		result += v
	}

	return result
}

// In this function we return a pointer instead of
// copying the variable passed to the "return" keyword
// and destroy it.
func sumThatReturnsAPointer(values ...int) *int {
	fmt.Println(values)

	result := 0

	for _, v := range values {
		result += v
	}

	return &result
}

// Functions in Go can return multiple values by
// wrapping their expected types in parenthesis
// and listing, separated by commas, the
// corresponding values.
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}

	return a / b, nil
}

type greeter struct {
	greeting string
	name     string
}

// Methods in Go use "receivers" which are the
// connection between an object type and the
// function itself.
// The syntaxis for this is the name of the
// receiver and its type inside parenthesis.
func (gr greeter) greet() {
	fmt.Println(gr.greeting, gr.name)
}

// Like before, Go will copy the entire structure
// of the object used as a receiver.
// This can be modified by converting the receiver
// into a pointer receiver.
// Here, the "*" symbol will dereference the reveiver
// variable on its own in the entire scope of the method.
func (gr *greeter) salute() {
	fmt.Println(gr.greeting, gr.name)
	gr.name = "Cruz"
}

// In this example we add a method to a primitive type.
type intAdder int

func (a intAdder) add(values ...intAdder) int {
	for _, v := range values {
		a += v
	}

	return int(a)
}
