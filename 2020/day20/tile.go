package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Tile struct {
	ID     int
	pixels [][]string

	Edges       map[string]string
	Matches     map[string]*Tile
	Connections map[string]*Tile

	AllTiles []*Tile

	refreshCounter int
}

type Tiles []*Tile

func (t Tiles) TopLeft() *Tile {
	corners := t.Corners()
	for _, corner := range corners {
		if _, ok := corner.Matches[top]; ok {
			continue
		}
		if _, ok := corner.Matches[left]; ok {
			continue
		}
		return corner
	}
	return nil
}
func (t Tiles) Corners() []*Tile {
	corners := make([]*Tile, 0)
	for _, tile := range t {
		count := tile.matchCount()
		if count == 2 {
			corners = append(corners, tile)
		}
	}
	if len(corners) != 4 {
		panic(fmt.Errorf("expected 4 corners, got %d", len(corners)))
	}
	return corners
}

func NewTile(id int, pixels [][]string) *Tile {
	t := &Tile{
		ID:     id,
		pixels: pixels,

		Edges:       make(map[string]string),
		Matches:     make(map[string]*Tile),
		Connections: make(map[string]*Tile),
	}
	t.populateEdges()
	return t
}

func (t *Tile) Init(others []*Tile) {
	t.AllTiles = others
	t.refresh()
}

func (t *Tile) Layout() string {
	lines := make([]string, 0)
	for _, row := range t.pixels {
		lines = append(lines, fmt.Sprintf("%s\n", strings.Join(row, "")))
	}
	return strings.Join(lines, "")
}

func (t *Tile) String() string {
	return strconv.Itoa(t.ID)
}

func (t *Tile) matchCount() int {
	return len(t.Matches)
}

func (t *Tile) populateMatches() {
	matches := make(map[string]*Tile)
	for _, other := range t.AllTiles {
		if other.ID == t.ID {
			continue
		}
		for key, edge := range t.Edges {
			for _, otherEdge := range other.Edges {
				if edge == otherEdge {
					matches[key] = other
				} else if edge == reverse(otherEdge) {
					matches[key] = other
				}
			}
		}
	}
	t.Matches = matches
}

func (t *Tile) populateEdges() {
	r, l := make([]string, len(t.pixels)), make([]string, len(t.pixels))
	for i, row := range t.pixels {
		r[i], l[i] = row[len(row)-1], row[0]
	}
	t.Edges = map[string]string{
		right:  strings.Join(r, ""),
		left:   strings.Join(l, ""),
		top:    strings.Join(t.pixels[0], ""),
		bottom: strings.Join(t.pixels[len(t.pixels)-1], ""),
	}
}

func (t *Tile) Refresh(i int) {
	if i == t.refreshCounter {
		return
	}
	t.populateEdges()
	t.populateMatches()
	t.refreshCounter = i
	for _, other := range t.Connections {
		other.Refresh(i)
	}
}

func (t *Tile) refresh() {
	t.Refresh(t.refreshCounter + 1)
}

func (t *Tile) flipped() [][]string {
	clone := make([][]string, len(t.pixels))
	for r := range clone {
		clone[r] = strings.Split(
			reverse(strings.Join(t.pixels[r], "")),
			"",
		)
	}
	return clone
}

func (t *Tile) Flip() {
	t.pixels = t.flipped()
	t.refresh()
}

func (t *Tile) Rotate() {
	clone := make([][]string, len(t.pixels))
	for r := range clone {
		clone[r] = make([]string, len(t.pixels[r]))
		for c := range clone[r] {
			clone[r][c] = t.pixels[len(clone[r])-1-c][r]
		}
	}
	t.pixels = clone
	t.refresh()
}

func (t *Tile) TrimmedPixels() [][]string {
	trimmed := make([][]string, 0)
	for _, row := range t.pixels[1 : len(t.pixels)-1] {
		trimmed = append(trimmed, row[1:len(row)-1])
	}
	return trimmed
}

func (t *Tile) ConnectDown(other *Tile) {
	if other.SearchForOrientation(func() bool {
		reverseMatch, ok := other.Matches[top]
		return ok && reverseMatch == t
	}) {
		t.Connections[bottom] = other
		other.Connections[top] = t
	} else {
		panic(fmt.Errorf("Tried to connect: %d to %d (down) and failed", t.ID, other.ID))
	}
}

func (t *Tile) ConnectRight(other *Tile) {
	if other.SearchForOrientation(func() bool {
		reverseMatch, ok := other.Edges[left]
		return ok && reverseMatch == t.Edges[right]
	}) {
		t.Connections[right] = other
		other.Connections[left] = t
	} else {
		panic(fmt.Errorf("Tried to connect: %d to %d (right) and failed", t.ID, other.ID))
	}
}

func (t *Tile) SearchForOrientation(condition func() bool) bool {
	rotations := 4
	flipped := false
	for {
		if condition() {
			return true
		}
		if rotations > 0 {
			t.Rotate()
			rotations--
			continue
		}
		if !flipped {
			t.Flip()
			flipped = true
			rotations = 4
			continue
		}
		break
	}
	return false
}

func (t *Tile) Valid() bool {
	return len(t.InvalidEdges()) == 0
}

func (t *Tile) InvalidEdges() []string {
	invalidEdges := make([]string, 0)
	for edge, other := range t.Connections {
		if t.Edges[edge] != other.Edges[connectingKey(edge)] {
			invalidEdges = append(invalidEdges, edge)
		}
	}
	return invalidEdges
}

func (t *Tile) CountSeaMonsters() int {
	count := 0
	if t.SearchForOrientation(func() bool {
		for i := 3; i < len(t.pixels); i++ {
			grid := t.pixels[i-3 : i]
			count += containsSeaMonster(grid)
		}
		return count > 0
	}) {
		return count
	}
	return -1
}

func (t *Tile) CountHashes() int {
	count := 0
	for _, row := range t.pixels {
		for _, val := range row {
			if val == "#" {
				count++
			}
		}
	}
	return count
}
