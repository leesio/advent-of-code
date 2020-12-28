package main

import (
	"fmt"
	"strings"

	"github.com/leesio/advent-of-code/2020/util"
)

const (
	se = "se"
	sw = "sw"
	ne = "ne"
	nw = "nw"
	w  = "w"
	e  = "e"
)

type Colour int

const (
	white Colour = iota
	black
)

var directions = []string{e, w, se, sw, ne, nw}

func main() {
	input, err := util.GetInput("input")
	if err != nil {
		panic(err)
	}
	fmt.Printf("part 1: %d\n", PartOne(input))
	fmt.Printf("part 2: %d\n", PartTwo(input))
}

func PartOne(input []string) int {
	sequences := ParseInput(input)
	floor := NewFloor()
	floor.FlipTiles(sequences)
	return len(floor.BlackTiles())
}

func PartTwo(input []string) int {
	sequences := ParseInput(input)
	floor := NewFloor()
	floor.FlipTiles(sequences)

	for i := 0; i < 100; i++ {
		floor.ElapseDay()
	}
	return len(floor.BlackTiles())
}

func ParseInput(input []string) [][]string {
	sequences := make([][]string, len(input))
	for l, line := range input {
		tail := line
		sequence := make([]string, 0)
		for len(tail) > 0 {
			chars := 1
			if strings.HasPrefix(tail, "s") || strings.HasPrefix(tail, "n") {
				chars = 2
			}
			dir := tail[:chars]
			tail = tail[chars:]
			sequence = append(sequence, dir)
		}
		sequences[l] = sequence
	}
	return sequences
}

type Tile struct {
	x      int
	y      int
	colour Colour
}

type Tiles struct {
	underlying map[string]*Tile
}

func NewTiles() *Tiles {
	return &Tiles{underlying: make(map[string]*Tile)}
}

func (t *Tiles) key(x, y int) string {
	return fmt.Sprintf("%d:%d", x, y)
}

func (t *Tiles) Get(x, y int) Colour {
	key := t.key(x, y)
	if tile, ok := t.underlying[key]; ok {
		return tile.colour
	}
	return white
}

func (t *Tiles) BlackTiles() map[string]*Tile {
	return t.underlying
}

func (t *Tiles) Flip(x, y int) {
	key := t.key(x, y)
	if _, ok := t.underlying[key]; ok {
		delete(t.underlying, key)
	} else {
		t.underlying[key] = &Tile{x: x, y: y, colour: black}
	}
}

func NewTile(x, y int) *Tile {
	return &Tile{x: x, y: y, colour: white}
}

type Floor struct {
	tiles *Tiles
}

func NewFloor() *Floor {
	return &Floor{tiles: NewTiles()}
}

func getRelativeTileCoords(x, y int, dir string) (int, int) {
	if dir == e {
		return x + 1, y
	}
	if dir == w {
		return x - 1, y
	}
	if strings.HasSuffix(dir, "w") {
		x--
	}
	if abs(y%2) == 1 {
		x++
	}
	if strings.HasPrefix(dir, "s") {
		y++
	} else if strings.HasPrefix(dir, "n") {
		y--
	}
	return x, y
}

func (f *Floor) FlipTile(sequence []string) {
	x, y := 0, 0
	for _, dir := range sequence {
		x, y = getRelativeTileCoords(x, y, dir)
	}
	f.tiles.Flip(x, y)
}

func (f *Floor) FlipTiles(sequences [][]string) {
	for _, sequence := range sequences {
		f.FlipTile(sequence)
	}
}

func (f *Floor) BlackTiles() map[string]*Tile {
	return f.tiles.BlackTiles()
}

type Neighbour struct {
	*Tile
	blackNeighbours int
}

func (f *Floor) ElapseDay() {
	allTiles := make(map[string]*Neighbour)
	blackTiles := f.BlackTiles()
	for _, tile := range blackTiles {
		if _, ok := allTiles[key(tile.x, tile.y)]; !ok {
			allTiles[key(tile.x, tile.y)] = &Neighbour{Tile: tile}
		}
		for _, dir := range directions {
			x, y := getRelativeTileCoords(tile.x, tile.y, dir)
			neighbour, ok := allTiles[key(x, y)]
			if !ok {
				neighbourTile, ok := blackTiles[key(x, y)]
				if !ok {
					neighbourTile = NewTile(x, y)
				}
				neighbour = &Neighbour{Tile: neighbourTile}
				allTiles[f.tiles.key(x, y)] = neighbour
			}
			neighbour.blackNeighbours++
		}
	}
	toSwitch := make([]*Tile, 0)
	for _, neighbour := range allTiles {

		if (neighbour.colour == black && neighbour.blackNeighbours > 2 || neighbour.blackNeighbours == 0) ||
			neighbour.colour == white && neighbour.blackNeighbours == 2 {
			toSwitch = append(toSwitch, neighbour.Tile)
		}
	}
	for _, tile := range toSwitch {
		f.tiles.Flip(tile.x, tile.y)
	}
}

func key(x, y int) string {
	return fmt.Sprintf("%d:%d", x, y)
}
func abs(n int) int {
	if n >= 0 {
		return n
	}
	return 0 - n
}
