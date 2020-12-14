package main

import (
	"fmt"
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

func PartOne(input []string) int64 {
	output := make(map[int]int64)
	operations, err := parseInput(input)
	if err != nil {
		panic(err)
	}
	for _, operation := range operations {
		output[operation.register] = operation.mask.Apply(operation.decimal)
	}
	var res int64
	for _, val := range output {
		res += val
	}
	return res
}
func PartTwo(input []string) int64 {
	output := make(map[int]int64)
	operations, err := parseInput(input)
	if err != nil {
		panic(err)
	}
	for _, operation := range operations {
		for _, register := range operation.mask.GetRegisters(operation.register) {
			output[register] = operation.decimal
		}
	}
	var res int64
	for _, val := range output {
		res += val
	}
	return res
}

func parseInput(input []string) ([]*Operation, error) {
	operations := make([]*Operation, 0)
	var mask *Mask
	for _, line := range input {
		if strings.HasPrefix(line, "mask = ") {
			mask = NewMask(strings.TrimPrefix(line, "mask = "))
			continue
		}
		parts := strings.Split(line, " = ")
		if len(parts) != 2 {
			return nil, fmt.Errorf("unexpected input format")
		}
		registerRaw := strings.TrimSuffix(strings.TrimPrefix(parts[0], "mem["), "]")
		register, err := strconv.Atoi(registerRaw)
		if err != nil {
			return nil, err
		}
		decimal, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		operations = append(operations, &Operation{
			decimal:  int64(decimal),
			register: register,
			mask:     mask,
		})
	}
	return operations, nil
}

type Operation struct {
	mask     *Mask
	register int
	decimal  int64
}

type Mask struct {
	ones   int64
	zeroes int64
	raw    string
}

func NewMask(s string) *Mask {
	ones := make([]string, 36)
	zeroes := make([]string, 36)
	for i, s := range strings.Split(s, "") {
		if s == "X" {
			zeroes[i], ones[i] = "1", "0"
		} else {
			zeroes[i], ones[i] = s, s
		}
	}
	onesVal, err := strconv.ParseInt(strings.Join(ones, ""), 2, 64)
	if err != nil {
		panic(err)
	}
	zeroesVal, err := strconv.ParseInt(strings.Join(zeroes, ""), 2, 64)
	if err != nil {
		panic(err)
	}
	return &Mask{
		ones:   onesVal,
		zeroes: zeroesVal,
		raw:    s,
	}
}

func (m *Mask) Apply(i int64) int64 {
	return i&m.zeroes | m.ones
}

type registerAddresses []int

func (r registerAddresses) shift() {
	for i := range r {
		r[i] = r[i] << 1
	}
}
func (r registerAddresses) add(val int) {
	for i := range r {
		r[i] = r[i] + val
	}
}
func (m *Mask) GetRegisters(i int) []int {
	inAddr := strings.Split(fmt.Sprintf("%036b", i), "")
	outAddrs := registerAddresses{0}
	for b, bit := range strings.Split(m.raw, "") {
		outAddrs.shift()
		switch bit {
		case "0":
			if inAddr[b] == "1" {
				outAddrs.add(1)
			}
		case "1":
			outAddrs.add(1)
		case "X":
			split := make(registerAddresses, len(outAddrs))
			copy(split, outAddrs)
			outAddrs.add(1)
			outAddrs = append(outAddrs, split...)
		}
	}
	return outAddrs
}
