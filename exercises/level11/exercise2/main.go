// Hands-on exercise #2
// Start with this code. Create a custom error message using “fmt.Errorf”.

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

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))

}

// // Before my edits:
// // toJSON needs to return an error also
// func toJSON(a interface{}) []byte {
// 	bs, _ := json.Marshal(a)
// }

// toJSON After my edits:
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	// // This following line is just for testing purposes.
	// err = fmt.Errorf("hi")
	if err != nil {
		err = fmt.Errorf("The JSON could not be Marshalized - %v", err)
		return bs, err
	}

	return bs, nil
}
