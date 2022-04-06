// Hands-on exercise #6
// Create a slice to store the names of all of the states in the United States of America. Use make
// and append to do this. Goal: do not have the array that underlies the slice created more than
// once. What is the length of your slice? What is the capacity? Print out all of the values, along
// with their index position, without using the range clause. Here is a list of the 50 states:
// ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, `
// Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, `
// Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, `
// Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`,
// ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`,
// ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, `
// Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,

package main

import "fmt"

func main() {
	// Slice of strings with 0 values and a capacity of 50 items.
	x := make([]string, 0, 50)
	x = append(
		x, "Alabama",
		"Alaska", "Arizona",
		"Arkansas", "California",
		"Colorado", "Connecticut",
		"Delaware", "Florida",
		"Georgia", "Hawaii",
		"Idaho", "Illinois",
		"Indiana", "Iowa",
		"Kansas", "Kentucky",
		"Louisiana", "Maine",
		"Maryland", "Massachusetts",
		"Michigan", "Minnesota",
		"Mississippi", "Missouri",
		"Montana", "Nebraska",
		"Nevada", "New Hampshire",
		"New Jersey", "New Mexico",
		"New York", "North Carolina",
		"North Dakota", "Ohio",
		"Oklahoma", "Oregon",
		"Pennsylvania", "Rhode Island",
		"South Carolina", "South Dakota",
		"Tennessee", "Texas",
		"Utah", "Vermont",
		"Virginia", "Washington",
		"West Virginia", "Wisconsin",
		"Wyoming",
	)

	fmt.Println(x)
	fmt.Printf("Length: %v, capacity: %v\n", len(x), cap(x))
}
