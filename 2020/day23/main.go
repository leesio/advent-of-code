// https://adventofcode.com/2020/day/23

package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/leesio/advent-of-code/2020/util"
)

func main() {
	fmt.Printf("part 1: %s\n", PartOne("487912365"))
	fmt.Printf("part 2: %d\n", PartTwo("487912365"))
}

func PartOne(input string) string {
	ring := NewRing(input)
	for i := 0; i < 100; i++ {
		ring.Move()
	}
	idx := findIndex(ring.underlying, 1)
	offset := idx + 1
	output := make([]string, 0)
	for i := 0; i < len(ring.underlying)-1; i++ {
		output = append(output, strconv.Itoa(ring.underlying[(i+offset)%len(ring.underlying)]))
	}
	return strings.Join(output, "")
}

func PartTwo(input string) int {
	ring := NewRing(input)
	ring.Pad()
	first := NewLinkedCups(ring)
	c := first
	for i := 0; i < 10_000_000; i++ {
		c = c.Move()
	}
	one := c.index[1]
	a := one.NextCup
	b := a.NextCup
	return a.Value * b.Value
}

type Ring struct {
	underlying []int
	cursor     int
	cursorVal  int
	min        int
	max        int

	moves int
}

func findIndex(sl []int, val int) int {
	for i, v := range sl {
		if val == v {
			return i
		}
	}
	return -1
}

func (r *Ring) updateCursor() {
	cursor := findIndex(r.underlying, r.cursorVal)
	r.cursor = (cursor + 1) % len(r.underlying)
	r.cursorVal = r.underlying[r.cursor]
}

func (r *Ring) Pad() {
	pad := make([]int, 1_000_000-len(r.underlying))
	for i := range pad {
		pad[i] = i + r.max + 1
	}
	r.underlying = append(r.underlying, pad...)
	r.max = 1_000_000
}

func (r *Ring) Move() {
	clone := make([]int, len(r.underlying)-3)
	pickup := make([]int, 3)
	copy(clone, r.underlying)
	length := len(r.underlying)

	offset := r.cursor + 1
	for i := 0; i < 3; i++ {
		idx := (i + offset) % length
		pickup[i] = r.underlying[idx]
	}
	if (offset + 3) > length {
		for i := 3; i < length; i++ {
			idx := (i + offset) % length
			clone[i-3] = r.underlying[idx]
		}
	} else {
		clone = append(r.underlying[:offset], r.underlying[offset+3:]...)
	}
	var dstIdx int
	label := r.underlying[r.cursor] - 1
	for {
		idx := findIndex(clone, label)
		if idx >= 0 {
			dstIdx = idx
			break
		}
		if label-1 < r.min {
			label = r.max
		} else {
			label--
		}
	}
	newUnderlying := make([]int, 0)
	newUnderlying = append(newUnderlying, clone[:dstIdx+1]...)
	newUnderlying = append(newUnderlying, pickup...)
	newUnderlying = append(newUnderlying, clone[dstIdx+1:]...)
	r.underlying = newUnderlying
	r.updateCursor()
	r.moves++
}

func NewRing(input string) *Ring {
	underlying := make([]int, len(input))
	for i, s := range strings.Split(input, "") {
		underlying[i] = util.MustAtoi(s)
	}
	min, max := 10000, 0
	for _, n := range underlying {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return &Ring{
		underlying: underlying,
		cursorVal:  underlying[0],
		min:        min,
		max:        max,
	}
}

type CupIndex map[int]*Cup

type Cup struct {
	Value   int
	NextCup *Cup

	index  CupIndex
	picked CupIndex
	max    int
	min    int
}

// NewLinkedCups creates a linked list of Cups.  Duplicating the slice each move
// makes the naive implementation too slow for part 2
func NewLinkedCups(ring *Ring) *Cup {
	index := make(map[int]*Cup)
	first := &Cup{
		Value: ring.underlying[0],
		index: index,
	}
	index[first.Value] = first
	prev := first
	max, min := prev.Value, prev.Value
	for i := 1; i < len(ring.underlying); i++ {
		current := &Cup{
			Value: ring.underlying[i],
			index: index,
		}
		if current.Value > max {
			max = current.Value
		}
		if current.Value < min {
			min = current.Value
		}
		index[current.Value] = current
		prev.NextCup = current
		prev = current
	}
	for _, cup := range index {
		cup.min = min
		cup.max = max
	}
	last := index[ring.underlying[len(ring.underlying)-1]]
	last.NextCup = first
	return first
}

func (c *Cup) String() string {
	return strconv.Itoa(c.Value)
}

func (c *Cup) Move() *Cup {
	disconnectedCups := make(map[int]bool)
	head, tail := c.Pick()
	c.NextCup = tail.NextCup
	tail.NextCup = nil
	cur := head
	for {
		disconnectedCups[cur.Value] = true
		if cur.NextCup == nil {
			break
		}
		cur = cur.NextCup
	}
	dst := c.Value - 1
	var dstCup *Cup
	var ok bool
	for {
		if dst == 0 {
			dst = c.max
		}
		if !disconnectedCups[dst] {
			dstCup, ok = c.index[dst]
			if !ok {
				fmt.Println("looking for", dst, "in", c.index)
			}
			break
		}
		dst--
	}
	link := dstCup.NextCup
	dstCup.NextCup = head
	tail.NextCup = link
	return c.NextCup
}

func (c *Cup) Pick() (*Cup, *Cup) {
	current := c
	head := c.NextCup
	for i := 0; i < 3; i++ {
		next := current.NextCup
		current = next
	}
	return head, current
}
