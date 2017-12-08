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

var test map[string]Node

type Tree map[string]Node

type Node struct {
	name         string
	weight       int
	subNodeNames []string
	subNodes     []Node
	parent       string
}

func (t Tree) BranchMatch(path string) bool {
	pathSums := make(map[int][]string)
	node := t[path]
	for _, subNodeName := range node.subNodeNames {
		sum := t.Sum(subNodeName)
		if _, ok := pathSums[sum]; ok {
			pathSums[sum] = append(pathSums[sum], subNodeName)
		} else {
			pathSums[sum] = []string{subNodeName}
		}
	}

	if len(pathSums) > 1 {
		// if there are multiple path sums, the branch is unbalanced
		for sum, pathSlice := range pathSums {
			if len(pathSlice) == 1 {
				unbalancedBranchPath := pathSlice[0]
				subsMatch := t.BranchMatch(unbalancedBranchPath)
				if subsMatch {
					// If all the branch's subNodes are balanced, this node must have the
					// wrong weight, check what it SHOULD be
					unbalancedNode := t[unbalancedBranchPath]
					for otherSum := range pathSums {
						if otherSum != sum {
							delta := otherSum - sum
							// this is the correct weight
							fmt.Println(
								"The broken node is",
								unbalancedNode.name,
								"It weighs", unbalancedNode.weight, "and should weigh",
								unbalancedNode.weight+delta,
							)
						}
					}
				}
			}
		}
		return false
	}
	return true
}

func (t Tree) Sum(path string) int {
	if node, ok := t[path]; ok {
		var sum int
		sum += node.weight
		for _, s := range node.subNodeNames {
			sum += t.Sum(s)
		}
		return sum
	}
	panic(fmt.Errorf("Trying to sum absent path: %v\n", path))
}

func (t Tree) GetBaseNode() Node {
	subNodeMap := make(map[string]bool)
	for key := range t {
		for _, sub := range t[key].subNodeNames {
			subNodeMap[sub] = true
		}
	}

	for key := range t {
		if _, ok := subNodeMap[key]; !ok {
			return t[key]
		}
	}
	return Node{}
}

func (t Tree) PopulateTree(lines []string) {
	nameReg := regexp.MustCompile("[a-z]{1,}")
	weightReg := regexp.MustCompile("[0-9]{1,}")

	for _, line := range lines {
		names := nameReg.FindAllString(line, -1)
		name := names[0]
		weightStr := weightReg.FindString(line)

		weight, err := strconv.Atoi(weightStr)
		if err != nil {
			panic(err)
		}

		node := Node{
			name:         name,
			weight:       weight,
			subNodeNames: make([]string, 0),
		}

		subNodes := names[1:]
		for _, subNode := range subNodes {
			node.subNodeNames = append(node.subNodeNames, subNode)
		}
		t[name] = node
	}
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

func main() {
	test = make(map[string]Node)

	input, err := getInput()
	if err != nil {
		panic(err)
	}

	tree := make(Tree, 0)
	tree.PopulateTree(input)
	baseNode := tree.GetBaseNode()
	fmt.Println("Base node", baseNode.name)
	tree.BranchMatch(baseNode.name)
}
