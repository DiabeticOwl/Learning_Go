// Hands-on exercise #4
// Fix the race condition you created in the previous exercise by using a mutex
// 	● it makes sense to remove runtime.Gosched()

package main

import (
	"fmt"
	"sync"
)

func main() {
	// Since these two variables are only going to be used in the main function
	// then it makes more sense to declare them here.
	var wg sync.WaitGroup
	var mu sync.Mutex

	counter := 0
	nRoutines := 50

	fmt.Println("---- Initial value of the counter:", counter, "----")

	wg.Add(nRoutines)

	for i := 0; i < nRoutines; i++ {
		go func() {
			// Locking the next chunk of code so others goroutines doesn't access
			// it until the one already working with it finishes.
			// Note: The lock should be before the following print because if it isn't
			// then other goroutines can access the counter variable, creating another
			// race condition.
			mu.Lock()

			fmt.Println("Counter's current value:", counter)

			c := counter
			c++
			counter = c

			// Unlocking the chunk of code.
			mu.Unlock()

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("---- Final value of the counter:", counter, "----")
}
