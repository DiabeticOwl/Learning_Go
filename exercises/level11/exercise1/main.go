// Hands-on exercise #1
// Start with this code. Instead of using the blank identifier, make sure the code is checking and
// handling the error.

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	// Before my edits:
	// bs, _ := json.Marshal(p1)
	// After:
	bs, err := json.Marshal(p1)
	if err != nil {
		// Since there is nothing to print exit the application.
		log.Fatalln(err)
	}

	fmt.Println(string(bs))

}
