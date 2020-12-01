package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getInput() ([]string, error) {
	f, err := os.Open("input")
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(f)
	vals := make([]string, 0)
	for scanner.Scan() {
		vals = append(vals, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return vals, nil
}

func Atoi(ss []string) ([]int, error) {
	vals := make([]int, len(ss))
	for i, s := range ss {
		val, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("error casting input (%s) to int %s", s, err)
		}
		vals[i] = val
	}
	return vals, nil
}

func partOne(input []int) {
	m := make(map[int]struct{})
	for _, num := range input {
		m[num] = struct{}{}
		if _, ok := m[2020-num]; ok {
			fmt.Println(num * (2020 - num))
			break
		}
	}
}
func partTwo(input []int) {
	h := make(map[int]struct{})
	for _, n := range input {
		for _, m := range input {
			h[n] = struct{}{}
			if _, ok := h[2020-m-n]; ok {
				fmt.Println(m * n * (2020 - m - n))
				return
			}
		}
	}
}
func main() {
	input, err := getInput()
	if err != nil {
		panic(err)
	}
	inputNums, err := Atoi(input)
	if err != nil {
		panic(err)
	}
	partOne(inputNums)
	partTwo(inputNums)
}
