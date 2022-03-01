package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func isAnOption(option string, options [3]string) bool {
	for _, value := range options {
		if option == value {
			return true
		}
	}

	return false
}

func main() {
	options := [3]string{
		"Rock",
		"Paper",
		"Scissors",
	}

	// Initializing the seed for the random number generator.
	// It uses the time now as Unix time, an integer
	// version of the current time.
	rand.Seed(time.Now().UnixNano())

	// Picks a number from 0 to 3, not including 3.
	compSel := options[rand.Int31n(3)]

	var userSel string

	if len(os.Args) > 1 {
		userSel = strings.Title(os.Args[1])

		if !isAnOption(userSel, options) {
			fmt.Print("---- Please enter a valid option. (Rock, Paper or Scissors). ----\n")
		} else {
			// After checking which was the computer's choice...
			// key being the computer's selection as a human-readable way.
			// value being the computer's selection.
			switch {
			case compSel == "Rock" && userSel == "Scissors":
				fmt.Println("**** Computer Wins! Scissors was selected by the player while Rock was selected by the PC. ****")
			case compSel == "Rock" && userSel == "Paper":
				fmt.Println("**** Player Wins! Paper was selected by the player while Rock was selected by the PC. ****")
			case compSel == "Rock" && userSel == "Rock":
				fmt.Println("**** It's a draw! Rock was selected by both the PC and the player. ****")
			case compSel == "Paper" && userSel == "Rock":
				fmt.Println("**** Computer Wins! Rock was selected by the player while Paper was selected by the PC. ****")
			case compSel == "Paper" && userSel == "Scissors":
				fmt.Println("**** Player Wins! Scissors was selected by the player while Paper was selected by the PC. ****")
			case compSel == "Paper" && userSel == "Paper":
				fmt.Println("**** It's a draw! Paper was selected by both the PC and the player. ****")
			case compSel == "Scissors" && userSel == "Rock":
				fmt.Println("**** User Wins! Rock was selected by the player while Scissors was selected by the PC. ****")
			case compSel == "Scissors" && userSel == "Paper":
				fmt.Println("**** Computer Wins! Paper was selected by the player while Scissors was selected by the PC. ****")
			case compSel == "Scissors" && userSel == "Scissors":
				fmt.Println("**** It's a draw! Scissors was selected by both the PC and the player. ****")
			}
		}
		fmt.Println("---- Thanks for playing! ----")
	} else {
		fmt.Println("---- Welcome to Rock, Papers, Scissors; The game! ----")
		fmt.Println("---- Please insert one of the following options to play: Rock, Paper or Scissors. ----")
	}
}
