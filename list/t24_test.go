package list

import (
	"testing"
)

func Test24(t *testing.T) {
	cases := []struct {
		start     int
		end       int
		count     int
		resultLen int
	}{
		{0, 10, 5, 5},
		{-10, 0, 5, 5},
		{0, 10, -5, 0},
		{-10, 0, -5, 0},
		{11, 20, 10, 10},
		{20, 10, 5, 0},
		{-10, -20, 5, 0},
		{-20, -10, 5, 5},
	}

	for _, c := range cases {
		result := RandomInRange(c.start, c.end, c.count)
		if len(result) != c.resultLen {
			t.Errorf("Length or result %v is not equal to expected value %v", len(result), c.resultLen)
		}
		for _, v := range result {
			if v < c.start || v > c.end {
				t.Errorf("Value %v in result is outside required range [%v, %v]", v, c.start, c.end)
			}
		}
	}
}
