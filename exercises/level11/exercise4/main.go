// Hands-on exercise #4
// Starting with this code, use the sqrt.Error struct as a value of type error. If you would like, use
// these numbers for your
// 	● lat "50.2289 N"
// 	● long "99.4656 W"

package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// write your error code here
		retErr := sqrtError{
			lat:  "50.2289 N",
			long: "99.4656 W",
			err:  fmt.Errorf("Passed number is less than 0. This is not allowed."),
		}
		return 0, retErr
	}
	return 42, nil
}
