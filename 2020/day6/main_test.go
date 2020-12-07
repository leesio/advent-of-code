package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

var (
	testInput = []string{"abc", "", "a", "b", "c", "", "ab", "ac", "", "a", "a",
		"a", "a", "", "b",
	}
)

func TestGetCounts(t *testing.T) {

	countOne, countTwo := GetCounts(testInput)
	if countOne != 11 {
		t.Errorf("Got count one: %d, expected 11", countOne)
	}
	if countTwo != 6 {
		t.Errorf("Got count two: %d, expected 6", countTwo)
	}
}

func TestWill(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		t.Fail()
	}
	s := bufio.NewScanner(f)
	grid := [][]byte{}
	lines := make([]string, 0)
	for s.Scan() {
		grid = append(grid, s.Bytes())
		lines = append(lines, s.Text())
	}
	fmt.Println(grid)
	match, mismatch := 0, 0
	for l, line := range grid {
		if safeLine := lines[l]; safeLine != string(line) {
			fmt.Println(l, string(line), "vs", safeLine)
			mismatch++
		} else {
			match++
		}
	}
	fmt.Println(match, mismatch)
}
