package list

import (
	"testing"
)

func Test07(t *testing.T) {

	cases := []struct {
		input  []interface{}
		output []int
	}{
		{[]interface{}{
			[]interface{}{1, 2},
			3,
			[]interface{}{
				5,
				[]interface{}{6, 7},
			},
			8,
			9,
		}, []int{1, 2, 3, 5, 6, 7, 8, 9}},
	}

	for _, c := range cases {
		flat := Flatten(c.input)
		if len(flat) != len(c.output) {
			t.Errorf("Length of flattened slice is not expected value")
		}

		for i := range flat {
			if flat[i] != c.output[i] {

				t.Errorf("Result %v is not equal to expected value %v", flat, c.output)
			}
		}
	}
}
