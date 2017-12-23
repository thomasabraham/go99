package list

import (
	"reflect"
	"testing"
)

func Test09(t *testing.T) {

	cases := []struct {
		input  []int
		output [][]int
	}{
		{[]int{}, [][]int{}},
		{[]int{1}, [][]int{{1}}},
		{[]int{1, 1}, [][]int{{1, 1}}},
		{[]int{1, 2, 2, 1}, [][]int{{1}, {2, 2}, {1}}},
		{[]int{1, 1, 2, 2, 2, 1}, [][]int{{1, 1}, {2, 2, 2}, {1}}},
		{[]int{1, 1, 2, 2, 2, 1, 3, 3, 2, 1}, [][]int{{1, 1}, {2, 2, 2}, {1}, {3, 3}, {2}, {1}}},
	}

	for _, c := range cases {
		result := PackDuplicates(c.input)
		if !reflect.DeepEqual(result, c.output) {
			t.Errorf("Result %v dont match the expected output %v", result, c.output)
		}
	}
}
