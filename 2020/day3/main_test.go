package main

import (
	"testing"
)

var (
	testInput = []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}

	testInputParsed = [][]string{
		{".", ".", "#", "#", ".", ".", ".", ".", ".", ".", "."},
		{"#", ".", ".", ".", "#", ".", ".", ".", "#", ".", "."},
		{".", "#", ".", ".", ".", ".", "#", ".", ".", "#", "."},
		{".", ".", "#", ".", "#", ".", ".", ".", "#", ".", "#"},
		{".", "#", ".", ".", ".", "#", "#", ".", ".", "#", "."},
		{".", ".", "#", ".", "#", "#", ".", ".", ".", ".", "."},
		{".", "#", ".", "#", ".", "#", ".", ".", ".", ".", "#"},
		{".", "#", ".", ".", ".", ".", ".", ".", ".", ".", "#"},
		{"#", ".", "#", "#", ".", ".", ".", "#", ".", ".", "."},
		{"#", ".", ".", ".", "#", "#", ".", ".", ".", ".", "#"},
		{".", "#", ".", ".", "#", ".", ".", ".", "#", ".", "#"},
	}
)

func TestInputParsing(t *testing.T) {
	parsed := ParseInput(testInput)
	if len(parsed) != len(testInputParsed) {
		t.Errorf("Received input with length: %d, expected: %d", len(parsed), len(testInputParsed))
	}
	for i := 0; i < len(parsed); i++ {
		exp, act := testInputParsed[i], parsed[i]
		if len(exp) != len(act) {
			t.Errorf("Row %d input has length: %d, expected: %d", i, len(act), len(exp))
		}
		for j := 0; j < len(exp); j++ {
			if exp[j] != act[j] {
				t.Errorf("Row: %d Col: %d is: %s, expected: %s", i, j, act[j], exp[j])
			}
		}
	}
}
func TestPartOne(t *testing.T) {
	if n := PartOne(testInputParsed); n != 7 {
		t.Errorf("Received: %d, expected 7", n)
	}
}
func TestPartTwo(t *testing.T) {
	if n := PartTwo(testInputParsed); n != 336 {
		t.Errorf("Received: %d, expected 336", n)
	}
}
