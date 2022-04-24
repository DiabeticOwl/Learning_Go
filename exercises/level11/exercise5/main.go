// Hands-on exercise #5
// We are going to learn about testing next. For this exercise, take a moment and see how much
// you can figure out about testing by reading the testing documentation & also Caleb Doxsey’s
// article on testing. See if you can get a basic example of testing working.

package main

import (
	"math/rand"
	"time"
)

// GetNumber returns a random number.
func GetNumber() int32 {
	rand.Seed(time.Now().UnixNano())

	num := rand.Int31()

	// Random choice of whether n is negative or not.
	// rand.Int31() returns a random positive 'int32' number,
	// rand.Intn(n) returns a random positive number within a [0, n) range.
	// Being n equal to 1 the next statement will return a number between 0 and 1,
	// and that will be used to determine the sign of the number "num".
	isPositive := rand.Intn(2)

	if isPositive == 0 {
		num = -num
	}

	return num
}

func main() {
	// Check "basic_test.go" for the test itself.
}
