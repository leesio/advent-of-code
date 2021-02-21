package main

import "fmt"

const (
	n  = 2020
	n2 = 30000000
)

var input = []int{17, 1, 3, 16, 19, 0}

func main() {
	fmt.Printf("part 1: %d\n", getNthNumberSpoken(input, n))
	fmt.Printf("part 2: %d\n", getNthNumberSpoken(input, n2))
}

func getNthNumberSpoken(input []int, n int) int {
	m := make(map[int][]int)

	for s, startingNum := range input {
		m[startingNum] = append(m[startingNum], s)
	}
	prev := input[len(input)-1]
	for cursor := len(input); ; cursor++ {
		spoken := 0
		if turn, ok := m[prev]; ok && len(turn) > 1 {
			spoken = turn[len(turn)-1] - turn[len(turn)-2]
		}
		if cursor == n-1 {
			return spoken
		}
		m[spoken] = append(m[spoken], cursor)
		prev = spoken
	}
}
