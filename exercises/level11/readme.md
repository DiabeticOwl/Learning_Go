# Level 11 - Error handling

## Hands-on exercise #1
Start with [this](https://play.golang.org/p/3W69TH4nON) code. Instead of using the blank identifier, make sure the code is checking and handling the error.

## Hands-on exercise #2
Start with [this](https://play.golang.org/p/9a1IAWy5E6) code. Create a custom error message using “fmt.Errorf”.

## Hands-on exercise #3
Create a struct “customErr” which implements the builtin.error interface. Create a func “foo” that has a value of type error as a parameter.
Create a value of type “customErr” and pass it into “foo”. If you need a hint, [here](https://play.golang.org/p/L5QWV8-p11) is one.

## Hands-on exercise #4
Starting with [this](https://play.golang.org/p/wlEM1tgfQD) code, use the `sqrt.Error` struct as a value of type error.
If you would like, use these numbers for your coordinates:
* lat "50.2289 N"
* long "99.4656 W"

## Hands-on exercise #5
We are going to learn about testing in level 13. For this exercise, take a moment and see how much you can figure out about testing by reading the [testing documentation](http://godoc.org/testing) & also [Caleb Doxsey’s article on testing](http://www.golang-book.com/books/intro/12). See if you can get a basic example of testing working.
