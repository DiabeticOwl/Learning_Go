# Level 7 - Pointers

## Hands-on exercise #1
* Create a value and assign it to a variable.
* Print the address of that value.

## Hands-on exercise #2
* create a person struct
* create a func called “changeMe” with a *person as a parameter
    * change the value stored at the *person address
        * *important*
            * to dereference a struct, use (*value).field
                * `p1.first`
                * `(*p1).first`
            * “As an exception, if the type of x is a named pointer type and `(*x).f` is a valid selector expression denoting a field (but not a method), `x.f` is shorthand for `*x).f`.”
                * https://golang.org/ref/spec#Selectors
