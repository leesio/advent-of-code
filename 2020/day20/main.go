package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/leesio/advent-of-code/2020/util"
)

const (
	top    = "top"
	bottom = "bottom"
	right  = "right"
	left   = "left"
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
	tiles := ParseInput(input)
	corners := tiles.Corners()
	res := corners[0].ID
	for _, corner := range corners[1:] {
		res *= corner.ID
	}
	return res
}

func printRow(row []*Tile) {
	for t, tile := range row {
		fmt.Printf("%d (%d)", tile.ID, tile.matchCount())
		if t < len(row)-1 {
			fmt.Printf("->")
		}
	}
	fmt.Printf("\n")
}

func buildRow(farLeft *Tile, size int) []*Tile {
	current := farLeft
	row := make([]*Tile, 0)

	for len(row) < size {
		// time.Sleep(1 * time.Second)
		row = append(row, current)
		if len(row) == size {
			break
		}
		nextAcross, ok := current.Matches[right]
		if ok {
			current.ConnectRight(nextAcross)
			current = nextAcross
			continue
		}
		current.Flip()
		nextAcross, ok = current.Matches[right]
		if ok {
			current.ConnectRight(nextAcross)
			current = nextAcross
			continue
		}

		panic(fmt.Errorf("no right to go to: %d %v %v", current.ID, current.Matches, row))
	}
	return row
}

func PartTwo(input []string) int {
	tiles := ParseInput(input)
	size := int(math.Sqrt(float64(len(tiles))))
	_ = size
	topLeft := tiles.TopLeft()
	current := topLeft
	rows := make([][]*Tile, 0)

	for {
		row := buildRow(current, size)
		rows = append(rows, row)

		if len(rows) == size {
			break
		}
		nextDown, ok := current.Matches[bottom]
		if !ok {
			panic(fmt.Errorf("no down to go to: %d %v", current.ID, current.Matches))
		}
		current.ConnectDown(nextDown)
		current = nextDown
	}
	t := buildMegaTile(rows)
	return t.CountHashes() - (15 * t.CountSeaMonsters())
}

func buildMegaTile(grid [][]*Tile) *Tile {
	pixels := make([][]string, 0)
	for r, row := range grid {
		for c, tile := range row {
			trimmed := tile.TrimmedPixels()
			for i, innerRow := range trimmed {
				_, _, _ = pixels, innerRow, c
				rowNum := (r * len(trimmed)) + i
				if len(pixels) <= rowNum {
					pixels = append(pixels, make([]string, 0))
				}
				pixels[rowNum] = append(pixels[rowNum], innerRow...)
			}

		}
	}
	t := &Tile{pixels: pixels}
	return t
}

func ParseInput(input []string) Tiles {
	tiles := make(Tiles, 0)
	id := 0
	pixels := make([][]string, 0)
	for _, line := range input {
		if line == "" {
			tiles = append(tiles, NewTile(id, pixels))
			pixels = make([][]string, 0)
			id = 0
			continue
		}
		if strings.Contains(line, "Tile") {
			id = util.MustAtoi(
				strings.TrimSuffix(strings.TrimPrefix(line, "Tile "), ":"),
			)
			continue
		}
		pixels = append(pixels, strings.Split(line, ""))
	}
	tiles = append(tiles, NewTile(id, pixels))
	for _, tile := range tiles {
		tile.Init(tiles)
	}
	return tiles
}

func connectingKey(key string) string {
	switch key {
	case bottom:
		return top
	case top:
		return bottom
	case left:
		return right
	case right:
		return left
	default:
		panic(fmt.Errorf("unknown key: %s", key))
	}
}

func reverse(s string) string {
	r := make([]string, len(s))
	for i, c := range strings.Split(s, "") {
		r[len(s)-1-i] = c
	}
	return strings.Join(r, "")
}
