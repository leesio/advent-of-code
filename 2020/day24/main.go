package main

import (
	"fmt"
	"math"
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
	store := NewTileStore()
	reference := store.Get(0, 0)
	for _, sequence := range sequences {
		final := reference.FollowSequence(store, sequence)
		final.SwitchColour()
	}
	return store.CountBlack()
}

func PartTwo(input []string) int {
	sequences := ParseInput(input)
	store := NewTileStore()
	reference := store.Get(0, 0)
	for _, sequence := range sequences {
		final := reference.FollowSequence(store, sequence)
		final.SwitchColour()
	}

	for i := 0; i < 100; i++ {
		store.ElapseDay()
	}
	return store.CountBlack()
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

type TileStore struct {
	store map[string]*Tile
}

func NewTileStore() *TileStore {
	return &TileStore{
		store: make(map[string]*Tile),
	}
}

func (ts *TileStore) key(x, y int) string {
	return fmt.Sprintf("%d:%d", x, y)
}
func (ts *TileStore) CountBlack() int {
	count := 0
	for _, tile := range ts.store {
		if tile.colour == "black" {
			count++
		}
	}
	return count
}

func (ts *TileStore) Get(x, y int) *Tile {
	key := ts.key(x, y)
	if tile, ok := ts.store[key]; ok {
		return tile
	}
	tile := &Tile{
		x:      x,
		y:      y,
		colour: "white",
	}
	ts.store[key] = tile
	return tile
}

type Tile struct {
	x      int
	y      int
	colour string
}

func (t *Tile) FollowSequence(store *TileStore, sequence []string) *Tile {
	current := t
	for _, dir := range sequence {
		current = current.GetNeighbour(store, dir)
	}
	return current
}

func (t *Tile) GetNeighbour(store *TileStore, dir string) *Tile {
	y := t.y
	if strings.HasPrefix(dir, "s") {
		y = t.y + 1
	} else if strings.HasPrefix(dir, "n") {
		y = t.y - 1
	}
	var x int
	if dir == e {
		x = t.x + 1
	} else if dir == w {
		x = t.x - 1
	} else if strings.HasSuffix(dir, "e") {
		if abs(t.y%2) == 1 {
			x = t.x + 1
		} else {
			x = t.x
		}
	} else if strings.HasSuffix(dir, "w") {
		if abs(t.y%2) == 1 {
			x = t.x
		} else {
			x = t.x - 1
		}
	}
	return store.Get(x, y)
}

func (ts *TileStore) ElapseDay() {
	toSwitch := make([]*Tile, 0)
	for _, tile := range ts.store {
		if tile.colour == "black" {
			tile.BlackNeighbours(ts)
		}
	}
	for _, tile := range ts.store {
		n := tile.BlackNeighbours(ts)
		if tile.colour == "black" {
			if n > 2 || n == 0 {
				toSwitch = append(toSwitch, tile)
			}
		} else {
			if n == 2 {
				toSwitch = append(toSwitch, tile)
			}
		}
	}
	for _, tile := range toSwitch {
		tile.SwitchColour()
	}
}
func (t *Tile) BlackNeighbours(store *TileStore) int {
	count := 0
	for _, dir := range directions {
		neighbour := t.GetNeighbour(store, dir)
		if neighbour.colour == "black" {
			count++
		}
	}
	return count
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}

func (t *Tile) SwitchColour() {
	if t.colour == "black" {
		t.colour = "white"
	} else {
		t.colour = "black"
	}
}
