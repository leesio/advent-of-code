// A cardinal sin of package naming
package util

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

func ExtractNamedMatches(r *regexp.Regexp, s string) map[string]string {
	names := r.SubexpNames()[1:]
	matches := r.FindStringSubmatch(s)
	m := make(map[string]string)
	for i, match := range matches[1:] {
		m[names[i]] = match
	}
	return m
}
func ExtractNamedSubmatches(r *regexp.Regexp, s string) []map[string]string {
	names := r.SubexpNames()[1:]
	matches := r.FindAllStringSubmatch(s, -1)
	ms := make([]map[string]string, len(matches))
	for i, match := range matches {
		m := make(map[string]string)
		for j, submatch := range match[1:] {
			m[names[j]] = submatch
		}
		ms[i] = m
	}
	return ms
}

func MustAtoi(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return val
}
