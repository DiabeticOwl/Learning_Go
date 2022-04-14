// Hands-on exercise #1
// ● get this code working:
// 	○ using func literal, aka, anonymous self-executing func
// 		■ solution: https://play.golang.org/p/SHr3lpX4so
// 	○ a buffered channel
// 		■ solution: https://play.golang.org/p/Y0Hx6IZc3U

package main

import (
	"fmt"
)

func main() {
	// ---- using func literal, aka, anonymous self-executing func ----
	fmt.Println("---- Using Func Literal ----")

	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

	fmt.Println("---- Using Func Literal ----")
	// ---- using func literal, aka, anonymous self-executing func ----

	fmt.Println()

	// ---- using a buffered channel ----
	fmt.Println("---- Using a Buffered Channel ----")

	c = make(chan int, 1)

	c <- 42

	fmt.Println(<-c)

	fmt.Println("---- Using a Buffered Channel ----")
	// ---- using a buffered channel ----
}
