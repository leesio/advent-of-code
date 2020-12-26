package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/leesio/advent-of-code/2020/util"
)

func main() {
	fmt.Printf("part 1: %s\n", PartOne("487912365"))
	fmt.Printf("part 2: %s\n", PartTwo("487912365"))
}

func PartOne(input string) string {
	ring := NewRing(input)
	for i := 0; i < 100; i++ {
		ring.Move()
	}
	fmt.Println(ring)
	idx := findIndex(ring.underlying, 1)
	offset := idx + 1
	output := make([]string, 0)
	for i := 0; i < len(ring.underlying)-1; i++ {
		output = append(output, strconv.Itoa(ring.underlying[(i+offset)%len(ring.underlying)]))
	}
	return strings.Join(output, "")
}

func PartTwo(input string) string {
	ring := NewRing(input)
	ring.Pad()
	a := time.Now()
	for i := 0; i < 10_000_000; i++ {
		ring.Move()
		if i > 0 && i%1000 == 0 {
			fmt.Println(i, time.Now().Sub(a))
			a = time.Now()
		}
	}
	idx := findIndex(ring.underlying, 1)
	fmt.Println(ring.underlying[idx+1], ring.underlying[idx+2])
	return ""
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

func (r *Ring) String() string {
	s := make([]string, len(r.underlying))
	for i, v := range r.underlying {
		if v == r.cursorVal {
			s[i] = fmt.Sprintf("(%s)", strconv.Itoa(v))
		} else {
			s[i] = fmt.Sprintf(" %s ", strconv.Itoa(v))
		}
	}
	return strings.Join(s, "")
}
func (r *Ring) Pad() {
	pad := make([]int, 999999-len(r.underlying))
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
