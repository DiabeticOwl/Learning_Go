// Using the "range" keyword and "close" function we can iterate through channels.

package main

import "fmt"

// -------- Sender Function --------

// This function sends an int-typed value to the passed channel.
// The "n" parameter describes how many values will be send to the channel.
// After all the values (or the desired behavior happened to our channel)
// are passed the channel needs to close so the receiver function gets
// notified that it needs to stop looking for more values, causing a
// channel block.
func foo(c chan<- int, n int) {
	for i := 0; i < n; i++ {
		c <- i
	}

	close(c)
}

// -------- Sender Function --------

// -------- Receiver Function --------

// This function receives a range of int-typed values from the passed channel
// and prints them. When it is notified of the channel being closed stops
// reading from it.
func bar(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

// -------- Receiver Function --------

func main() {
	// The channel "c" is instantiated without a buffer.
	c := make(chan int)

	// Sender of values function. This will be sent to another gorutine.
	go foo(c, 25)

	// Receiver of values. Since the prior goroutine performs in parallel
	// of the "main" goroutine, as values are added to the channel, this
	// function is extracting and performing the desired behavior with them.
	bar(c)
}
