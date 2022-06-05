# Level 8 - Application of external libraries

## Hands-on exercise #1

Starting with [this](https://play.golang.org/p/_fQUGm9Utvl) code, marshal the []user to JSON. There is a little bit of a curve ball in this
hands-on exercise - remember to ask yourself what you need to do to EXPORT a value from a
package.

## Hands-on exercise #2

Starting with [this](https://play.golang.org/p/b_UuCcZag9) code, unmarshal the JSON into a Go data structure. Hint: <https://mholt.github.io/json-to-go/>

## Hands-on exercise #3

Starting with [this](https://play.golang.org/p/BVRZTdlUZ_) code, encode to JSON the []user sending the results to Stdout.
Hint: you will need to use `json.NewEncoder(os.Stdout).encode(v interface{})`.

## Hands-on exercise #4

Starting with [this](https://play.golang.org/p/H_q75mpmHW) code, sort the []int and []string for each person.

## Hands-on exercise #5

Starting with [this](https://play.golang.org/p/BVRZTdlUZ_) code, sort the []user by

* age
* last

Also sort each []string “Sayings” for each user print everything in a way that is pleasant.
