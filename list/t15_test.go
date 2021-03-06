package list

import (
	"testing"
)

func Test15(t *testing.T) {

	cases := []struct {
		input  []int
		N      int
		output []int
	}{
		{[]int{}, 2, []int{}},
		{[]int{1}, 4, []int{1, 1, 1, 1}},
		{[]int{1, 1}, 1, []int{1, 1}},
		{[]int{1, 2, 2, 1}, 3, []int{1, 1, 1, 2, 2, 2, 2, 2, 2, 1, 1, 1}},
		{[]int{1, 1, 2, 2, 2, 1}, 2, []int{1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 1, 1}},
		{[]int{1, 1, 2, 2, 2, 1, 3, 3, 2, 1}, 2, []int{1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 1, 1, 3, 3, 3, 3, 2, 2, 1, 1}},
	}

	for _, c := range cases {
		result := DuplicateElementsN(c.input, c.N)
		if len(result) != len(c.output) {
			t.Errorf("Length of result %v don't match with expected output %v", result, c.output)
		}
		for k, v := range c.output {
			if v != result[k] {
				t.Errorf("Result %v is not equal to expected value %v", result, c.output)
			}
		}
	}
}
