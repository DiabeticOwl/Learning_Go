package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

// Go doesn't have inheritance of structs but it has a
// composition that relates one's atributes to another.
// This is called Embedding.
type Animal struct {
	// Tags are like descriptions
	// strings related to an atribute.
	Name   string `required,max:"100"`
	Origin string
}

type Bird struct {
	// This reads as "Bird has Animal caracteristics"
	// and not as "Bird is Animal".
	Animal
	SpeedKPH float32
	CanFly   bool
}

func main() {
	fmt.Println("----- Maps -----")

	// Alternative declaration for an empy map:
	// statePopulations := make(map[string]int{})
	statePopulations := map[string]int{
		"California":   39237836,
		"Texas":        29527941,
		"Florida":      21781128,
		"New York":     19835913,
		"Pennsylvania": 12964056,
		"Illinois":     12671469,
		"Ohio":         11780017,
	}

	statePopulations["Georgia"] = 10799566

	delete(statePopulations, "Georgia")

	fmt.Println(statePopulations["Texas"])
	fmt.Println(statePopulations)
	fmt.Println(len(statePopulations))

	fmt.Println()

	wrongState, ok := statePopulations["ejiweo"]

	fmt.Println(wrongState, ok)

	fmt.Println()

	// As slices, referring to a map with another variable
	// with point to the initial one.
	sp := statePopulations
	delete(sp, "Texas")
	fmt.Println(sp)
	fmt.Println(statePopulations)

	fmt.Println("----- Maps -----")

	fmt.Println()

	fmt.Println("----- Structs -----")

	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.companions[2])

	fmt.Println()

	fmt.Println("----- Structs | Anonymous Structs -----")

	// This anonymous strutcs are recommended to be used
	// as a quick and short-lived variable inside a script.
	// Creating another anonymous struct based on an initial
	// one won't point to the original in memory.
	aDoctor2 := struct{ name string }{name: "John Pertwee"}

	anotherDoctor := aDoctor2
	anotherDoctor.name = "Tom baker"

	fmt.Println(aDoctor2)
	fmt.Println(anotherDoctor)

	// Like arrays, for pointing to one specific struct we
	// need to use the "&" symbol.
	anotherDoctor2 := &aDoctor2
	anotherDoctor2.name = "David Tennant"

	fmt.Println(anotherDoctor2)
	fmt.Println(aDoctor2)

	fmt.Println("----- Structs | Anonymous Structs -----")
	fmt.Println("----- Structs -----")

	fmt.Println()

	fmt.Println("----- Embedding -----")

	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false

	fmt.Println(b.Name)

	fmt.Println("----- Embedding -----")
	fmt.Println("----- Embedding -----")

	fmt.Println()

	fmt.Println("----- Tags -----")

	// To extract a tag of a struct we need to
	// use the reflect package.
	// A tag is just a string.
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")

	fmt.Printf("%v, %T\n", field.Tag, field.Tag)

	fmt.Println("----- Tags -----")
}
