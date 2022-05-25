// Hands-on exercise #3
// 	● Create a new type: vehicle.
// 		○ The underlying type is a struct.
// 		○ The fields:
// 			■ doors
// 			■ color
// 	● Create two new types: truck & sedan.
// 		○ The underlying type of each of these new types is a struct.
// 		○ Embed the “vehicle” type in both truck & sedan.
// 		○ Give truck the field “fourWheel” which will be set to bool.
// 		○ Give sedan the field “luxury” which will be set to bool.
// 	● Using the vehicle, truck, and sedan structs:
// 		○ using a composite literal, create a value of type truck and assign values to the
// 		  fields;
// 		○ using a composite literal, create a value of type sedan and assign values to the
// 		  fields.
// 	● Print out each of these values.
// 	● Print out a single field from each of these values.

package main

import "fmt"

type vehicle struct {
	doors uint8
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	mytruck := truck{
		// Since we declare the "fourWheel" field by composite
		// literal it is required to declare the field "vehicle"
		// the same way.
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		fourWheel: true,
	}

	mySedan := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "white",
		},
		luxury: false,
	}

	fmt.Println(mytruck)
	fmt.Println(mySedan)

	fmt.Println(mytruck.doors)
	fmt.Println(mySedan.luxury)
}
