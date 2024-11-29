package main

import (
	"reflect"
	"testing"
)

func TestDifference(t *testing.T) {
	tests := []struct {
		slice1   []string
		slice2   []string
		expected []string
	}{
		{
			[]string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			[]string{"banana", "date", "fig"},
			[]string{"apple", "cherry", "43", "lead", "gno1"},
		},
		{
			[]string{"a", "b", "c"},
			[]string{"a", "b"},
			[]string{"c"},
		},
		{
			[]string{"x", "y", "z"},
			[]string{"a", "b", "c"},
			[]string{"x", "y", "z"},
		},
		{
			[]string{"common", "unique"},
			[]string{"common"},
			[]string{"unique"},
		},
		{
			[]string{""},
			[]string{"a", "b"},
			[]string{""},
		},
	}

	for _, test := range tests {
		result := difference(test.slice1, test.slice2)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("difference(%v, %v) = %v; expected %v", test.slice1, test.slice2, result, test.expected)
		}
	}
}
