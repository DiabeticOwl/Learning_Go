package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "Tym",
		Last:  "Burton",
		Age:   45,
	}

	p2 := person{
		First: "Myllenna",
		Last:  "Burton",
		Age:   42,
	}

	people := []person{p1, p2}

	fmt.Println(people)

	sb, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(sb)
	fmt.Println(string(sb))

	var people2 []person

	// Unmarshal takes a slice of bytes and the address
	// of where to store the data.
	// Normally it is useful to check the struct type
	// of the data using the following page:
	// https://mholt.github.io/json-to-go/
	err = json.Unmarshal(sb, &people2)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(people2)
}
