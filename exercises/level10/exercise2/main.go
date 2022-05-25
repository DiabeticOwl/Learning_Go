// Hands-on exercise #2
// ● Get this code running:
// 	○ https://play.golang.org/p/oB-p3KMiH6
// 		■ solution: https://play.golang.org/p/isnJ8hMMKg
// 	○ https://play.golang.org/p/_DBRueImEq
// 		■ solution: https://play.golang.org/p/mgw750EPp4

package main

import (
	"fmt"
)

// https://play.golang.org/p/oB-p3KMiH6
func firstCodePart() {
	// This line creates the error because it instantiates a sender only channel
	// and later tried to receive a value from it.
	// cs := make(chan<- int)

	// Solution:
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}

// https://play.golang.org/p/_DBRueImEq
func secondCodePart() {
	// This line creates the error because it instantiates a receiver only channel
	// and later tried to send a value to it.
	// cr := make(<-chan int)

	// Solution:
	cr := make(chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}

func main() {
	fmt.Println("---- First Code Part ----")
	// https://play.golang.org/p/oB-p3KMiH6
	firstCodePart()
	// https://play.golang.org/p/oB-p3KMiH6
	fmt.Println("---- First Code Part ----")

	fmt.Println()

	fmt.Println("---- Second Code Part ----")
	// https://play.golang.org/p/_DBRueImEq
	secondCodePart()
	// https://play.golang.org/p/_DBRueImEq
	fmt.Println("---- Second Code Part ----")
}
