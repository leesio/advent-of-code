package main

import "fmt"

func containsSeaMonster(grid [][]string) int {
	if len(grid) != 3 {
		panic(fmt.Errorf("non 3 sized grid"))
	}
	count := 0
	for x := 20; x < len(grid[0]); x++ {
		grid := [][]string{
			grid[0][x-20 : x],
			grid[1][x-20 : x],
			grid[2][x-20 : x],
		}
		if isSeaMonster(grid) {
			count++
		}
	}
	return count

}
func isSeaMonster(grid [][]string) bool {
	h := "#"
	if len(grid[0]) != 20 {
		panic(fmt.Errorf("non 20 sized grid"))
	}
	//                   #
	// #    ##    ##    ###
	//  #  #  #  #  #  #
	coords := []struct {
		x int
		y int
	}{
		{0, 18},
		{1, 0}, {1, 5}, {1, 6}, {1, 11}, {1, 12}, {1, 17}, {1, 18}, {1, 19},
		{2, 1}, {2, 4}, {2, 7}, {2, 10}, {2, 13}, {2, 16},
	}
	for _, c := range coords {
		if grid[c.x][c.y] != h {
			return false
		}
	}
	return true
}
