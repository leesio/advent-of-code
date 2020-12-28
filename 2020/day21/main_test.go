package main

import (
	"testing"
)

var testInput = []string{
	"mxmxvkd kfcds sqjhc nhms (contains dairy, fish)",
	"trh fvjkl sbzzf mxmxvkd (contains dairy)",
	"sqjhc fvjkl (contains soy)",
	"sqjhc mxmxvkd sbzzf (contains fish)",
}

func TestPartOne(t *testing.T) {
	if res := PartOne(testInput); res != 5 {
		t.Errorf("Got %d, expected %d for part one", res, 5)
	}
}
func TestPartTwo(t *testing.T) {
	exp := "mxmxvkd,sqjhc,fvjkl"
	if res := PartTwo(testInput); res != exp {
		t.Errorf("Got %s, expected %s for part two", res, exp)
	}
}
