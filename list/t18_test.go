package list

import (
	"reflect"
	"testing"
)

func Test18(t *testing.T) {
	cases := []struct {
		input  []int
		start  int
		end    int
		output []int
	}{
		{[]int{}, -1, -10, []int{}},
		{[]int{1}, -1, -10, []int{}},
		{[]int{1}, 0, 0, []int{}},
		{[]int{1}, 2, 2, []int{}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 2, 4, []int{2, 3, 4}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 4, 2, []int{2, 3, 4}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, -4, 3, []int{1, 2, 3}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 4, 13, []int{4, 5, 6, 7, 8}},
	}

	for _, c := range cases {
		result := SliceList(c.input, c.start, c.end)
		if !reflect.DeepEqual(result, c.output) {
			t.Errorf("Result %v dont match the expected output for case %v", result, c)
		}
	}
}
