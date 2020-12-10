package main

import (
	"testing"
)

var (
	testInputOne = []int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}
	testInputTwo = []int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45,
		19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3}
)

func TestPartOne(t *testing.T) {
	if res := PartOne(testInputOne); res != 35 {
		t.Errorf("Got answer of %d, expected %d", res, 35)
	}
	if res := PartOne(testInputTwo); res != 220 {
		t.Errorf("Got answer of %d, expected %d", res, 220)
	}
}

func TestPartTwo(t *testing.T) {
	if res := PartTwo(testInputOne); res != 8 {
		t.Errorf("Got answer of %d, expected %d", res, 8)
	}
	if res := PartTwo(testInputTwo); res != 19208 {
		t.Errorf("Got answer of %d, expected %d", res, 19208)
	}
}
