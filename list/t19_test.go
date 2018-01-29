package list

import (
	"reflect"
	"testing"
)

func Test19(t *testing.T) {
	cases := []struct {
		input  []int
		N      int
		output []int
	}{
		{[]int{}, 0, []int{}},
		{[]int{}, 10, []int{}},
		{[]int{1}, 0, []int{1}},
		{[]int{1}, 10, []int{1}},
		{[]int{1, 2}, 0, []int{1, 2}},
		{[]int{1, 2}, 1, []int{2, 1}},
		{[]int{1, 2}, 2, []int{1, 2}},
		{[]int{1, 2}, 3, []int{2, 1}},
		{[]int{1, 2, 3}, 0, []int{1, 2, 3}},
		{[]int{1, 2, 3}, 1, []int{2, 3, 1}},
		{[]int{1, 2, 3}, 2, []int{3, 1, 2}},
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}},
		{[]int{1, 2, 3}, 4, []int{2, 3, 1}},
	}
	for _, c := range cases {
		result := RotateListLeft(c.input, c.N)
		if !reflect.DeepEqual(result, c.output) {
			t.Errorf("Result %v is not equal to the expected output %v", result, c.output)
		}
	}
}
