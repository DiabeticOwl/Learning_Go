// Hands-on exercise #4
// ● Starting with this code, pull the values off the channel using a select statement

package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}

		q <- 1

		close(c)
	}()

	return c
}

func receive(c, q <-chan int) {
	for {
		select {
		case v, _ := <-c:
			fmt.Printf("The value %v is even and comes from the \"c\" channel.\n", v)
		case <-q:
			fmt.Println("\"q\" channel has returned a value, therefore I will quit.")
			return
		}
	}
}
