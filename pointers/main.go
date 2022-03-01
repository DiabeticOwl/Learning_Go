// Pointers are the location in memory of objects.
package main

import "fmt"

func main() {
	fmt.Println("----- Pointers -----")
	fmt.Println("----- Pointers | Primitives -----")

	// Primitives and slices, when assgined to another
	// variable, creates a copy of themselves insted
	// of point to the same object.
	a := 44
	b := a

	fmt.Println(a, b)

	a = 85

	fmt.Println(a, b)

	fmt.Println("----- Pointers | Primitives -----")

	fmt.Println()

	fmt.Println("----- Pointers | Object Type -----")

	var c int = 20
	// The '*' symbol is for describing the type of
	// the pointer and the '&' symbol for the variable
	// to point to.
	var d *int = &c

	// The value of d will be a string representing
	// the location in memory of the contents in
	// the variable c.
	fmt.Println(c, d)

	// If we put the "&" symbol (or "address") before
	// the variable c we can see the same value being
	// printed for both variables.
	fmt.Println(&c, d)

	// By using the "*" symbol (or "derefencing operator")
	// before the d variable we can find the value that
	// is located on the address in memory in question.
	fmt.Println(c, *d)

	// Since both values of the variables c and d are
	// connected changing the value in one will change
	// it in the other.
	c = 564
	fmt.Println(c, *d)

	fmt.Println("----- Pointers | Object Type -----")

	fmt.Println()

	fmt.Println("----- Pointers | Pointer Arithmetics -----")

	// In this section we can observe an example of
	// "Pointers Aritmetics" not working in Go due
	// to the decisions of maintaining this language
	// simple. For this kind of operations, the
	// package "unsafe" will become useful.

	// By declaring an integer array and creating
	// two pointers towards two of its values,
	// the fact that the addresses are very close
	// is visible.
	// In C, and other languages, would be possible
	// to add or substract to these addresses in
	// order to very optimally point to another
	// value in memory.
	arr := [3]int{1, 2, 3}
	arr1 := &arr[0]
	arr2 := &arr[1]

	fmt.Printf("%v %p %p\n", arr, arr1, arr2)

	fmt.Println("----- Pointers | Pointer Arithmetics -----")

	fmt.Println()

	fmt.Println("----- Pointers | Point to an unassigned value (Pointer type) -----")

	// Go allows us to point to even unassigned values
	// just by declaring the pointer type of the object
	// that we want to point at. Then point to where
	// we want to.
	// We can tell that this works thanks to the blue
	// underline in the pointer type declaration.
	var ms *myStruct
	ms = &myStruct{foo: 85}
	fmt.Println(ms)

	// Go gives us the option to initialize pointers
	// to empty objects with the "new" function.
	var ms2 *myStruct
	// The initial value of an empty pointer is "nil".
	fmt.Println(ms2)
	ms2 = new(myStruct)
	fmt.Println(ms2)

	// In order to access to the field "foo" contained
	// in "myStruct" and by using this pointer we need
	// to dereference it. This process is done using
	// the following syntax: (*pointer).foo .

	// This syntax is needed because the "*" symbol
	// has lower precedence than the "." in Go, meaning
	// that the language will first execute the "." and
	// then the "*".
	(*ms2).foo = 6
	fmt.Println(ms2)
	fmt.Println((*ms2).foo)

	// The syntax described before sets a limit in
	// Go's code readibility so the compiler
	// identifies what we are trying to do and
	// dereference the pointer for us without
	// the need of the syntax.
	ms2.foo = 8
	fmt.Println(ms2)
	fmt.Println(ms2.foo)

	fmt.Println("----- Pointers | Point to an unassigned value (Pointer type) -----")

	fmt.Println("----- Pointers -----")
}

type myStruct struct {
	foo int
}
