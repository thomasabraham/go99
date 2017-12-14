package list

import (
	"testing"
)

func Test02(t *testing.T) {

	cases := []struct {
		input  []int
		output int
	}{
		{[]int{1, 2, 3, 4}, 3},
		{[]int{55, 23}, 55},
	}

	for _, c := range cases {
		result := LastButOne(c.input)
		if result != c.output {
			t.Errorf("Last number but one %d is not equal to expected value %d", result, c.output)
		}
	}
}
