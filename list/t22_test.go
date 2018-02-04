package list

import (
	"reflect"
	"testing"
)

func Test22(t *testing.T) {
	cases := []struct {
		start  int
		end    int
		result []int
	}{
		{1, -1, []int{}},
		{3, 2, []int{}},
		{4, 4, []int{4}},
		{-3, 3, []int{-3, -2, -1, 0, 1, 2, 3}},
		{-33, -28, []int{-33, -32, -31, -30, -29, -28}},
		{22, 26, []int{22, 23, 24, 25, 26}},
	}

	for _, c := range cases {
		result := RangeAsList(c.start, c.end)
		if !reflect.DeepEqual(c.result, result) {
			t.Errorf("Result %v is not expected result %v", result, c.result)
		}
	}
}
