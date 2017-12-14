package list

import (
	"testing"
)

func Test06(t *testing.T) {

	cases := []struct {
		input  []int
		output bool
	}{
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 2}, false},
		{[]int{55}, true},
		{[]int{1, 55, 1}, true},
		{[]int{32, 1, 55, 1, 32}, true},
		{[]int{}, true},
	}

	for _, c := range cases {
		isPalindRome := IsPalindRome(c.input)
		if isPalindRome != c.output {
			t.Errorf("Result %t is not equal to expected value %t", isPalindRome, c.output)
		}
	}
}
