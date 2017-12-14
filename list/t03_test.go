package list

import (
	"testing"
)

func Test03(t *testing.T) {

	cases := []struct {
		input  []int
		index  int
		output int
	}{
		{[]int{1, 2, 3, 4}, 2, 3},
		{[]int{55, 23}, 0, 55},
	}

	for _, c := range cases {
		result := GetValueAt(c.input, c.index)
		if result != c.output {
			t.Errorf("Last result %d is not equal to expected value %d", result, c.output)
		}
	}
}
