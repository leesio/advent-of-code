package main

import (
	"fmt"
	"sort"

	"github.com/leesio/advent-of-code/2020/util"
)

const stepSize = 25

func main() {
	rawInput, err := util.GetInput("input")
	if err != nil {
		panic(err)
	}
	input, err := util.Atoi(rawInput)
	if err != nil {
		panic(err)
	}

	partOne := PartOne(input, stepSize)
	fmt.Printf("part 1: %d\n", partOne)

	partTwo := PartTwo(input, partOne)
	fmt.Printf("part 2: %d\n", partTwo)

}

func PartOne(input []int, stepSize int) int {
	for i := stepSize; i < len(input); i++ {
		if val, window := input[i], input[i-stepSize:i]; !Valid(window, val) {
			return val
		}
	}
	return -1
}
func PartTwo(input []int, target int) int {
	for i := 2; i < len(input); i++ {
		for j := 0; j+i < len(input); j++ {
			if window := input[j : j+i]; SumAll(window) == target {
				return SumLowestAndHighest(window)
			}
		}
	}
	return -1
}

func Valid(window []int, val int) bool {
	for i := 0; i < len(window); i++ {
		for j := i + 1; j < len(window); j++ {
			sum := window[i] + window[j]
			if sum == val {
				return true
			}
		}
	}
	return false
}

func SumLowestAndHighest(w []int) int {
	sorted := make(sort.IntSlice, len(w))
	copy(sorted, w)
	sorted.Sort()
	return sorted[0] + sorted[len(sorted)-1]
}

func SumAll(w []int) int {
	s := 0
	for _, v := range w {
		s = s + v
	}
	return s
}
