package main

import (
	"fmt"
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
	rules, messages := ParseInput(input)
	matchers := make(map[string]*Matcher)
	root := NewMatcher("0", rules["0"], rules, matchers)
	matchers["0"] = root
	root.populateSubMatchers()

	count := 0
	for _, msg := range messages {
		if match, remainder := root.Match(msg); match && remainder == "" {
			count++
		}
	}
	return count
}

func PartTwo(input []string) int {
	rules, messages := ParseInput(input)
	matchers := make(map[string]*Matcher)
	rules["8"] = "42 | 42 8"
	rules["11"] = "42 31 | 42 11 31"

	root := NewMatcher("0", rules["0"], rules, matchers)
	matchers["0"] = root
	root.populateSubMatchers()

	count := 0
	for _, msg := range messages {
		if match, remainder := root.Match(msg); match && remainder == "" {
			count++
		}
	}
	return count
}

type Matcher struct {
	rules       map[string]string
	allMatchers map[string]*Matcher

	ruleBody string

	paths [][]*Matcher

	char *string
	key  string
}

func (m *Matcher) Loops() bool {
	return m.key == "8" || m.key == "11"
}
func (m *Matcher) Chars() int {
	if len(m.paths) == 0 {
		return 1
	}
	chars := 0
	for _, subMatcher := range m.paths[0] {
		chars += subMatcher.Chars()
	}
	return chars
}

func (m *Matcher) populateSubMatchers() {
	if strings.Contains(m.ruleBody, `"`) {
		return
	}
	var clauses []string
	if strings.Contains(m.ruleBody, "|") {
		clauses = strings.Split(m.ruleBody, " | ")
	} else {
		clauses = []string{m.ruleBody}
	}
	// each clause is a potential path, if any of the paths match, the overall
	// matcher matches
	m.paths = make([][]*Matcher, len(clauses))
	for c, clause := range clauses {
		m.paths[c] = make([]*Matcher, 0)
		keys := strings.Split(clause, " ")
		for _, key := range keys {
			if matcher, ok := m.allMatchers[key]; ok {
				m.paths[c] = append(m.paths[c], matcher)
				continue
			}
			matcher := NewMatcher(key, m.rules[key], m.rules, m.allMatchers)
			m.allMatchers[key] = matcher
			m.paths[c] = append(m.paths[c], matcher)
			matcher.populateSubMatchers()
		}
	}
}

func (m *Matcher) tryPath(path []*Matcher, s string) (bool, string) {
	cursor := 0
	for i := 0; i < len(path) && len(s) > 0; i++ {
		subMatcher := path[i]
		if cursor+subMatcher.Chars() > len(s) {
			return false, ""
		}

		match, remainder := subMatcher.Match(s)
		if !match {
			return false, ""
		}
		s = remainder
	}
	return true, s
}

func (m *Matcher) Match(s string) (bool, string) {
	if m.char != nil && *m.char == string(s[0]) {
		return true, s[1:]
	}
	for _, path := range m.paths {
		if match, remainder := m.tryPath(path, s); match {
			return true, remainder
		}
	}
	return false, ""
}

func NewMatcher(
	key string,
	ruleBody string,
	rules map[string]string,
	allMatchers map[string]*Matcher,
) *Matcher {
	m := &Matcher{
		rules:       rules,
		allMatchers: allMatchers,
		ruleBody:    ruleBody,
		paths:       make([][]*Matcher, 0),
		char:        nil,
		key:         key,
	}
	if strings.Contains(ruleBody, `"`) {
		// there are no submatchers, this is the most basic matcher
		char := strings.ReplaceAll(ruleBody, `"`, "")
		m.char = &char
	}
	return m
}

func ParseRules(input []string) map[string]string {
	r := make(map[string]string)
	for _, line := range input {
		parts := strings.Split(line, ": ")
		key, ruleBody := parts[0], parts[1]
		r[key] = ruleBody
	}
	return r
}

func ParseInput(input []string) (map[string]string, []string) {
	var messages []string
	var rules map[string]string
	for i, line := range input {
		if line == "" {
			messages = input[i+1:]
			rules = ParseRules(input[:i])
			break
		}
	}
	return rules, messages
}
