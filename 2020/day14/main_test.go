package main

import (
	"testing"
)

var (
	testInputOne = []string{
		"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
		"mem[8] = 11",
		"mem[7] = 101",
		"mem[8] = 0",
	}
	testInputTwo = []string{
		"mask = 000000000000000000000000000000X1001X",
		"mem[42] = 100",
		"mask = 00000000000000000000000000000000X0XX",
		"mem[26] = 1",
	}
)

func TestMask(t *testing.T) {
	mask := NewMask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")

	if res := mask.Apply(11); res != 73 {
		t.Errorf("Got %d applying mask to %d, expected %d", res, 11, 73)
	}
	if res := mask.Apply(101); res != 101 {
		t.Errorf("Got %d applying mask to %d, expected %d", res, 101, 101)
	}

	mask = NewMask("000000000000000000000000000000X1001X")
	expResults := map[int]bool{26: true, 27: true, 58: true, 59: true}
	registers := mask.GetRegisters(42)
	for _, r := range registers {
		if expResults[r] {
			delete(expResults, r)
			continue
		}
		t.Errorf("Got %d register, expected one of %v", r, expResults)
	}
	if len(expResults) != 0 {
		t.Errorf("Missed some results %v", expResults)
	}
}
func TestPartOne(t *testing.T) {
	if res := PartOne(testInputOne); res != 165 {
		t.Errorf("Got %d for part one, expected %d", res, 165)
	}

}
func TestPartTwo(t *testing.T) {
	if res := PartTwo(testInputTwo); res != 208 {
		t.Errorf("Got %d for part two, expected %d", res, 208)
	}
}
