// Benchmarking, Examples and Testing of package mymath.
package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	xi := []int{7, 2, 3, 5, 4, 6, 1}
	expected := 20.0 / 5.0

	gotten := CenteredAvg(xi)

	if gotten != expected {
		t.Errorf("Expected %v got %v", expected, gotten)
	}
}

func TestTableCenteredAvg(t *testing.T) {
	type test struct {
		data   []int
		result float64
	}

	tests := []test{
		{[]int{7, 2, 3, 5, 4, 6, 1}, 20.0 / 5.0},
		{[]int{8, 8, 5, 6, 2, 1, 3}, 24.0 / 5.0},
		{[]int{5, 4, 2, 3, 9, 4, 7}, 23.0 / 5.0},
		{[]int{3, 3, 4, 4, 2, 9, 7}, 21.0 / 5.0},
		{[]int{1, 0, 5, 6, 7, 2, 3}, 17.0 / 5.0},
		{[]int{6, 5, 6, 3, 8, 7, 2}, 27.0 / 5.0},
	}

	for _, v := range tests {
		gotten := CenteredAvg(v.data)
		if gotten != v.result {
			t.Errorf("Expected %v got %v", v.result, gotten)
		}
	}
}

func ExampleCenteredAvg() {
	xi := []int{7, 2, 3, 5, 4, 6, 1}
	result := CenteredAvg(xi)

	fmt.Println(result)
	// Output:
	// 4
}

func BenchmarkCenteredAvg(b *testing.B) {
	xi := []int{7, 2, 3, 5, 4, 6, 1}

	for i := 0; i < b.N; i++ {
		CenteredAvg(xi)
	}
}
