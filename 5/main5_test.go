package main

import (
	"reflect"
	"testing"
)

func TestFindIntersection(t *testing.T) {
	tests := []struct {
		a          []int
		b          []int
		expectedOk bool
		expected   []int
	}{
		{
			[]int{65, 3, 58, 678, 64},
			[]int{64, 2, 3, 43},
			true,
			[]int{64, 3},
		},
		{
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			false,
			[]int{},
		},
		{
			[]int{10, 20, 30},
			[]int{30, 40, 50},
			true,
			[]int{30},
		},
		{
			[]int{},
			[]int{1, 2, 3},
			false,
			[]int{},
		},
		{
			[]int{5, 10, 15},
			[]int{},
			false,
			[]int{},
		},
	}

	for _, test := range tests {
		ok, result := findIntersection(test.a, test.b)
		if ok != test.expectedOk || !reflect.DeepEqual(result, test.expected) {
			t.Errorf("findIntersection(%v, %v) = (%v, %v); expected (%v, %v)",
				test.a, test.b, ok, result, test.expectedOk, test.expected)
		}
	}
}
