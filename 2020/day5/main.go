package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/leesio/advent-of-code/2020/util"
)

const (
	rows    = 127
	columns = 7
)

func main() {
	rawInput, err := util.GetInput("input")
	if err != nil {
		panic(err)
	}
	input := ParseInput(rawInput)

	max := 0
	seatIDs := make(sort.IntSlice, len(input))
	for i, pass := range input {
		row, col := GetRowNum(pass), GetColNum(pass)
		seatID := GetSeatID(row, col)
		if seatID > max {
			max = seatID
		}
		seatIDs[i] = seatID
	}
	fmt.Printf("part 1: %d\n", max)

	seatIDs.Sort()
	prev := seatIDs[0]
	var myID int
	for _, seatID := range seatIDs[1:] {
		if seatID-prev != 1 {
			myID = seatID - 1
			break
		}
		prev = seatID
	}
	fmt.Printf("part 2: %d\n", myID)
}

func ParseInput(input []string) [][]string {
	parsed := make([][]string, len(input))
	for l, line := range input {
		parsed[l] = strings.Split(line, "")
	}
	return parsed
}

func BinarySearch(
	input []string,
	upperMarker string,
	lowerMarker string,
	max int,
) int {
	min := 0
	for _, x := range input {
		diff := (max - min) / 2
		switch x {
		case lowerMarker:
			max = min + diff
		case upperMarker:
			min = min + diff
			if max%2 == 1 {
				min++
			}
		}
	}

	if input[len(input)-1] == upperMarker {
		return max
	}
	return min
}

func GetRowNum(input []string) int {
	return BinarySearch(input[:7], "B", "F", rows)
}
func GetColNum(input []string) int {
	return BinarySearch(input[7:], "R", "L", columns)
}

func GetSeatID(row, col int) int {
	return (row * 8) + col
}
