// Hands-on exercise #7
// ● write a program that
// 	○ launches 10 goroutines
// 		■ each goroutine adds 10 numbers to a channel
// 	○ pull the numbers off the channel and print them

package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)

	// I need a WaitGroup so the following close statement doesn't
	// run before all the goroutines are done.
	var wg sync.WaitGroup

	go func() {
		for i := 0; i < 10; i++ {
			wg.Add(1)

			go func() {
				// Run "wg.Done()" after the literal function ends.
				defer wg.Done()
				for i := 0; i < 10; i++ {
					c <- i
				}
			}()
		}

		// Wait until all goroutines are done and then close.
		wg.Wait()
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}
