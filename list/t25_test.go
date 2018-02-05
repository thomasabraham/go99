package list

import (
	"testing"
)

func Test25(t *testing.T) {
	input := []int{3, 2, 66, 4, 8, 43, 56, 67, 34, 33, 22, 34, 8, 93, 73}
	result := RandomizeList(input)
	if len(result) != len(input) {
		t.Errorf("Result length did not match what is expected")
	}
	if !testContains(result, input) {
		t.Errorf("Result is not contained in input")
	}
}
