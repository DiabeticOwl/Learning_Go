// Hands-on exercise #9
// Using the code from the previous example, add a record to your map. Now print the map out
// using the “range” loop

package main

import "fmt"

func main() {
	people := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	people[`cruz_johan`] = []string{`PCs`, `RPGs`, `Cats`, `Yogurt`}

	for key, value := range people {
		fmt.Println("Key: ", key)
		for i, favThing := range value {
			fmt.Printf("\tIndex: %v \t Value: %v\n", i, favThing)
		}
	}
}
