// A cardinal sin of package naming
package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GetInput(name string) ([]string, error) {
	f, err := os.Open(name)
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
