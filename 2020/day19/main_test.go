package main

import (
	"fmt"
	"testing"
)

var testInput = []string{
	"0: 4 1 5",
	"1: 2 3 | 3 2",
	"2: 4 4 | 5 5",
	"3: 4 5 | 5 4",
	`4: "a"`,
	`5: "b"`,
	"",
	"ababbb",
	"bababa",
	"abbbab",
	"aaabbb",
	"aaaabbb",
}
var testInput2 = []string{
	"42: 9 14 | 10 1",
	"9: 14 27 | 1 26",
	"10: 23 14 | 28 1",
	`1: "a"`,
	"11: 42 31",
	"5: 1 14 | 15 1",
	"19: 14 1 | 14 14",
	"12: 24 14 | 19 1",
	"16: 15 1 | 14 14",
	"31: 14 17 | 1 13",
	"6: 14 14 | 1 14",
	"2: 1 24 | 14 4",
	"0: 8 11",
	"13: 14 3 | 1 12",
	"15: 1 | 14",
	"17: 14 2 | 1 7",
	"23: 25 1 | 22 14",
	"28: 16 1",
	"4: 1 1",
	"20: 14 14 | 1 15",
	"3: 5 14 | 16 1",
	"27: 1 6 | 14 18",
	`14: "b"`,
	"21: 14 1 | 1 14",
	"25: 1 1 | 1 14",
	"22: 14 14",
	"8: 42",
	"26: 14 22 | 1 20",
	"18: 15 15",
	"7: 14 5 | 1 21",
	"24: 14 1",
	"",
	"abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa",
	"bbabbbbaabaabba",
	"babbbbaabbbbbabbbbbbaabaaabaaa",
	"aaabbbbbbaaaabaababaabababbabaaabbababababaaa",
	"bbbbbbbaaaabbbbaaabbabaaa",
	"bbbababbbbaaaaaaaabbababaaababaabab",
	"ababaaaaaabaaab",
	"ababaaaaabbbaba",
	"baabbaaaabbaaaababbaababb",
	"abbbbabbbbaaaababbbbbbaaaababb",
	"aaaaabbaabaaaaababaa",
	"aaaabbaaaabbaaa",
	"aaaabbaabbaaaaaaabbbabbbaaabbaabaaa",
	"babaaabbbaaabaababbaabababaaab",
	"aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba",
}

func TestPartOne(t *testing.T) {
	rules, messages := ParseInput(testInput)
	matchers := make(map[string]*Matcher)
	root := NewMatcher("0", rules["0"], rules, matchers)
	matchers["0"] = root
	root.populateSubMatchers()
	for key, matcher := range matchers {
		fmt.Println(key, matcher.key)
	}

	for _, msg := range messages {
		match, remainder := root.Match(msg)
		fmt.Println(msg, match, remainder)
	}
}
func TestPartTwo(t *testing.T) {
	rules, _ := ParseInput(testInput2)
	rules["8"] = "42 | 42 8"
	rules["11"] = "42 31 | 42 11 31"
	matchers := make(map[string]*Matcher)
	root := NewMatcher("0", rules["0"], rules, matchers)
	matchers["0"] = root
	root.populateSubMatchers()

	// for _, msg := range messages {
	// 	fmt.Println(msg, root.Match(msg))
	// }
	lol := "bbabbbbaabaabba"
	match, remainder := root.Match(lol)
	fmt.Println(lol, match, remainder)
}
