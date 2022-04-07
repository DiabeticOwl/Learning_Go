// Hands-on exercise #4
// Create and use an anonymous struct.

package main

import "fmt"

func main() {
	anonStruct := struct {
		firstName string
		lastname  string
	}{
		firstName: "Johann",
		lastname:  "Cruz",
	}

	fmt.Println(anonStruct)
}
