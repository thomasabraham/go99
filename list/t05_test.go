package list

import (
	"testing"
)

func Test05(t *testing.T) {

	cases := []struct {
		input  []int
		output []int
	}{
		{[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{[]int{55}, []int{55}},
		{[]int{}, []int{}},
	}

	for _, c := range cases {
		reversed := Reverse(c.input)
		for k, v := range c.output {
			if v != reversed[k] {
				t.Errorf("Result %d is not equal to expected value %d", reversed, c.output)
			}
		}
	}
}
