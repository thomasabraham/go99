package arithmetic

import (
	"testing"
)

func Test01(t *testing.T) {
	cases := []struct {
		input  int
		output bool
	}{
		{-1, false},
		{-10, false},
		{-13, false},
		{0, false},
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
		{10, false},
		{11, true},
		{12, false},
		{13, true},
		{14, false},
		{15, false},
		{16, false},
	}

	for _, c := range cases {
		result := PrimeOrNot(c.input)
		if result != c.output {
			t.Errorf("Result %v is not equal to expected output %v for input %v", result, c.output, c.input)
		}
	}
}
