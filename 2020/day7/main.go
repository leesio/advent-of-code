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
	bags := BuildTree(input)
	count := 0
	for _, bag := range bags {
		if bag.Contains("shiny", "gold") {
			count++
		}
	}
	fmt.Printf("part 1: %d\n", count)
	fmt.Printf("part 2: %d\n", bags["shiny:gold"].TotalChildren())
}

type Tree map[string]*Bag

func (t Tree) Get(adj, colour string) *Bag {
	if b, ok := t[ID(adj, colour)]; ok {
		return b
	}
	bag := CreateBag(adj, colour)
	t[bag.ID] = bag
	return bag
}

type ChildBag struct {
	*Bag
	Count int
}

type Bag struct {
	ID        string
	Adjective string
	Colour    string
	Children  []*ChildBag
}

func (b *Bag) Contains(adj, colour string) bool {
	id := ID(adj, colour)
	for _, child := range b.Children {
		if child.ID == id || child.Contains(adj, colour) {
			return true
		}
	}
	return false
}

func (b *Bag) TotalChildren() int {
	total := 0
	for _, child := range b.Children {
		total = total + child.Count + (child.Count * child.TotalChildren())
	}
	return total
}

func BuildTree(input []string) map[string]*Bag {
	tree := make(Tree)
	for _, line := range input {
		line = strings.Trim(line, ".")
		parts := strings.Split(line, " contain ")
		root := tree.Get(extractParentData(parts[0]))
		children := strings.Split(parts[1], ", ")
		if len(children) == 1 && children[0] == "no other bags" {
			root.Children = make([]*ChildBag, 0)
			continue
		}
		root.Children = make([]*ChildBag, len(children))
		for c, child := range children {
			adj, colour, count := extractChildData(child)
			root.Children[c] = &ChildBag{
				Count: count,
				Bag:   tree.Get(adj, colour),
			}
		}
	}
	return tree
}

func ID(adj, colour string) string {
	return fmt.Sprintf("%s:%s", adj, colour)
}

func extractChildData(raw string) (string, string, int) {
	parts := strings.Split(raw, " ")
	if len(parts) != 4 {
		panic(fmt.Errorf("Got a child with %d parts, expected: 4: %v", len(parts), parts))
	}
	count, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(fmt.Errorf("error converting child count to int: %s", err))
	}
	return parts[1], parts[2], count
}

func extractParentData(parent string) (string, string) {
	parts := strings.Split(parent, " ")
	if len(parts) != 3 {
		panic(fmt.Errorf("Got a bag description with %d parts, expected: 3: %s", len(parts), parent))
	}
	return parts[0], parts[1]
}

func CreateBag(adj, colour string) *Bag {
	return &Bag{
		Adjective: adj,
		Colour:    colour,
		ID:        ID(adj, colour),
	}
}
