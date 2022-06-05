# Level 9 - Concurrency

## Hands-on exercise #1

* in addition to the main goroutine, launch two additional goroutines
  * each additional goroutine should print something out
* use waitgroups to make sure each goroutine finishes before your program exists

## Hands-on exercise #2

This exercise will reinforce our understanding of method sets:

* create a type person struct
* attach a method speak to type person using a pointer receiver
  * *person
* create a type human interface
  * to implicitly implement the interface, a human must have the speak method
* create func “saySomething”
  * have it take in a human as a parameter
  * have it call the speak method
* show the following in your code
  * you CAN pass a value of type *person into saySomething
  * you CANNOT pass a value of type person into saySomething
* here is a hint if you need some help
  * <https://play.golang.org/p/FAwcQbNtMG>

## Hands-on exercise #3

* Using goroutines, create an incrementor
 program
  * have a variable to hold the incrementor
     value
  * launch a bunch of goroutines
    * each goroutine should
      * read the incrementor
             value
      * store it in a new variable
      * yield the processor with runtime.Gosched()
      * increment the new variable
      * write the value in the new variable back to the incrementor
             variable
* use waitgroups to wait for all of your goroutines to finish
* the above will create a race condition.
* Prove that it is a race condition by using the -race flag
* if you need help, here is a hint: <https://play.golang.org/p/FYGoflKQej>

## Hands-on exercise #4

Fix the race condition you created in the previous exercise by using a mutex.

It makes sense to remove `runtime.Gosched()`.

## Hands-on exercise #5

Fix the race condition you created in exercise #4 by using package atomic.

## Hands-on exercise #6

Create a program that prints out your OS and ARCH. Use the following commands to run it

* go run
* go build
* go install
