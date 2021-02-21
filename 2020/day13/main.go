package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/leesio/advent-of-code/2020/util"
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
	ts, ids := parseInput(input)
	bestID := 0
	min := 1 << 31
	for _, id := range ids {
		if id == -1 {
			continue
		}
		firstAfterTs := firstDepartureAfterTS(id, ts)
		if firstAfterTs < min {
			min, bestID = firstAfterTs, id
		}
	}
	waitTime := min - ts
	return waitTime * bestID
}
func firstDepartureAfterTS(id int, ts int) int {
	if ts%id == 0 {
		return ts
	}
	return (1 + (ts / id)) * id
}

func PartTwo(input []string) int {
	_, ids := parseInput(input)
	as := make([]*big.Int, 0)
	ns := make([]*big.Int, 0)
	for i, id := range ids {
		if id == -1 {
			continue
		}
		as = append(as, big.NewInt(int64(id)))
		ns = append(ns, big.NewInt(int64(id-i)))
	}
	res := chineseRemainder(ns, as)
	return int(res.Int64())
}

func parseInput(input []string) (int, []int) {
	ts, err := strconv.Atoi(input[0])
	if err != nil {
		panic(err)
	}
	ids := make([]int, 0)
	for _, idRaw := range strings.Split(input[1], ",") {
		var id int
		if idRaw == "x" {
			id = -1
		} else {
			var err error
			id, err = strconv.Atoi(idRaw)
			if err != nil {
				panic(err)
			}
		}
		ids = append(ids, id)
	}
	return ts, ids
}

// Code stolen from the internet and adapted slightly
// https://rosettacode.org/wiki/Chinese_remainder_theorem
// https://www.dave4math.com/mathematics/chinese-remainder-theorem/
func chineseRemainder(a, n []*big.Int) *big.Int {
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		// the pairs are coprime so the GCD should be 1
		z.GCD(nil, &s, n1, &q)
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p)
}
