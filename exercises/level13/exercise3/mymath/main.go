// mymath is a package that describes an implementation of the centered average. This is a subset of the moving average calculation.
// More information can be found here:
// https://en.wikipedia.org/wiki/Moving_average
package mymath

import "sort"

// CenteredAvg computes the average of a list of numbers
// after removing the largest and smallest values.
func CenteredAvg(xi []int) float64 {
	sort.Ints(xi)
	xi = xi[1:(len(xi) - 1)]

	n := 0
	for _, v := range xi {
		n += v
	}

	f := float64(n) / float64(len(xi))
	return f
}
