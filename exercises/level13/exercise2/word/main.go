// Package word provides some functions useful for working witj strings.
package word

import "strings"

// no need to write an example for this one
// writing a test for this one is a bonus challenge; harder

// UseCount returns a map with each unique string separated by a space
// in s and the amount of repetitions that it has.
func UseCount(s string) map[string]int {
	// strings.Fields(s) is equivallent to strings.Split(s, " ")
	xs := strings.Fields(s)

	m := make(map[string]int)

	// v is instantiated once and then its value is incremented
	// each time is found.
	for _, v := range xs {
		m[v]++
	}

	return m
}

// Count returns the amount of string that are separated by spaces in s.
func Count(s string) int {
	return len(strings.Fields(s))
}
