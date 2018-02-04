package list

import (
	"reflect"
	"testing"
)

func Test21(t *testing.T) {
	cases := []struct {
		inputList []int
		value     int
		index     int
		result    []int
	}{
		{[]int{}, 25, -10, []int{}},
		{[]int{}, 0, -1, []int{}},
		{[]int{}, -29, 0, []int{-29}},
		{[]int{}, 33, 1, []int{}},
		{[]int{}, 13, 10, []int{}},
		{[]int{33}, 5, -1, []int{33}},
		{[]int{33}, 4, 0, []int{4, 33}},
		{[]int{33}, 3, 1, []int{33, 3}},
		{[]int{33}, 2, 2, []int{33}},
		{[]int{33, 2}, 5, -1, []int{33, 2}},
		{[]int{33, 2}, 4, 0, []int{4, 33, 2}},
		{[]int{33, 2}, 3, 1, []int{33, 3, 2}},
		{[]int{33, 2}, 2, 2, []int{33, 2, 2}},
		{[]int{33, 2}, 3, 2, []int{33, 2, 3}},
		{[]int{33, 2}, 3, 4, []int{33, 2}},
	}

	for _, c := range cases {
		result := ListInsert(c.inputList, c.value, c.index)
		if !reflect.DeepEqual(c.result, result) {
			t.Errorf("Result %v is not the expected result %v", result, c.result)
		}
	}

}
