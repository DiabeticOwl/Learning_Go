// My first program that uses channels.

// A channel is a mechanism that allows the share of values through goroutines
// in a clean, powerful and more reliable way than using WaitGroups and GoScheds.

// A channel can be created with the following sintaxis:
// make(chan TYPE, n) where "TYPE" refers to the primitive type of values that
// the channel will work with, and "n" to the buffer that it will have.

// The buffer of a channel describes how many values the channel can have
// and, although it is very useful, it is advised (for optimization purposes
// as I'm aware of) to prefer using channels with short life cycle.
// Channels without a buffer can not work on a singular goroutine (causing
// a so called "channel block") since as they are instructed to send the
// value that they receive in the instant that they receive it, so multiple
// goroutines would be necessary in order to satisfy this design.

// Channels can be bidirectional, meaning that can both receive and send values,
// or directional, meaning that can do one of the prior behaviors.
// The verbs "receive" means that Go will receive values from the channel
// and "send" will describe a channel that allows Go to send values to it.
// A bidirectional channel can be converted into a directional channel but a
// directional one can not be a bidirectional.

// In order to declare a directional channel the following sintaxis is required:
// 	* Receiver channel:
// 		- c := make(<-chan TYPE, n)
// 		- func XXX(<-c TYPE){}
// 	* Sender channel:
// 		- c := make(chan <- TYPE, n)
// 		- func XXX(c <- TYPE){}

package main

import "fmt"

// -------- Sender Function --------

// This function sends an int-typed value to the passed channel.
func foo(c chan<- int) {
	c <- 6
}

// -------- Sender Function --------

// -------- Receiver Function --------

// This function receives an int-typed value from the passed channel and prints it.
func bar(c <-chan int) {
	fmt.Println(<-c)
}

// -------- Receiver Function --------

func main() {
	// The channel "c" is instantiated without a buffer.
	c := make(chan int)

	// Sender of values function. This will be sent to another gorutine.
	go foo(c)

	// receiver of values. As the prior goroutine performs in parallel
	// of the "main" go routine, satisfying the channel that received
	// a value in it by extracting the expected value.
	bar(c)
}
