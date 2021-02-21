package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/leesio/advent-of-code/2020/util"
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
	rules, _, nearby := parseInput(input)
	errors := 0
	for _, ticket := range nearby {
		if valid, sum := validTicket(ticket, rules); !valid {
			errors += sum
		}
	}
	return errors
}

func PartTwo(input []string) int {
	rules, myTicket, nearby := parseInput(input)
	candidatesPerField := make([]map[string]bool, len(myTicket))

	for i := range candidatesPerField {
		candidates := make(map[string]bool)
		for name := range rules {
			candidates[name] = true
		}
		candidatesPerField[i] = candidates
	}

	for _, ticket := range append(nearby, myTicket) {
		if valid, _ := validTicket(ticket, rules); !valid {
			continue
		}
		for n, field := range ticket {
			for name, rule := range rules {
				if !rule(field) {
					delete(candidatesPerField[n], name)
				}
			}
		}
	}

	confirmed := make(map[int]bool)
	finalFieldNames := make([]string, len(myTicket))
	for i := 0; i < len(candidatesPerField); {
		field := candidatesPerField[i]
		if len(field) == 1 && !confirmed[i] {
			var fieldName string
			for k := range field {
				fieldName = k
			}
			candidatesPerField = deleteExcept(candidatesPerField, fieldName, i)
			confirmed[i] = true
			finalFieldNames[i] = fieldName
			i = 0
			continue
		}
		i++
	}

	result := 1
	for i, n := range myTicket {
		if strings.HasPrefix(finalFieldNames[i], "departure") {
			result *= n
		}
	}
	return result
}

func deleteExcept(fields []map[string]bool, fieldName string, confirmedPos int) []map[string]bool {
	for f := range fields {
		if f == confirmedPos {
			continue
		}
		delete(fields[f], fieldName)
	}
	return fields
}

func validTicket(ticket []int, r rules) (bool, int) {
	sum := 0
	ticketValid := true
	for _, n := range ticket {
		valid := false
		for _, rule := range r {
			if rule(n) {
				valid = true
				break
			}
		}
		if !valid {
			ticketValid = false
			sum += n
		}
	}
	return ticketValid, sum
}

func getValidator(min, max int) func(int) bool {
	return func(val int) bool {
		return val >= min && val <= max
	}
}

func splitRange(r string) (int, int) {
	parts := strings.Split(r, "-")
	min, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	max, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	return min, max
}

type rules map[string]func(int) bool

func parseRules(lines []string) rules {
	m := make(map[string]func(int) bool)
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		name := parts[0]
		ranges := strings.Split(parts[1], " or ")
		m[name] = func(val int) bool {
			return getValidator(splitRange(ranges[0]))(val) || getValidator(splitRange(ranges[1]))(val)
		}
	}
	return m
}

func parseInput(input []string) (rules, []int, [][]int) {
	sections := make([][]string, 0)
	section := make([]string, 0)
	for _, line := range input {
		if line == "" {
			sections = append(sections, section)
			section = make([]string, 0)
			continue
		}
		section = append(section, line)
	}
	sections = append(sections, section)
	rules := parseRules(sections[0])
	ticket, err := util.Atoi(strings.Split(sections[1][1], ","))
	if err != nil {
		panic(err)
	}
	nearby := make([][]int, 0)
	for _, line := range sections[2][1:] {
		ticket, err := util.Atoi(strings.Split(line, ","))
		if err != nil {
			panic(err)
		}
		nearby = append(nearby, ticket)
	}
	return rules, ticket, nearby
}
