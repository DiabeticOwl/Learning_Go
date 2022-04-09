// Hands-on exercise #1
// ● in addition to the main goroutine, launch two additional goroutines
// 	○ each additional goroutine should print something out
// ● use waitgroups to make sure each goroutine finishes before your program exists

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func foo() {
	fmt.Println("---- I'm the foo function. ----")
	for i := 0; i < 10; i++ {
		fmt.Println(i + 1)
	}
	fmt.Println("---- I'm the foo function. ----")

	// Here we say to the WaitGroup that this function is "done".
	wg.Done()
}

func bar() {
	fmt.Println("---- I'm the bar function. ----")
	for i := 0; i < 10; i++ {
		fmt.Println(i + 1)
	}
	fmt.Println("---- I'm the bar function. ----")

	// Here we say to the WaitGroup that this function is "done".
	wg.Done()
}

func main() {
	fmt.Println("------------------------- Start -------------------------")
	fmt.Println("\tArch:", runtime.GOARCH)
	fmt.Println("\tOS:", runtime.GOOS)
	fmt.Println("\tNumber of CPUs:", runtime.NumCPU())
	fmt.Println("\tNumber of subroutunes:", runtime.NumGoroutine())
	fmt.Println("------------------------- Start -------------------------")

	fmt.Println()

	// Here we say to the WaitGroup that there will be two additional goroutunes.
	wg.Add(2)
	// Here we say Go to create the goroutunes, each with the "go" keyword.
	go foo()
	go bar()

	fmt.Println()

	fmt.Println("------------------------- Middle -------------------------")
	fmt.Println("\tNumber of CPUs:", runtime.NumCPU())
	fmt.Println("\tNumber of subroutunes:", runtime.NumGoroutine())
	fmt.Println("------------------------- Middle -------------------------")

	// Here we say to the WaitGroup to make the application wait until the functions
	// in the goroutunes are "done".
	wg.Wait()

	fmt.Println("------------------------- End -------------------------")
	fmt.Println("\tNumber of CPUs:", runtime.NumCPU())
	fmt.Println("\tNumber of subroutunes:", runtime.NumGoroutine())
	fmt.Println("------------------------- End -------------------------")
}
