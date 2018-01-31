package list

import (
	"reflect"
	"testing"
)

func Test20(t *testing.T) {
	cases := []struct {
		input  []int
		K      int
		value  int
		output []int
		panic  bool
	}{
		{[]int{}, -1, 0, []int{}, true},
		{[]int{}, 0, 0, []int{}, true},
		{[]int{}, 1, 0, []int{}, true},
		{[]int{3}, -1, 0, []int{}, true},
		{[]int{3}, 0, 3, []int{}, false},
		{[]int{3}, 1, 0, []int{}, true},
		{[]int{5, 4}, -2, 0, []int{}, true},
		{[]int{5, 4}, -1, 0, []int{}, true},
		{[]int{5, 4}, 0, 5, []int{4}, false},
		{[]int{5, 4}, 1, 4, []int{5}, false},
		{[]int{5, 4}, 2, 0, []int{}, true},
		{[]int{5, 4}, 3, 0, []int{}, true},
		{[]int{4, 3, 5}, -1, 0, []int{}, true},
		{[]int{4, 3, 5}, 0, 4, []int{3, 5}, false},
		{[]int{4, 3, 5}, 1, 3, []int{4, 5}, false},
		{[]int{4, 3, 5}, 2, 5, []int{4, 3}, false},
		{[]int{4, 3, 5}, 3, 0, []int{}, true},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 3, 4, []int{1, 2, 3, 5, 6, 7, 8}, false},
	}

	for _, c := range cases {
		if c.panic {
			if !testPanic(c.input, c.K) {
				t.Errorf("Test case %v didn't result in panic", c)
			}
		} else {
			value, result := RemoveKthFromList(c.input, c.K)
			if !reflect.DeepEqual(result, c.output) {
				t.Errorf("Result %v is not expected output %v", result, c.output)
			}
			if value != c.value {
				t.Errorf("Result %v is not expected output for case %v", value, c)
			}
		}
	}
}

func testPanic(input []int, K int) (didPanic bool) {
	defer func() {
		if r := recover(); r != nil {
			didPanic = true
		}
	}()
	RemoveKthFromList(input, K)
	return false
}
