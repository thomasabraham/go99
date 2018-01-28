package list

import (
	"testing"
)

func Test17(t *testing.T) {
	cases := []struct {
		input []int
		pivot int
		left  []int
		right []int
	}{
		{[]int{}, -1, []int{}, []int{}},
		{[]int{}, 0, []int{}, []int{}},
		{[]int{}, 1, []int{}, []int{}},
		{[]int{}, 2, []int{}, []int{}},
		{[]int{1}, -1, []int{}, []int{1}},
		{[]int{1}, 0, []int{}, []int{1}},
		{[]int{1}, 1, []int{1}, []int{}},
		{[]int{1}, 2, []int{1}, []int{}},
		{[]int{1, 2}, -1, []int{}, []int{1, 2}},
		{[]int{1, 2}, 0, []int{}, []int{1, 2}},
		{[]int{1, 2}, 1, []int{1}, []int{2}},
		{[]int{1, 2}, 2, []int{1, 2}, []int{}},
		{[]int{1, 2}, 3, []int{1, 2}, []int{}},
		{[]int{1, 2}, 5, []int{1, 2}, []int{}},
		{[]int{1, 2, 3}, -1, []int{}, []int{1, 2, 3}},
		{[]int{1, 2, 3}, 0, []int{}, []int{1, 2, 3}},
		{[]int{1, 2, 3}, 1, []int{1}, []int{2, 3}},
		{[]int{1, 2, 3}, 2, []int{1, 2}, []int{3}},
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}, []int{}},
		{[]int{1, 2, 3}, 4, []int{1, 2, 3}, []int{}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 4, []int{1, 2, 3, 4}, []int{5, 6, 7, 8, 9}},
	}

	for _, c := range cases {
		left, right := SplitList(c.input, c.pivot)
		if len(left) != len(c.left) {
			t.Errorf("Length of left %v don't match with expected output %v", left, c.left)
		}
		if len(right) != len(c.right) {
			t.Errorf("Length of right %v don't match with expected output %v", right, c.right)
		}
		for k, v := range c.left {
			if v != left[k] {
				t.Errorf("Left %v is not equal to expected value %v", left, c.left)
			}
		}
		for k, v := range c.right {
			if v != right[k] {
				t.Errorf("Right %v is not equal to expected value %v", right, c.right)
			}
		}
	}
}
