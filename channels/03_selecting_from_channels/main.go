// Using the "select" clause we can describe different behaviors with
// multiple channels based on conditions.

package main

import "fmt"

func main() {
	fmt.Printf("---- Program starts succesfully. :D ----\n\n")

	even := make(chan int)
	oddN := make(chan int)
	quit := make(chan int)

	// Sender function that will send values to the passed channels
	// corresponding on a condition (if the value is even or odd).
	go send(even, oddN, quit, 100)

	// Receiver function that will describe condition based behaviors
	// on the passed channels.
	receive(even, oddN, quit)

	fmt.Printf("---- Program exited succesfully. :D ----\n")
}

// -------- Sender Function --------

// This function will send values to the passed channels based on a condition.
// There will be no need to close the channels since the "quit" channel will
// notify the application to stop in the receiver function.
// The parameter "n" will determine how many values are going to be passed.
func send(e, o, q chan<- int, n int) {
	// Passing values based of if they are even or odd numbers.
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}

	// After all the values have been passed the "quit" channel will receive
	// a value and that will notify the receive funcion to stop receiving.
	q <- 0
}

// -------- Sender Function --------

// -------- Receiver Function --------

// This function will receive all the values prior passed and if the "quit"
// channel receives a values makes the function stop.
// Based from which channel the values comes (even or oddN channels) the
// "select" clause will determine what block of code will be executed.
func receive(e, o, q <-chan int) {
	// Infinite loop since the "quit" channel will be the determing factor
	// on when this function will stop reading values from the other channels.
	for {
		select {
		// The ":=" syntaxis will declare the variable "v" as the value
		// received from each channel.
		case v := <-e:
			fmt.Println("The following value comes from the 'even' channel:", v)
		case v := <-o:
			fmt.Println("The following value comes from the 'oddN' channel:", v)
		case <-q:
			// If the "quit" channel would get multiple values which would
			// determine different behaviors here could be an if condition
			// or switch clause that will execute those behaviors.
			fmt.Printf("The 'quit' channel received a value, therefore I will stop looking for more values.\n\n")
			return
		}
	}
}

// -------- Receiver Function --------
