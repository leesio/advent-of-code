package main

import (
	"testing"
)

var testInput = []string{
	"light red bags contain 1 bright white bag, 2 muted yellow bags.",
	"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
	"bright white bags contain 1 shiny gold bag.",
	"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
	"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
	"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
	"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
	"faded blue bags contain no other bags.",
	"dotted black bags contain no other bags.",
}

func TestBagContains(t *testing.T) {
	adj, colour := "shiny", "gold"
	tree := BuildTree(testInput)
	count := 0
	for _, bag := range tree {
		if bag.Contains(adj, colour) {
			count++
		}
	}
	if count != 4 {
		t.Errorf("Got %d bags containing %s, expected: 4", count, ID(adj, colour))
	}
}
func TestBagTotalChildren(t *testing.T) {
	tree := BuildTree(testInput)
	if children := tree["shiny:gold"].TotalChildren(); children != 32 {
		t.Errorf("Got %d children, expected 32", children)
	}
}
