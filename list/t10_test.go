package list

import (
	"reflect"
	"testing"
)

func Test10(t *testing.T) {

	cases := []struct {
		input  []int
		output [][]int
	}{
		{[]int{}, [][]int{}},
		{[]int{1}, [][]int{{1, 1}}},
		{[]int{1, 1}, [][]int{{2, 1}}},
		{[]int{1, 2, 2, 1}, [][]int{{1, 1}, {2, 2}, {1, 1}}},
		{[]int{1, 1, 2, 2, 2, 1}, [][]int{{2, 1}, {3, 2}, {1, 1}}},
		{[]int{1, 1, 2, 2, 2, 1, 3, 3, 2, 1}, [][]int{{2, 1}, {3, 2}, {1, 1}, {2, 3}, {1, 2}, {1, 1}}},
	}

	for _, c := range cases {
		result := RunlengthEncode(c.input)
		if !reflect.DeepEqual(result, c.output) {
			t.Errorf("Result %v dont match the expected output %v", result, c.output)
		}
	}
}
