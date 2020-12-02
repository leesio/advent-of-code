package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/leesio/advent-of-code/2020/util"
)

type Item struct {
	Password string
	Policy   *PasswordPolicy
}

func (i *Item) Valid() bool {
	return i.Policy.PasswordValid(i.Password)
}
func (i *Item) ValidOTCA() bool {
	return i.Policy.PasswordValidOTCA(i.Password)
}

type PasswordPolicy struct {
	MinOccurences int
	MaxOccurences int
	Pattern       string
}

func (p *PasswordPolicy) PasswordValidOTCA(pwd string) bool {
	if len(pwd) < p.MinOccurences || len(pwd) < p.MaxOccurences {
		return false
	}
	count := 0
	if strings.Split(pwd, "")[p.MinOccurences-1] == p.Pattern {
		count++
	}
	if strings.Split(pwd, "")[p.MaxOccurences-1] == p.Pattern {
		count++
	}
	return count == 1
}
func (p *PasswordPolicy) PasswordValid(pwd string) bool {
	occurences := strings.Count(pwd, p.Pattern)
	return occurences >= p.MinOccurences && occurences <= p.MaxOccurences
}

func parseMinMax(s string) (int, int, error) {
	limitParts := strings.Split(s, "-")
	if len(limitParts) != 2 {
		return 0, 0, fmt.Errorf("unexpected limit format type: %s", s)
	}
	min, err := strconv.Atoi(limitParts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("unable to parse lower limit as int: %s in %s", err, s)
	}
	max, err := strconv.Atoi(limitParts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("unable to parse upper limit as int: %s in %s", err, s)
	}
	return min, max, nil
}
func parseInput(input []string) ([]*Item, error) {
	items := make([]*Item, len(input))
	for n, line := range input {
		parts := strings.Fields(line)
		if len(parts) != 3 {
			return nil, fmt.Errorf("unexpected input format: %s", line)
		}
		min, max, err := parseMinMax(parts[0])
		if err != nil {
			return nil, fmt.Errorf("error parsing input limits: %s in %s", err, line)
		}
		ch := strings.TrimSuffix(parts[1], ":")
		items[n] = &Item{
			Password: parts[2],
			Policy: &PasswordPolicy{
				MinOccurences: min,
				MaxOccurences: max,
				Pattern:       ch,
			},
		}

	}
	return items, nil
}

func main() {
	input, err := util.GetInput("input")
	if err != nil {
		panic(err)
	}
	items, err := parseInput(input)
	if err != nil {
		panic(err)
	}
	valid := 0
	validOTCA := 0
	for _, item := range items {
		if item.Valid() {
			valid++
		}
		if item.ValidOTCA() {
			validOTCA++
		}
	}
	fmt.Printf("part one: %d (%d)\n", valid, len(items))
	fmt.Printf("part two: %d (%d)\n", validOTCA, len(items))
}
