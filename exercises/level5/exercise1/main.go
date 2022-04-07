// Hands-on exercise #1
// Create your own type “person” which will have an underlying type of “struct” so that it can store
// the following data:
// ● first name
// ● last name
// ● favorite ice cream flavors
// Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice
// which stores the favorite flavors.

package main

import "fmt"

type person struct {
	firstName       string
	lastName        string
	favIceCreamFlav []string
}

func main() {
	person1 := person{
		firstName:       "Johann",
		lastName:        "Cruz",
		favIceCreamFlav: []string{"Cacao", "Cinnamon"},
	}

	person2 := person{
		firstName:       "Fulano",
		lastName:        "Esteban",
		favIceCreamFlav: []string{"Mani", "Vainilla"},
	}

	people := []person{person1, person2}

	for i, p := range people {
		id := i + 1

		fmt.Printf("---- Person #%v ----\n", id)

		fmt.Printf("First Name: %v\n", p.firstName)
		fmt.Printf("Last Name: %v\n", p.lastName)
		fmt.Println("Favorite Ice Cream Flavors:")

		for _, flavor := range p.favIceCreamFlav {
			fmt.Printf("\t* %v\n", flavor)
		}

		fmt.Printf("---- Person #%v ----\n", id)
	}
}
