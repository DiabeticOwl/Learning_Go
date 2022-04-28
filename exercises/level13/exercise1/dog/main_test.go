// Testing and examples for dog package.
package dog

import (
	"fmt"
	"testing"
)

// TestYears will check that the number returned by the Years
// function will return the actual dog age of a dog.
func TestYears(t *testing.T) {
	n := Years(5)

	if n != 35 {
		t.Errorf("Expected 35. Got %v.", n)
	}
}

// TestYearsTwo will check that the number returned by the Years
// function will return the actual dog age of a dog.
func TestYearsTwo(t *testing.T) {
	n := YearsTwo(5)

	if n != 35 {
		t.Errorf("Expected 35. Got %v.", n)
	}
}

func ExampleYears() {
	fmt.Println(Years(5))
	// output:
	// 35
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(5))
	// output:
	// 35
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(5)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(5)
	}
}
