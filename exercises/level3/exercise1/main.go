// Hands-on exercise #1
// Print every number from 1 to 10,000

package main

import (
	// It will be needed to run the command
	// "go get golang.org/x/text/language" after "go mod init" is ran.
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	p := message.NewPrinter(language.English)

	for i := 1; i <= 1e4; i++ {
		p.Println(i)
	}
}
