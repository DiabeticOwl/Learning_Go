// Hands-on exercise #2
// ● create a func with the identifier foo that
// 	○ takes in a variadic parameter of type int
// 	○ pass in a value of type []int into your func (unfurl the []int)
// 	○ returns the sum of all values of type int passed in
// ● create a func with the identifier bar that
// 	○ takes in a parameter of type []int
// 	○ returns the sum of all values of type int passed in

package main

import "fmt"

func main() {
	numsToAdd := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sum1 := foo(numsToAdd...)
	fmt.Println(sum1)

	sum2 := bar(numsToAdd)
	fmt.Println(sum2)
}

func foo(nums ...int) int {
	result := 0

	for _, num := range nums {
		result += num
	}

	return result
}

func bar(nums []int) int {
	result := 0

	for _, num := range nums {
		result += num
	}

	return result
}
