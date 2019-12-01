package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	file, err := getInputFile()
	if err != nil {
		panic(err)
	}
	score, garbageCount, err := parseFile(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 1:", score)
	fmt.Println("Part 2:", garbageCount)

}

func parseFile(file *os.File) (int, int, error) {
	reader := bufio.NewReader(file)

	groupValue := 0
	score := 0
	currentlyIgnoring := false
	negateNextChar := false
	garbageCount := 0

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return score, garbageCount, nil
			}
			return 0, 0, err
		}

		char := string(r)
		if negateNextChar {
			negateNextChar = false
			continue
		}

		if currentlyIgnoring && char != ">" && char != "!" {
			garbageCount++
			continue
		}

		switch char {
		case "{":
			groupValue++
		case "}":
			score += groupValue
			groupValue--
		case "<":
			currentlyIgnoring = true
		case ">":
			currentlyIgnoring = false
		case "!":
			negateNextChar = true
		}
	}
	return 0, 0, fmt.Errorf("Inexplicably reached the end of the parse function")
}

func getInputFile() (*os.File, error) {
	var f *os.File
	cwd, err := os.Getwd()
	if err != nil {
		return f, err
	}
	return os.Open(filepath.Join(cwd, "input"))
}
