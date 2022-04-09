// Hands-on exercise #5
// Fix the race condition you created in exercise #3 by using package atomic

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	var counter int64
	nRoutines := 50

	fmt.Println("---- Initial value of the counter:", counter, "----")

	wg.Add(nRoutines)

	for i := 0; i < nRoutines; i++ {
		go func() {
			// Atomic will yield the value and increment it by 1 automatically.
			atomic.AddInt64(&counter, 1)

			// If we read the counter value without loading it from atomic
			// (or before the value is incremented) then other goroutines
			// could access it before it would be incremented.
			fmt.Println("Counter's value after increment:", atomic.LoadInt64(&counter))

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("---- Final value of the counter:", counter, "----")
}
