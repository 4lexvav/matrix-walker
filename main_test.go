package main

import (
	"testing"
)

func TestFlatten(t *testing.T) {
	var tests = []struct {
		matrix [][]int32
		want   []int32
		result bool
	}{
		{
			[][]int32{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			[]int32{1, 2, 3, 6, 9, 8, 7, 4, 5},
			true,
		},
		{
			[][]int32{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			[]int32{1, 234, 3, 6, 9, 8, 7, 4, 5},
			false,
		},
	}

	for _, test := range tests {
		result := true
		got := flatten(test.matrix)
		for i := range got {
			if got[i] != test.want[i] {
				result = false
			}
		}

		if result != test.result {
			t.Errorf("flatten(%v) == %v, expected %v, got %v", test.matrix, test.want, test.result, result)
		}
	}
}
