package list

import (
	"testing"
)

func Test16(t *testing.T) {
	cases := []struct {
		input  []int
		N      int
		output []int
	}{
		{[]int{}, 2, []int{}},
		{[]int{}, 10, []int{}},
		{[]int{1}, 2, []int{1}},
		{[]int{1, 2}, 2, []int{1}},
		{[]int{1, 2, 3}, 2, []int{1, 3}},
		{[]int{1, 2, 3}, 3, []int{1, 2}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2, []int{1, 3, 5, 7, 9}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3, []int{1, 2, 4, 5, 7, 8}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5, []int{1, 2, 3, 4, 6, 7, 8, 9}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9, []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 10, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}

	for _, c := range cases {
		result := DropEveryNth(c.input, c.N)
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
