package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Object struct {
	Name     string
	Orbiting []*Object
	Orbiters []*Object
}

func main() {
	f, err := os.Open("./input")
	if err != nil {
		fmt.Println("error opening file", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("error reading file", err)
		os.Exit(1)
	}
	orbits := strings.Split(strings.TrimRight(string(b), "\n"), ("\n"))
	fmt.Println(partOne(orbits))
	fmt.Println(partTwo(orbits))
}
func partOne(orbits []string) int {
	tree := buildTree(orbits)
	count := 0
	for _, leaf := range tree {
		count = count + countOrbits(leaf)
	}
	return count
}

func partTwo(orbits []string) int {
	tree := buildTree(orbits)
	san, you := tree["SAN"], tree["YOU"]
	sanPaths := plotPaths(san, 0, make(map[string]int))
	youPaths := plotPaths(you, 0, make(map[string]int))

	shortestPath := -1
	for key, sanPath := range sanPaths {
		if youPath, ok := youPaths[key]; ok {
			// this means we have a common ancestor
			// sum the paths and we get a route
			if shortestPath == -1 || shortestPath > (sanPath+youPath) {
				shortestPath = sanPath + youPath
			}
		}
	}
	return shortestPath
}

func buildTree(orbits []string) map[string]*Object {
	tree := make(map[string]*Object)
	for _, orbit := range orbits {
		parts := strings.Split(orbit, (")"))
		orbiteeName, orbiterName := parts[0], parts[1]
		var orbiter, orbitee *Object
		if entry, ok := tree[orbiterName]; ok {
			orbiter = entry
		} else {
			orbiter = &Object{
				Name:     orbiterName,
				Orbiting: make([]*Object, 0),
				Orbiters: make([]*Object, 0),
			}
			tree[orbiter.Name] = orbiter
		}
		if entry, ok := tree[orbiteeName]; ok {
			orbitee = entry
		} else {
			orbitee = &Object{
				Name:     orbiteeName,
				Orbiting: make([]*Object, 0),
				Orbiters: make([]*Object, 0),
			}
			tree[orbitee.Name] = orbitee
		}
		orbitee.Orbiters = append(orbitee.Orbiters, orbiter)
		orbiter.Orbiting = append(orbiter.Orbiting, orbitee)
	}
	return tree
}

func plotPaths(object *Object, n int, tree map[string]int) map[string]int {
	for _, parent := range object.Orbiting {
		if currentVal, ok := tree[parent.Name]; ok {
			if currentVal < n {
				tree[parent.Name] = n
			}
		} else {
			tree[parent.Name] = n
		}
		plotPaths(parent, n+1, tree)
	}
	return tree
}

func countOrbits(object *Object) int {
	if len(object.Orbiting) == 0 {
		return 0
	}
	count := len(object.Orbiting)
	for _, c := range object.Orbiting {
		l := countOrbits(c)
		count = count + l
	}
	return count
}
