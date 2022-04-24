package main

import (
	"testing"
)

// TestGreaterThanOne will check whether the number that the function
// called "GetNumber" returns is positive or not.
func TestGreaterThanOne(t *testing.T) {
	// Getting number.
	num := GetNumber()

	if num < 0 {
		t.Errorf("Number '%v' is negative. A positive number is expected.", num)
	}
}
