package main

import (
	"fmt"
	"strings"

	"github.com/leesio/advent-of-code/2020/util"
)

func main() {
	input, err := util.GetInput("input")
	if err != nil {
		panic(err)
	}
	one, two := GetCounts(input)
	fmt.Printf("part 1: %d\n", one)
	fmt.Printf("part 2: %d\n", two)
}

type answerCounts map[string]int

func (a answerCounts) FirstCount() int {
	return len(a)
}

func (a answerCounts) SecondCount(groupSize int) int {
	count := 0
	for _, val := range a {
		if val == groupSize {
			count++
		}
	}
	return count
}

func GetCounts(input []string) (int, int) {
	count := 0
	countTwo := 0
	groupSize := 0
	m := make(answerCounts)
	for _, line := range input {
		if line == "" {
			count = count + m.FirstCount()
			countTwo = countTwo + m.SecondCount(groupSize)

			m = make(map[string]int)
			groupSize = 0
			continue
		}
		parts := strings.Split(line, "")
		groupSize++
		for _, part := range parts {
			m[part]++
		}
	}
	count = count + m.FirstCount()
	countTwo = countTwo + m.SecondCount(groupSize)
	return count, countTwo
}
