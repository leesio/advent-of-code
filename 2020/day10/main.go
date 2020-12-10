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
	input, err := util.Atoi(rawInput)
	if err != nil {
		panic(err)
	}
	partOne := PartOne(input)
	fmt.Printf("part 1: %d\n", partOne)

	partTwo := PartTwo(input)
	fmt.Printf("part 2: %d\n", partTwo)
	partTwo = PartTwoCopiedFromJon(input)
	fmt.Printf("part 2: %d\n", partTwo)
}

type Adapters sort.IntSlice

func (a Adapters) ValidForInput(inputVoltage int) []int {
	min, max := inputVoltage+1, inputVoltage+3
	return a.numBetween(min, max)
}

func (a Adapters) ValidForOutput(outputVoltage int) []int {
	max, min := outputVoltage-1, outputVoltage-3
	return a.numBetween(min, max)
}

func (a Adapters) numBetween(min, max int) []int {
	valid := make([]int, 0)
	for _, adapter := range a {
		if adapter < min {
			continue
		}
		if adapter > max {
			break
		}
		valid = append(valid, adapter)
	}
	return valid
}

func isPinchPoint(adapters Adapters, n int) bool {
	validPrevious := adapters.ValidForOutput(n)
	if len(validPrevious) != 1 {
		return false
	}
	onlyPrevious := validPrevious[0]
	validOnward := adapters.ValidForInput(onlyPrevious)
	return len(validOnward) == 1
}

func PartOne(input []int) int {
	sorted := make(sort.IntSlice, len(input))
	copy(sorted, input)
	sorted.Sort()
	inputVoltage := 0
	gaps := make(map[int]int)
	for _, outputVoltage := range sorted {
		gap := outputVoltage - inputVoltage
		gaps[gap]++
		inputVoltage = outputVoltage
	}
	// The device voltage is 3V higher than the input, so add a last 3V gap
	gaps[3]++
	return gaps[1] * gaps[3]
}

func PartTwo(input []int) int {
	sorted := make(sort.IntSlice, len(input))
	copy(sorted, input)
	sorted.Sort()
	sorted = append(sorted, sorted[len(sorted)-1]+3)

	adapters := Adapters(sorted)
	answer, collected := 1, 0
	for _, val := range sorted {
		if !isPinchPoint(adapters, val) {
			collected++
			continue
		}
		answer *= tribonacci(collected + 3)
		collected = 0
	}
	return answer
}

func PartTwoCopiedFromJon(input []int) int {
	sorted := make(sort.IntSlice, len(input))
	copy(sorted, input)
	sorted.Sort()

	sorted = append(sorted, sorted[len(sorted)-1]+3)
	prev := 0
	answer := 1
	current := 0
	for i := 0; i < len(sorted); i++ {
		val := sorted[i]
		diff := val - prev
		if diff < 3 {
			current++
		} else if current > 0 {
			switch current {
			case 4:
				answer = answer * 7
			case 3:
				answer = answer * 4
			case 2:
				answer = answer * 2
			}
			current = 0
		}
		prev = val
	}
	return answer
}

func tribonacci(n int) int {
	if n < 3 {
		return 0
	}
	if n == 3 {
		return 1
	}
	return tribonacci(n-1) + tribonacci(n-2) + tribonacci(n-3)
}
