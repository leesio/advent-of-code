package main

import (
	"testing"
)

var testInput = []int{
	35,
	20,
	15,
	25,
	47,
	40,
	62,
	55,
	65,
	95,
	102,
	117,
	150,
	182,
	127,
	219,
	299,
	277,
	309,
	576,
}

func TestPartOne(t *testing.T) {
	if res := PartOne(testInput, 5); res != 127 {
		t.Errorf("Received: %d, expected: %d", res, 127)
	}
}

func TestPartOneAlternative(t *testing.T) {
	if res := PartOneAlternative(testInput, 5); res != 127 {
		t.Errorf("Received: %d, expected: %d", res, 127)
	}
}

func TestPartTwo(t *testing.T) {
	target := 127
	if res := PartTwo(testInput, target); res != 62 {
		t.Errorf("Received: %d, expected: %d", res, 62)
	}
}
