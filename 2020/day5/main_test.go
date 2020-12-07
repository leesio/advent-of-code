package main

import (
	"testing"
)

var testCases = []struct {
	input  string
	row    int
	column int
	seatID int
}{
	{"FBFBBFFRLR", 44, 5, 357},
	{"BFFFBBFRRR", 70, 7, 567},
	{"FFFBBBFRRR", 14, 7, 119},
	{"BBFFBBFRLL", 102, 4, 820},
}

var validChars = map[string]bool{"F": true, "B": true, "R": true, "L": true}

func TestSeatIDCalculation(t *testing.T) {
	for n, testCase := range testCases {
		col, row := GetColNum(testCase.input), GetRowNum(testCase.input)
		if row != testCase.row {
			t.Errorf("Got row %d, expected: %d (case: %d)", row, testCase.row, n)
		}
		if col != testCase.column {
			t.Errorf("Got column %d, expected: %d (case: %d)", col, testCase.column, n)
		}
		if seatID := GetSeatID(row, col); seatID != testCase.seatID {
			t.Errorf("Got seatID %d, expected: %d (case: %d)", seatID, testCase.seatID, n)
		}
	}
}
