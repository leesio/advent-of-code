package main

import (
	"testing"
)

var (
	testInput = []string{
		"L.LL.LL.LL",
		"LLLLLLL.LL",
		"L.L.L..L..",
		"LLLL.LL.LL",
		"L.LL.LL.LL",
		"L.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLLL",
		"L.LLLLLL.L",
		"L.LLLLL.LL",
	}
	testVisible = []string{
		".......#.",
		"...#.....",
		".#.......",
		".........",
		"..#L....#",
		"....#....",
		".........",
		"#........",
		"...#.....",
	}
)

func TestMakeGrid(t *testing.T) {
	grid := makeGrid(testInput)
	if len(grid) != len(testInput) {
		t.Errorf("Grid has size: %d, expected: %d", len(grid), len(testInput))
	}
	if occupied := grid.countOccupiedSeats(); occupied != 0 {
		t.Errorf("Grid has: %d occupied seats, expected: %d", occupied, 0)
	}
}

func TestGridVisible(t *testing.T) {
	g := makeGrid(testVisible)
	if visible := g.visibleOccupiedSeats(4, 3); visible != 8 {
		t.Errorf("Got %d visible, expected: %d", visible, 8)
	}
}

func TestPartOne(t *testing.T) {
	if res := PartOne(testInput); res != 37 {
		t.Errorf("Got %d for part two, expected: %d", res, 37)
	}
}
func TestPartTwo(t *testing.T) {
	if res := PartTwo(testInput); res != 26 {
		t.Errorf("Got %d for part two, expected: %d", res, 26)
	}
}
