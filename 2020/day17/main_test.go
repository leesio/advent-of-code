package main

import (
	"testing"
)

var testInput = []string{
	".#.",
	"..#",
	"###",
}

func TestPartOne(t *testing.T) {
	if res := PartOne(testInput); res != 112 {
		t.Errorf("Got %d for part one, expected: %d", res, 112)
	}
}

func TestPartTwo(t *testing.T) {
	if res := PartTwo(testInput); res != 848 {
		t.Errorf("Got %d for part two, expected: %d", res, 848)
	}
}
