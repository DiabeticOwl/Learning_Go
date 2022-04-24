// Pacakage dog provides a type and functionality related to dogs in real life.
package dog

// Dog is a simple struct that defines a name and age of a dog.
type Dog struct {
	Name string
	// yearsOld is the amount of humans years that the dog has.
	YearsOld int
}

// Years returns the amout of dog years that the attached Dog has.
func (d Dog) Years() int {
	return d.YearsOld * 7
}
