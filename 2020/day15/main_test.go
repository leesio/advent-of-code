package main

import (
	"testing"
)

var testCases = []struct {
	input []int
	exp   int
}{
	// {[]int{0, 3, 6}, 0},
	{[]int{1, 3, 2}, 1},
	{[]int{2, 1, 3}, 10},
	{[]int{1, 2, 3}, 27},
	{[]int{2, 3, 1}, 78},
	{[]int{3, 2, 1}, 438},
	{[]int{3, 1, 2}, 1836},
}

func TestGetNthNumberSpoken(t *testing.T) {
	n := 2020
	for i, tc := range testCases {
		if res := getNthNumberSpoken(tc.input, n); res != tc.exp {
			t.Errorf("Case: %d - Got %d, expected: %d", i, res, tc.exp)
		}
	}
}
