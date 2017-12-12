package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	input, err := getInput()
	if err != nil {
		panic(err)
	}
	partOne, partTwo := getLocation(input)

	fmt.Println("part one:", partOne)
	fmt.Println("part two:", partTwo)

}

func getLocation(dirs []string) (int, int) {
	x := 0
	y := 0
	z := 0
	max := 0
	distance := 0
	for _, dir := range dirs {
		switch dir {
		case "nn":
			y++
			z--
		case "ne":
			x++
			z--
		case "se":
			y--
			x++
		case "ss":
			z++
			y--
		case "sw":
			x--
			z++
		case "nw":
			y++
			x--
		}

		distance = getDistance(x, y, z)
		if distance > max {
			max = distance
		}

	}
	return distance, max

}

func getDistance(x, y, z int) int {
	distance := (math.Abs(float64(x)) + math.Abs(float64(y)) + math.Abs(float64(z))) / 2
	return int(distance)
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

	input = input[:len(input)-1]
	dirs := strings.Split(string(input), ",")
	return dirs, nil

}
