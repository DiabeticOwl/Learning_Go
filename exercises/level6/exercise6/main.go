// Hands-on exercise #6
// ● Build and use an anonymous func

package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("Hello World! :D")
	}

	f()
}
