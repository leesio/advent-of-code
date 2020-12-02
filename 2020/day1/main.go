package main

import (
	"fmt"

	"github.com/leesio/advent-of-code/2020/util"
)

func partOne(input []int) int {
	m := make(map[int]struct{})
	for _, num := range input {
		m[num] = struct{}{}
		if _, ok := m[2020-num]; ok {
			return num * (2020 - num)
		}
	}
	return -1
}
func partTwo(input []int) int {
	h := make(map[int]struct{})
	for _, n := range input {
		for _, m := range input {
			h[n] = struct{}{}
			if _, ok := h[2020-m-n]; ok {
				return m * n * (2020 - m - n)
			}
		}
	}
	return -1
}
func main() {
	input, err := util.GetInput("input")
	if err != nil {
		panic(err)
	}
	inputNums, err := util.Atoi(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("part 1: %d\n", partOne(inputNums))
	fmt.Printf("part 2: %d\n", partTwo(inputNums))
}
