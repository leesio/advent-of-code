package main

import (
	"testing"
)

var testInput = []string{
	"class: 1-3 or 5-7",
	"row: 6-11 or 33-44",
	"seat: 13-40 or 45-50",
	"",
	"your ticket:",
	"7,1,14",
	"",
	"nearby tickets:",
	"7,3,47",
	"40,4,50",
	"55,2,20",
	"38,6,12",
}

func TestPartOne(t *testing.T) {
	if res := PartOne(testInput); res != 71 {
		t.Errorf("part one, received: %d, expected: %d", res, 71)
	}
}
func TestPartTwo(t *testing.T) {
	if res := PartTwo(testInput); res != 71 {
		t.Errorf("part one, received: %d, expected: %d", res, 71)
	}
}
