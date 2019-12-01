package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type Node struct {
	id           int
	subNodeNames []int
}

type Tree map[int]Node

func main() {
	input, err := getInput()
	if err != nil {
		panic(err)
	}

	tree := make(Tree, 0)
	err = tree.PopulateTree(input)
	seen := make(map[int]bool)
	partOne := len(tree.GetBranch(0, seen))
	partTwo := tree.GetGroupCount()
	fmt.Println("part one:", partOne)
	fmt.Println("part two:", partTwo)
}

func (t Tree) GetGroupCount() int {
	count := 0
	for b := range t {
		seen := make(map[int]bool)
		branch := t.GetBranch(b, seen)
		for x := range branch {
			delete(t, x)
		}
		count++
	}
	return count
}

func (t Tree) GetBranch(id int, seen map[int]bool) map[int]bool {
	startNode := t[id]
	seen[id] = true

	for _, subNodeID := range startNode.subNodeNames {
		if seen[subNodeID] {
			continue
		}
		seen[subNodeID] = true
		t.GetBranch(subNodeID, seen)
	}
	return seen
}

func (t Tree) PopulateTree(lines []string) error {
	idReg := regexp.MustCompile("[0-9]{1,4}")

	for _, line := range lines {
		ids := idReg.FindAllString(line, -1)
		idStr := ids[0]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return err
		}

		subNodes := ids[1:]
		node := Node{
			id:           id,
			subNodeNames: make([]int, len(subNodes)),
		}

		for s, subNode := range subNodes {
			subNodeId, err := strconv.Atoi(subNode)
			if err != nil {
				return err
			}
			node.subNodeNames[s] = subNodeId
		}
		t[id] = node
	}
	return nil
}

func getInput() ([]string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return []string{}, err
	}
	input, err := ioutil.ReadFile(filepath.Join(cwd, "input"))
	if err != nil {
		return []string{}, err
	}
	lines := strings.Split(string(input), "\n")
	lines = lines[:len(lines)-1]

	return lines, nil
}
