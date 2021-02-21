package main

import (
	"fmt"
	"sort"

	"github.com/leesio/advent-of-code/2020/util"
)

func main() {
	rawInput, err := util.GetInput("input")
	if err != nil {
		panic(err)
	}
	max := 0
	rows := make([]sort.IntSlice, 129)
	for _, pass := range rawInput {
		row, col := GetRowNum(pass), GetColNum(pass)
		seatID := GetSeatID(row, col)
		if seatID > max {
			max = seatID
		}
		rows[row] = append(rows[row], seatID)
	}
	fmt.Printf("part 1: %d\n", max)

	var myID int
	for _, row := range rows {
		if len(row) != 8 && len(row) != 0 {
			row.Sort()
			prev := row[0]
			for _, seatID := range row[1:] {
				if seatID-prev != 1 {
					myID = seatID - 1
					break
				}
				prev = seatID
			}
		}
	}
	fmt.Printf("part 2: %d\n", myID)
}

func BinarySearch(input string, upperMarker string, lowerMarker string) int {
	maxIndex, val := len(input)-1, 0
	for n, x := range input {
		bit := 0
		if string(x) == upperMarker {
			bit = 1
		}
		val = val | (bit << (maxIndex - n))
	}
	return val
}

func GetRowNum(input string) int {
	return BinarySearch(input[:7], "B", "F")
}
func GetColNum(input string) int {
	return BinarySearch(input[7:], "R", "L")
}

func GetSeatID(row, col int) int {
	return (row * 8) + col
}
