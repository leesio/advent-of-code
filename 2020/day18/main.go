package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/leesio/advent-of-code/2020/util"
)

const (
	add      = "+"
	multiply = "*"
)

var (
	parenRegexp    = regexp.MustCompile("\\([0-9 +*]+\\)")
	additionRegexp = regexp.MustCompile("\\d+ \\+ \\d+")
)

func main() {
	input, err := util.GetInput("input")
	if err != nil {
		panic(err)
	}
	fmt.Printf("part 1: %d\n", PartOne(input))
	fmt.Printf("part 2: %d\n", PartTwo(input))
}

func PartOne(input []string) int {
	sum := 0
	for _, line := range input {
		sum += Evaluate(line)
	}
	return sum
}
func PartTwo(input []string) int {
	sum := 0
	for _, line := range input {
		sum += EvaluateAdvanced(line)
	}
	return sum
}

func evaluatePair(a int, operator string, b int) int {
	if operator == add {
		return a + b
	}
	return a * b
}

func collectParens(s string, evaluateFn func(string) int) string {
	collapsed := s
	for parenRegexp.MatchString(collapsed) {
		collapsed = parenRegexp.ReplaceAllStringFunc(collapsed, func(s string) string {
			return strconv.Itoa(evaluateFn(
				strings.Replace(strings.Replace(s, "(", "", -1), ")", "", -1),
			))
		})
	}
	return collapsed
}

func Evaluate(s string) int {
	s = collectParens(s, Evaluate)
	parts := strings.Split(s, " ")
	res := util.MustAtoi(parts[0])
	for i := 1; i < len(parts); i += 2 {
		operator := parts[i]
		b := util.MustAtoi(parts[i+1])
		res = evaluatePair(res, operator, b)
	}
	return res
}

func EvaluateAdvanced(s string) int {
	collapsed := collectParens(s, EvaluateAdvanced)
	for additionRegexp.MatchString(collapsed) {
		collapsed = additionRegexp.ReplaceAllStringFunc(collapsed, func(s string) string {
			parts := strings.Split(s, " ")
			as, operator, bs := parts[0], parts[1], parts[2]
			return strconv.Itoa(
				evaluatePair(util.MustAtoi(as), operator, util.MustAtoi(bs)),
			)
		})
	}
	return Evaluate(collapsed)
}
