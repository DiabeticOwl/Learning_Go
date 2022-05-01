// Benchmarking, Examples and Testing of package word.
package word

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUseCount(t *testing.T) {
	s := "Hello World Hello"

	m := UseCount(s)
	expected := map[string]int{
		"Hello": 2,
		"World": 1,
	}

	if reflect.DeepEqual(m, expected) != true {
		t.Errorf("Expected %v got %v", expected, m)
	}
}

func BenchmarkUseCount(b *testing.B) {
	s := "Hello World Hello"

	for i := 0; i < b.N; i++ {
		UseCount(s)
	}
}

func TestCount(t *testing.T) {
	count := Count("Hello World Hello")

	if count != 3 {
		t.Errorf("Expected 3 got %v", count)
	}
}

func ExampleCount() {
	fmt.Println(Count("Hello this string has 8 words in it."))
	// Output:
	// 8
}

func BenchmarkCount(b *testing.B) {
	Count("Hello World Hello")
}
