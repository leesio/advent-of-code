package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	lines, err := getInput()
	if err != nil {
		panic(err)
	}

	store := make(Store)
	max, err := store.Populate(lines)
	if err != nil {
		panic(err)
	}
	fmt.Println("part 1:", store.GetMax())
	fmt.Println("part 2:", max)
}

type Store map[string]int

func (s Store) GetMax() int {
	max := 0
	for _, val := range s {
		if val > max {
			max = val
		}
	}
	return max
}

func (s Store) Populate(lines []string) (int, error) {
	max := 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		registerName := parts[0]
		invert := parts[1] == "dec"

		valueStr := parts[2]
		value, err := strconv.Atoi(valueStr)
		if err != nil {
			return 0, err
		}

		condition := parts[4:]

		if invert {
			value = 0 - value
		}

		result, err := s.EvaluateCondition(condition)
		if err != nil {
			return 0, err
		}
		if result {
			s[registerName] += value
		}
		if s[registerName] > max {
			max = s[registerName]
		}
	}
	return max, nil
}

func (s Store) EvaluateCondition(condition []string) (bool, error) {
	register := s[condition[0]]
	operator := condition[1]
	compareTo, err := strconv.Atoi(condition[2])
	if err != nil {
		return false, err
	}

	switch operator {
	case "==":
		return register == compareTo, nil
	case ">=":
		return register >= compareTo, nil
	case ">":
		return register > compareTo, nil
	case "<":
		return register < compareTo, nil
	case "<=":
		return register <= compareTo, nil
	case "!=":
		return register != compareTo, nil
	}
	return false, fmt.Errorf("Unrecognised operator", operator)
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
