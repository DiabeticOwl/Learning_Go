// Here we are creating a custom sort set of behaviors
// in order to order our typed slices as we want to.
package main

import (
	"fmt"
	"sort"
)

// First we create the type of the previously mentioned
// typed slice.
type person struct {
	first string
	last  string
	age   uint16
}

// ---------------- Ordering by age. ----------------

// Then we create the typed slice so we can later
// attach specific behaviors that are already
// written in the "sort" package (and will be modified
// at our leisure) so the typed slice becomes of type
// "Interface" (not to mix with the "interface" type,
// this is an interface type within the "sort" package).
type byAge []person

// "Len()" seems to be just the length of the array.
// I assume it is used to determine the following
// "i" and "j" values.
func (a byAge) Len() int { return len(a) }

// "Swap(i, j int)" seems to exchange the values in the
//"i" and "j" positions in the array accordingly
// to some logic in the "sort" package.
func (a byAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// "Less(i, j int)" seems to have a heavy weight into the
// "Swap(i, j int)" logic by determining which of the
// values in the "i" and "j" positions is lesser than the other.
func (a byAge) Less(i, j int) bool { return a[i].age < a[j].age }

// ---------------- Ordering by age. ----------------

// ---------------- Ordering by full name. ----------------
// Here we are modifying the behaviors so the "Sort" method
// uses our logic to sort by the full name (first + " " + last)
// of the "person" type.
// Notice how the behaviors have the same name but by their
// receivers being different, Go knows that these are different
// behaviors than the ones described before.

type byName []person

func (a byName) Len() int      { return len(a) }
func (a byName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byName) Less(i, j int) bool {
	iFullName := a[i].first + " " + a[i].last
	jFullName := a[j].first + " " + a[j].last

	return iFullName < jFullName
}

// ---------------- Ordering by full name. ----------------

func main() {
	p1 := person{
		first: "Johann",
		last:  "Cruz",
		age:   28,
	}

	p2 := person{
		first: "Johann",
		last:  "Villas",
		age:   30,
	}

	p3 := person{
		first: "Luis",
		last:  "Tiburcio",
		age:   22,
	}

	p4 := person{
		first: "Ivan",
		last:  "Cruz",
		age:   10,
	}

	// Now we can populate a slice of "person"...
	people := []person{p1, p2, p3, p4}

	// ...Check that it is unordered accordingly to our needs...
	fmt.Println(people)

	// Converting it to the typed slice in which we attached
	// our custom "sort" package's behaviors, so then the
	// "Sort" method (which expects an "Interface" value)
	// can work with our behaviors.
	// This will sort the array as we desire.
	sort.Sort(byAge(people))

	// Here we can check the new order.
	fmt.Println(people)

	// Here we can sort by the new logic we described...
	sort.Sort(byName(people))

	// ...And print the ordered array.
	fmt.Println(people)
}
