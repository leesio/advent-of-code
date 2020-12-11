package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/leesio/advent-of-code/2020/util"
)

const (
	floor    = "."
	occupied = "#"
	empty    = "L"
)

func main() {
	input, err := util.GetInput("input")
	if err != nil {
		panic(err)
	}
	fmt.Printf("part 1: %d\n", PartOne(input))
	fmt.Printf("part 2: %d\n", PartTwo(input))
}

func PartOne(input []string) int {
	g, changed := makeGrid(input), -1
	for {
		if g, changed = g.stepAdjacent(); changed == 0 {
			break
		}
	}
	return g.countOccupiedSeats()
}

func PartTwo(input []string) int {
	g, changed := makeGrid(input), -1
	for {
		if g, changed = g.stepVisible(); changed == 0 {
			break
		}
	}
	return g.countOccupiedSeats()
}

type grid [][]string

func makeGrid(input []string) grid {
	grid := make(grid, len(input))
	for l, line := range input {
		grid[l] = strings.Split(line, "")
	}
	return grid
}

func (g grid) adjacentOccupiedSeats(x, y int) int {
	occupiedCount := 0
	for i := max(x-1, 0); i <= min(len(g)-1, x+1); i++ {
		for j := max(y-1, 0); j <= min(len(g[x])-1, y+1); j++ {
			if i == x && j == y {
				continue
			}
			neighbour := g[i][j]
			if neighbour == occupied {
				occupiedCount++
			}
		}
	}
	return occupiedCount
}

func (g grid) visibleOccupiedSeats(x, y int) int {
	occupiedCount := 0
	shouldBreak := func(val string, debugStr string) bool {
		if val == occupied {
			occupiedCount++
			return true
		}
		return val == empty
	}
	row := g[x]
	for i := y + 1; i < len(row); i++ {
		if shouldBreak(row[i], "right") {
			break
		}
	}
	for i := y - 1; i >= 0; i-- {
		if shouldBreak(row[i], "left") {
			break
		}
	}
	for i := x - 1; i >= 0; i-- {
		if shouldBreak(g[i][y], "up") {
			break
		}
	}
	for i := x + 1; i < len(g); i++ {
		if shouldBreak(g[i][y], "down") {
			break
		}
	}

	for i, j := x-1, y+1; i >= 0 && j < len(g[i]); i, j = i-1, j+1 {
		if shouldBreak(g[i][j], "up-right") {
			break
		}
	}

	for i, j := x-1, y-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if shouldBreak(g[i][j], "up-left") {
			break
		}
	}

	for i, j := x+1, y+1; i < len(g) && j < len(g[i]); i, j = i+1, j+1 {
		if shouldBreak(g[i][j], "down-right") {
			break
		}
	}

	for i, j := x+1, y-1; i < len(g) && j >= 0; i, j = i+1, j-1 {
		if shouldBreak(g[i][j], "down-left") {
			break
		}
	}
	return occupiedCount
}

func (g grid) mapGrid(fn func(int, int, string) string) grid {
	clone := make([][]string, len(g))
	for i := 0; i < len(g); i++ {
		row := g[i]
		clone[i] = make([]string, len(row))
		for j := 0; j < len(row); j++ {
			clone[i][j] = fn(i, j, g[i][j])
		}
	}
	return clone
}

func (g grid) stepVisible() (grid, int) {
	changes := 0
	return g.mapGrid(func(x, y int, val string) string {
		occupiedVisible := g.visibleOccupiedSeats(x, y)
		if val == occupied && occupiedVisible >= 5 {
			changes++
			return empty
		}
		if val == empty && occupiedVisible == 0 {
			changes++
			return occupied
		}
		return val
	}), changes
}

func (g grid) stepAdjacent() (grid, int) {
	changes := 0
	return g.mapGrid(func(x, y int, val string) string {
		occupiedNeighbours := g.adjacentOccupiedSeats(x, y)
		if val == occupied && occupiedNeighbours >= 4 {
			changes++
			return empty
		}
		if val == empty && occupiedNeighbours == 0 {
			changes++
			return occupied
		}
		return val
	}), changes
}

func (g grid) String() string {
	lines := make([]string, len(g))
	for r, row := range g {
		lines[r] = strings.Join(row, " ")
	}
	return strings.Join(lines, "\n")
}

func (g grid) countOccupiedSeats() int {
	count := 0
	for _, row := range g {
		for _, col := range row {
			if col == occupied {
				count++
			}
		}
	}
	return count
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
