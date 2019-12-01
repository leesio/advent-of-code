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
	input, err := getInput()
	if err != nil {
		panic(err)
	}
	inputOne := make([]int, len(input))
	inputTwo := make([]int, len(input))

	copy(inputOne, input)
	copy(inputTwo, input)

	answerOne := getAnswer(inputOne, func(current int) int {
		return 1
	})

	answerTwo := getAnswer(inputTwo, func(current int) int {
		if current >= 3 {
			return -1
		}
		return 1
	})

	fmt.Printf("part one: %v\n", answerOne)
	fmt.Printf("part two: %v\n", answerTwo)
}

func getAnswer(input []int, getInc func(current int) int) int {
	pointer := 0
	steps := 0
	for {
		if pointer >= len(input) {
			return steps
		}
		current := input[pointer]
		inc := getInc(current)
		input[pointer] = input[pointer] + inc
		pointer += current
		steps++
		continue
	}
}

func getInput() ([]int, error) {
	vals := make([]int, 0)
	cwd, err := os.Getwd()
	if err != nil {
		return vals, err
	}
	input, err := ioutil.ReadFile(filepath.Join(cwd, "input"))
	if err != nil {
		return vals, err
	}
	chars := strings.Split(string(input), "\n")
	for _, char := range chars {
		if char == "" {
			continue
		}
		val, err := strconv.Atoi(char)
		if err != nil {
			return []int{}, nil
		}
		vals = append(vals, val)
	}
	return vals, nil
}
