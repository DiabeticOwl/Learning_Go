// This is the same code as before but implementing the "comma ok" idiom,
// which is very used in all Go.

package main

import "fmt"

func main() {
	fmt.Printf("---- Program starts successfully. :D ----\n\n")

	even := make(chan int)
	oddN := make(chan int)
	quit := make(chan int)

	// Sender function that will send values to the passed channels
	// corresponding on a condition (if the value is even or odd).
	go send(even, oddN, quit, 100)

	// Receiver function that will describe condition based behaviors
	// on the passed channels.
	receive(even, oddN, quit)

	fmt.Printf("---- Program exited successfully. :D ----\n")
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

	// After all the values have been passed the "quit" channel will be closed,
	// giving it a zero value (the zero value of the type of the channel).
	// In this scenario it will return a 0.
	close(q)
}

// -------- Sender Function --------

// -------- Receiver Function --------

// This function will receive all the values prior passed and if the "quit"
// channel closes makes the function stop.
// Based from which channel the values comes (even or oddN channels) the
// "select" clause will determine what block of code will be executed.
func receive(e, o, q <-chan int) {
	// Infinite loop since the "quit" channel will be the determining factor
	// on when this function will stop reading values from the other channels.
	for {
		select {
		// The ":=" syntax will declare the variable "v" as the value
		// received from each channel.
		case v := <-e:
			fmt.Println("The following value comes from the 'even' channel:", v)
		case v := <-o:
			fmt.Println("The following value comes from the 'oddN' channel:", v)
		case _, ok := <-q:
			// A receiver channel will return the corresponding value and a
			// boolean normally called "ok" as the state of it
			// (open or closed i.e. true or false).
			// This allows us to describe more behaviors if needed.
			if !ok {
				fmt.Printf("The 'quit' channel closed, therefore I will stop looking for more values.\n\n")
				return
			} // else { do something else }
		}
	}
}

// -------- Receiver Function --------
