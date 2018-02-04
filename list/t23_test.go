package list

import (
	"testing"
)

func Test23(t *testing.T) {

	// Test for panic if length is greater than size of array
	if true != testRandomNPanic([]int{1, 2, 3}, 4) {
		t.Errorf("Did not panic when length > size of input")
	}

	// Test for panic if length is negative
	if true != testRandomNPanic([]int{1, 2, 3}, -1) {
		t.Errorf("Did not panic when length < 0")
	}

	input := []int{3, 2, 66, 4, 8, 43, 56, 67, 34, 33, 22, 34, 8, 93, 73}
	length := len(input) / 2
	result := RandomN(input, length)
	if len(result) != length {
		t.Errorf("Result length did not match what is expected")
	}
	if !testContains(result, input) {
		t.Errorf("Result is not contained in input")
	}

	input = []int{3, 2, 66, 4, 8, 43, 56, 67, 34, 33, 22, 34, 8, 93, 73}
	length = len(input)
	result = RandomN(input, length)
	if len(result) != length {
		t.Errorf("Result length did not match what is expected")
	}
	if !testContains(result, input) {
		t.Errorf("Result is not contained in input")
	}

}

func testContains(input []int, container []int) bool {
	for _, v := range input {
		index := findIndexInList(container, v)
		if index < 0 {
			return false
		}
		container[index] = -1
	}
	return true
}

func findIndexInList(input []int, value int) int {
	for k, v := range input {
		if v == value {
			return k
		}
	}
	return -1
}

func testRandomNPanic(input []int, length int) (didPanic bool) {
	defer func() {
		if r := recover(); r != nil {
			didPanic = true
		}
	}()
	RandomN(input, length)
	return false
}
