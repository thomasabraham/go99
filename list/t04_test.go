package list

import (
	"testing"
)

func Test04(t *testing.T) {

	cases := []struct {
		input  []int
		output int
	}{
		{[]int{1, 2, 3, 4}, 4},
		{[]int{55}, 1},
		{[]int{123, 654, 66, 78, 43}, 5},
		{[]int{}, 0},
	}

	for _, c := range cases {
		count := Count(c.input)
		if count != c.output {
			t.Errorf("Result %d is not equal to expected value %d", count, c.output)
		}
	}
}
