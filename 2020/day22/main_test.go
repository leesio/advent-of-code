package main

import (
	"testing"
)

var testInput = []string{
	"Player 1:",
	"9",
	"2",
	"6",
	"3",
	"1",
	"",
	"Player 2:",
	"5",
	"8",
	"4",
	"7",
	"10",
}

func TestPartOne(t *testing.T) {
	if res := PartOne(testInput); res != 306 {
		t.Errorf("Got %d for part one, expected: %d", res, 306)
	}
}
func TestPartTwo(t *testing.T) {
	if res := PartTwo(testInput); res != 291 {
		t.Errorf("Got %d for part one, expected: %d", res, 291)
	}
}
