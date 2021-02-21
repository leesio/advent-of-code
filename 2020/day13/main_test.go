package main

import (
	"fmt"
	"testing"
)

var testInput = []string{
	"939",
	"7,13,x,x,59,x,31,19",
}

func TestPartOne(t *testing.T) {
	ts, ids := parseInput(testInput)
	min := 1000000
	for _, id := range ids {
		firstAfterTs := (1 + (ts / id)) * id
		if firstAfterTs < min {
			min = firstAfterTs
		}
	}
	fmt.Println(min)
}

var cases = []struct {
	ids []string
	exp int
}{
	{testInput, 1068781},
	{[]string{"1", "17,x,13,19"}, 3417},
	{[]string{"1", "67,7,59,61"}, 754018},
	{[]string{"1", "67,x,7,59,61"}, 779210},
	{[]string{"1", "67,7,x,59,61"}, 1261476},
	{[]string{"1", "1789,37,47,1889"}, 1202161486},
}

func TestPartTwo(t *testing.T) {
	for c, ca := range cases {
		res := PartTwo(ca.ids)
		fmt.Println("---", c, res, ca.exp)
	}
}
