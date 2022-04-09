// Hands-on exercise #5
// Starting with this code, sort the []user by
// 	● age
// 	● last
// Also sort each []string “Sayings” for each user
// 	● print everything in a way that is pleasant

package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type byAge []user

func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type byLast []user

func (a byLast) Len() int           { return len(a) }
func (a byLast) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byLast) Less(i, j int) bool { return a[i].Last < a[j].Last }

// Prints a slice of users pleasantly.
func printUsers(users []user) {
	for i, u := range users {
		uId := i + 1
		fmt.Printf("\t-------------------------------- User #%v --------------------------------\n", uId)

		fmt.Printf("\t\tFirst Name: %v\n", u.First)
		fmt.Printf("\t\tLast Name: %v\n", u.Last)
		fmt.Printf("\t\tAge: %v\n", u.Age)
		fmt.Printf("\t\tSayings: \n")

		for j, saying := range u.Sayings {
			fmt.Printf("\t\t\tSaying #%v - \"%v\"\n", j+1, saying)
		}

		fmt.Printf("\t-------------------------------- User #%v --------------------------------\n", uId)
	}
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here

	// Sort by Age:
	fmt.Println("************************************************ Sorting by Age ***********************************************")
	sort.Sort(byAge(users))
	printUsers(users)
	fmt.Println("************************************************ Sorting by Age ***********************************************")

	fmt.Println()

	// Sort by Last:
	fmt.Println("*********************************************** Sorting by Last ***********************************************")
	sort.Sort(byLast(users))
	printUsers(users)
	fmt.Println("*********************************************** Sorting by Last ***********************************************")

	fmt.Println()

	// Sort each user's Sayings:
	fmt.Println("******************************** Sorting by Last - Sorting each User's Sayings ********************************")
	for _, u := range users {
		sort.Strings(u.Sayings)
	}
	printUsers(users)
	fmt.Println("******************************** Sorting by Last - Sorting each User's Sayings ********************************")
}
