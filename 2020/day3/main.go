package main

import (
	"fmt"
	"strings"

	"github.com/leesio/advent-of-code/2020/util"
)

const (
	tree = "#"
)

func main() {
	input, err := util.GetInput("input")
	if err != nil {
		panic(err)
	}
	parsedInput := ParseInput(input)
	fmt.Printf("part 1: %d\n", PartOne(parsedInput))
	fmt.Printf("part 2: %d\n", PartTwo(parsedInput))

}

func ParseInput(s []string) [][]string {
	p := make([][]string, len(s))
	for l, line := range s {
		p[l] = strings.Split(line, "")
	}
	return p
}

func PartTwo(input [][]string) int {
	traverse := []struct {
		x int
		y int
	}{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	res := 1
	for _, t := range traverse {
		res = res * CountTreesEncountered(input, t.x, t.y)
	}
	return res
}

func PartOne(input [][]string) int {
	return CountTreesEncountered(input, 3, 1)
}
func CountTreesEncountered(input [][]string, mvX, mvY int) int {
	trees := 0
	for x, y := 0, 0; y < len(input); y = y + mvY {
		if input[y][x] == tree {
			trees++
		}
		x = (x + mvX) % len(input[y])
	}
	return trees
}
