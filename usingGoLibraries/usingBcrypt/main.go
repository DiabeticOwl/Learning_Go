// Here we are going to use the "bcrypt" package that is under
// the experimental section of packages in Go (golang.org/x/...).
// In order to use these experimental libraries we need to run
// the command "go get" followed by the desired library.
// And, after the import is written in the code, run
// "go mod tidy" so it will clean our package.

package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pword := "passwordToHash"

	// The "GenerateFromPassword(password []byte, cost int) ([]byte, error)"
	// method needs a slice of bytes as the password to encrypt/hash
	// and an int referring to the cost that the computer is going
	// to take in order to generate the hashed password.
	// The cost number ranges from 4 to 31, as it is described in the documentation.
	// https://pkg.go.dev/golang.org/x/crypto@v0.0.0-20220331220935-ae2d96664a29/bcrypt
	sb, err := bcrypt.GenerateFromPassword([]byte(pword), 6)

	if err != nil {
		fmt.Println(err)
	}

	// The standard and raw password.
	fmt.Println(pword)
	// The hashed password as a slice of bytes.
	// fmt.Println(sb)
	// The hashed password as a string.
	fmt.Println(string(sb))

	fmt.Println()

	// So we can comparte the hashed password with the raw one
	// (for authentication for instance) the
	// "CompareHashAndPassword(hashedPassword, password []byte) error"
	// method is useful.
	err = bcrypt.CompareHashAndPassword(sb, []byte("hola"))

	if err != nil {
		fmt.Println("You are not properly authenticated.")
		// So we can close the "main" function.
		return
	}

	fmt.Println("Welcome. :D")
}
