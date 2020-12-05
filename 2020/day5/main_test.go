package main

import "testing"

var rawInput = []string{
	"FBFBBFFRLR",
	"BFFFBBFRRR",
	"FFFBBBFRRR",
	"BBFFBBFRLL",
}
var results = []struct {
	row    int
	column int
	seatID int
}{
	{44, 5, 357},
	{70, 7, 567},
	{14, 7, 119},
	{102, 4, 820},
}

var validChars = map[string]bool{"F": true, "B": true, "R": true, "L": true}

func TestParseInput(t *testing.T) {
	input := ParseInput(rawInput)
	if len(input) != len(rawInput) {
		t.Errorf("Got input with len: %d, expected: %d", len(input), len(rawInput))
	}
	for _, l := range input {
		if len(l) != 10 {
			t.Errorf("Got boarding pass with len: %d, expected: %d", len(l), 7)
		}
		for _, ch := range l {
			if !validChars[ch] {
				t.Errorf("Unexpected char: %s", ch)
			}
		}
	}
}
func TestSeatIDCalculation(t *testing.T) {
	input := ParseInput(rawInput)
	if len(input) != len(rawInput) {
		t.Errorf("Got input with len: %d, expected: %d", len(input), len(rawInput))
	}
	for n, l := range input {
		result := results[n]
		col, row := GetColNum(l), GetRowNum(l)
		if row != result.row {
			t.Errorf("Got row %d, expected: %d (case: %d)", row, result.row, n)
		}
		if col != result.column {
			t.Errorf("Got column %d, expected: %d (case: %d)", col, result.column, n)
		}
		if seatID := GetSeatID(row, col); seatID != result.seatID {
			t.Errorf("Got seatID %d, expected: %d (case: %d)", seatID, result.seatID, n)
		}
	}
}
