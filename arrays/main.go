package main

import "fmt"

func main() {
	fmt.Println("----- Arrays -----")
	grades := [...]int{97, 85, 93, 96}

	fmt.Printf("Grades: %v\n", grades)

	fmt.Println()

	var students [6]string
	students[0] = "Johann"
	students[1] = "Cruz"
	students[2] = "Antonio"

	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Student's count: %v\n", len(students))

	fmt.Println()

	fmt.Println("----- Arrays | Matrix -----")

	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}

	fmt.Println(identityMatrix)
	fmt.Println("----- Arrays | Matrix -----")

	fmt.Println()

	fmt.Println("----- Arrays | Trying to reference -----")
	// Arrays in Go are viewed as unique values,
	// therefore assigning one to another variable
	// will automatically create a copy of
	// the passed array, instead of referring
	// to the initial one.
	a := [...]float32{2.09, 3., 8.1}
	b := a
	b[1] = 7

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("----- Arrays | Trying to reference -----")

	fmt.Println()

	fmt.Println("----- Arrays | Pointing -----")
	// For "pointing" to the original array
	// we need to use pointers (the "&" symbol
	// before the variable name).
	c := [...]int{5, 8, 4, 6, 10}
	d := &c
	d[1] = 7
	d[3] = 2

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println("----- Arrays | Pointing -----")
	fmt.Println("----- Arrays -----")

	fmt.Println()

	fmt.Println("----- Slices -----")
	// For avoiding to have to give a size to our
	// array we can use what are called "slices".
	// They do not have the behavior of arrays
	// in which Go, by default, creates a new
	// copy instead of pointing to a space in memory.
	// Referring to the inicial slice.
	// This difference comes from the fact that
	// a slice will, behind the scenes, create an array.
	e := []int{5, 8, 4}
	f := e

	f[1] = 3

	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("Length: %v\n", len(e))
	fmt.Printf("Capacity: %v\n", cap(e))

	fmt.Println()

	fmt.Println("----- Slices | Slicing -----")

	g := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	h := g[:]   // Slice of all elements.
	i := g[3:]  // Slice from 4th element to end.
	j := g[:6]  // Slice of all first 6 elements.
	k := g[3:6] // Slice the 4th, 5th and 6th elements.

	g[3] = 11

	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)

	fmt.Println("----- Slices | Slicing -----")

	fmt.Println()

	fmt.Println("----- Slices | make function -----")

	l := make([]int, 3)
	fmt.Println(l)
	fmt.Printf("Length: %v\n", len(l))
	fmt.Printf("Capacity: %v\n", cap(l))

	m := make([]int, 3, 100)
	fmt.Println(m)
	fmt.Printf("Length: %v\n", len(m))
	fmt.Printf("Capacity: %v\n", cap(m))

	fmt.Println("----- Slices | make function -----")

	fmt.Println()

	fmt.Println("----- Slices | append function -----")

	n := []int{}
	fmt.Println(n)
	fmt.Printf("Length: %v\n", len(n))
	fmt.Printf("Capacity: %v\n", cap(n))

	n = append(n, 1)

	fmt.Println(n)
	fmt.Printf("Length: %v\n", len(n))
	fmt.Printf("Capacity: %v\n", cap(n))

	n = append(n, 2, 3)

	fmt.Println(n)
	fmt.Printf("Length: %v\n", len(n))
	fmt.Printf("Capacity: %v\n", cap(n))

	// "..." works like * in Python.
	n = append(n, []int{4, 5, 6}...)

	fmt.Println(n)
	fmt.Printf("Length: %v\n", len(n))
	fmt.Printf("Capacity: %v\n", cap(n))

	fmt.Println("----- Slices | append function -----")

	fmt.Println("----- Slices -----")
}
