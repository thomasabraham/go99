package list

import (
	"reflect"
	"testing"
)

func Test11(t *testing.T) {

	cases := []struct {
		input  []int
		output []interface{}
	}{
		{[]int{}, []interface{}{}},
		{[]int{1}, []interface{}{1}},
		{[]int{1, 1}, []interface{}{[]int{2, 1}}},
		{[]int{1, 2, 2, 1}, []interface{}{1, []int{2, 2}, 1}},
		{[]int{1, 1, 2, 2, 2, 1}, []interface{}{[]int{2, 1}, []int{3, 2}, 1}},
		{[]int{1, 1, 2, 2, 2, 1, 3, 3, 2, 1}, []interface{}{[]int{2, 1}, []int{3, 2}, 1, []int{2, 3}, 2, 1}},
	}

	for _, c := range cases {
		result := RunlengthEncode2(c.input)
		if !reflect.DeepEqual(result, c.output) {
			t.Errorf("Result %v dont match the expected output %v", result, c.output)
		}
	}
}
