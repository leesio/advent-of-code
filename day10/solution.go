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
	suffix := []int{17, 31, 73, 47, 23}
	lengths, err := getLengths(false)
	if err != nil {
		panic(err)
	}
	list := getList()
	partOne := hashList(lengths, list, 1)
	fmt.Println("part one:", partOne[0]*partOne[1])

	lengths, err = getLengths(true)
	lengths = append(lengths, suffix...)
	fmt.Println(lengths)
	if err != nil {
		panic(err)
	}
	list = getList()
	denseHash := getDenseHash(hashList(lengths, list, 64))
	partTwo := ""
	for _, part := range denseHash {
		partTwo += fmt.Sprintf("%02x", part)
	}
	fmt.Println("part two:", partTwo)
}

func getDenseHash(hash []int) []int {
	blockSize := 16
	blocks := len(hash) / blockSize
	denseHash := make([]int, blocks)
	for i := 0; i < blocks; i++ {
		result := hash[i*blockSize]
		for j := 1; j < blockSize; j++ {
			result = result ^ hash[(i*blockSize)+j]
		}
		denseHash[i] = result
	}
	return denseHash
}

func hashList(lengths, list []int, rounds int) []int {
	position := 0
	skipSize := 0

	for r := 0; r < rounds; r++ {
		for _, length := range lengths {
			swapList := make([]int, length)
			subList := make([]int, length)

			for s := range subList {
				subList[s] = list[(position+s)%len(list)]
				swapList[s] = list[(position+s)%len(list)]
			}

			for s := range subList {
				subList[s] = swapList[len(subList)-s-1]
			}

			for l := 0; l < length; l++ {
				list[position] = subList[l]
				position = (position + 1) % len(list)
			}
			position = (position + skipSize) % len(list)
			skipSize++
		}
	}
	return list
}

func getList() []int {
	list := make([]int, 256)
	for i := 0; i < 256; i++ {
		list[i] = i
	}
	return list
}

func getLengths(ascii bool) ([]int, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return []int{}, err
	}
	input, err := ioutil.ReadFile(filepath.Join(cwd, "input"))
	if err != nil {
		return []int{}, err
	}
	input = input[:len(input)-1]
	if ascii {
		asciis := make([]int, len(input))
		for l, char := range input {
			length := int(rune(char))
			asciis[l] = length
		}
		return asciis, nil
	}

	lengthStrs := strings.Split(string(input), ",")
	lengths := make([]int, len(lengthStrs))
	for l, lengthStr := range lengthStrs {
		length, err := strconv.Atoi(lengthStr)
		if err != nil {
			return []int{}, err
		}
		lengths[l] = length
	}
	return lengths, nil

}
