package main

import (
	"fmt"
	"strings"

	"github.com/leesio/advent-of-code/2020/util"
)

const (
	active   = "#"
	inactive = "."
)

type Range struct {
	min, max int
}

func NewRange(min, max int) *Range {
	return &Range{min, max}
}

func (r *Range) String() string {
	return fmt.Sprintf("%d-%d", r.min, r.max)
}

func (r *Range) clone() *Range {
	return &Range{r.min, r.max}
}
func (r *Range) Check(v int) {
	if v < r.min {
		r.min = v
	}
	if v > r.max {
		r.max = v
	}
}

type infiniteGrid struct {
	m      map[string]bool
	wRange *Range
	xRange *Range
	yRange *Range
	zRange *Range
}

func NewInfiniteGrid(sl []string) *infiniteGrid {
	g := &infiniteGrid{
		m:      make(map[string]bool),
		wRange: NewRange(0, 0),
		xRange: NewRange(0, 0),
		yRange: NewRange(0, 0),
		zRange: NewRange(0, 0),
	}
	for l, line := range sl {
		for c, char := range strings.Split(line, "") {
			if char == active {
				g.set(0, c, l, 0, active)
			}
		}
	}

	return g
}

func countActive(sl []string) int {
	count := 0
	for _, val := range sl {
		if val == active {
			count++
		}
	}
	return count
}
func (ig *infiniteGrid) fresh() *infiniteGrid {
	m := make(map[string]bool)
	return &infiniteGrid{
		xRange: ig.xRange.clone(),
		yRange: ig.yRange.clone(),
		zRange: ig.zRange.clone(),
		wRange: ig.wRange.clone(),
		m:      m,
	}
}
func (ig *infiniteGrid) step4d() *infiniteGrid {
	next := ig.fresh()
	for z := ig.zRange.min - 1; z <= ig.zRange.max+1; z++ {
		for y := ig.yRange.min - 1; y <= ig.yRange.max+1; y++ {
			for x := ig.xRange.min - 1; x <= ig.xRange.max+1; x++ {
				for w := ig.wRange.min - 1; w <= ig.wRange.max+1; w++ {
					neighbours := ig.neighbours4d(w, x, y, z)
					current := ig.get(w, x, y, z)
					if an := countActive(neighbours); an == 3 {
						next.set(w, x, y, z, active)
					} else if an == 2 && current == active {
						next.set(w, x, y, z, active)
					}
				}
			}
		}
	}
	return next
}
func (ig *infiniteGrid) step() *infiniteGrid {
	next := ig.fresh()
	for z := ig.zRange.min - 1; z <= ig.zRange.max+1; z++ {
		for y := ig.yRange.min - 1; y <= ig.yRange.max+1; y++ {
			for x := ig.xRange.min - 1; x <= ig.xRange.max+1; x++ {
				neighbours := ig.neighbours(x, y, z)
				current := ig.get(0, x, y, z)
				if an := countActive(neighbours); an == 3 {
					next.set(0, x, y, z, active)
				} else if an == 2 && current == active {
					next.set(0, x, y, z, active)
				}
			}
		}
	}
	return next
}

func (ig *infiniteGrid) countActive() int {
	active := 0
	for _, v := range ig.m {
		if v {
			active++
		}
	}
	return active
}

func (ig *infiniteGrid) key(w, x, y, z int) string {
	return fmt.Sprintf("%d:%d:%d:%d", w, x, y, z)
}

func (ig *infiniteGrid) setLimits(w, x, y, z int) {
	ig.xRange.Check(x)
	ig.yRange.Check(y)
	ig.zRange.Check(z)
	ig.wRange.Check(w)
}

func (ig *infiniteGrid) set(w, x, y, z int, val string) {
	if val == active {
		ig.m[ig.key(w, x, y, z)] = true
		ig.setLimits(w, x, y, z)
	} else {
		ig.m[ig.key(w, x, y, z)] = false
	}
}
func (ig *infiniteGrid) get(w, x, y, z int) string {
	if ig.m[ig.key(w, x, y, z)] {
		return active
	}
	return inactive
}

func (ig *infiniteGrid) neighbours4d(w, x, y, z int) []string {
	neighbours := make([]string, 0)
	for i := z - 1; i <= z+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			for k := x - 1; k <= x+1; k++ {
				for l := w - 1; l <= w+1; l++ {
					if k == x && j == y && i == z && l == w {
						// skip self
						continue
					}
					neighbours = append(neighbours, ig.get(l, k, j, i))
				}
			}
		}
	}
	return neighbours
}
func (ig *infiniteGrid) neighbours(x, y, z int) []string {
	neighbours := make([]string, 0)
	for i := z - 1; i <= z+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			for k := x - 1; k <= x+1; k++ {
				if k == x && j == y && i == z {
					// skip self
					continue
				}
				neighbours = append(neighbours, ig.get(0, k, j, i))
			}
		}
	}
	return neighbours
}

func PartOne(input []string) int {
	g := NewInfiniteGrid(input)
	for i := 0; i < 6; i++ {
		g = g.step()
	}
	return g.countActive()
}
func PartTwo(input []string) int {
	g := NewInfiniteGrid(input)
	for i := 0; i < 6; i++ {
		g = g.step4d()
	}
	return g.countActive()
}
func main() {
	input, err := util.GetInput("input")
	if err != nil {
		panic(err)
	}
	fmt.Printf("part 1: %d\n", PartOne(input))
	fmt.Printf("part 2: %d\n", PartTwo(input))
}
