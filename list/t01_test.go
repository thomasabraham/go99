package list

import (
	"testing"
)

func Test01(t *testing.T) {

	cases := []struct {
		input  []int
		output int
	}{
		{[]int{1, 2, 3, 4}, 4},
		{[]int{55}, 55},
	}

	for _, c := range cases {
		lastNumber := Last(c.input)
		if lastNumber != c.output {
			t.Errorf("Last number %d is not equal to expected value %d", lastNumber, c.output)
		}
	}
}
