// Hands-on exercise #2
// Take the code from the previous exercise, then store the values of type person in a map with the
// key of last name. Access each value in the map. Print out the values, ranging over the slice.

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

	people := map[string]person{
		person1.lastName: person1,
		person2.lastName: person2,
	}

	for key, p := range people {
		fmt.Printf("---- %v ----\n", key)

		fmt.Printf("First Name: %v\n", p.firstName)
		fmt.Printf("Last Name: %v\n", p.lastName)
		fmt.Println("Favorite Ice Cream Flavors:")

		for _, flavor := range p.favIceCreamFlav {
			fmt.Printf("\t* %v\n", flavor)
		}

		fmt.Printf("---- %v ----\n", key)
	}
}
