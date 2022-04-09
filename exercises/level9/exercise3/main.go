// Hands-on exercise #3
// 	● Using goroutines, create an incrementer program
// 		○ have a variable to hold the incrementer value
// 		○ launch a bunch of goroutines
// 			■ each goroutine should
// 				● read the incrementer value
// 					○ store it in a new variable
// 				● yield the processor with runtime.Gosched()
// 				● increment the new variable
// 				● write the value in the new variable back to the incrementer
// 				  variable
// 	● use waitgroups to wait for all of your goroutines to finish
// 	● the above will create a race condition.
// 	● Prove that it is a race condition by using the -race flag
// 	● if you need help, here is a hint: https://play.golang.org/p/FYGoflKQej

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	counter := 0
	nRoutines := 50

	fmt.Println("---- Initial value of the counter:", counter, "----")

	wg.Add(nRoutines)

	for i := 0; i < nRoutines; i++ {
		go func() {
			fmt.Println("Counter's current value:", counter)

			c := counter
			runtime.Gosched()
			c++
			counter = c

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("---- Final value of the counter:", counter, "----")
}
