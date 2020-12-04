package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/leesio/advent-of-code/2020/util"
)

var (
	hairColourRegex = regexp.MustCompile("^#[0-9a-f]{6}$")
	pidRegex        = regexp.MustCompile("^[0-9]{9}$")
	validEyeColours = map[string]bool{
		"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true,
	}
	validators = map[string]func(string) bool{
		"byr": func(value string) bool {
			// byr (Birth Year) - four digits; at least 1920 and at most 2002.
			return isNumBetween(value, 1920, 2002)
		},
		"iyr": func(value string) bool {
			// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
			return isNumBetween(value, 2010, 2020)
		},
		"eyr": func(value string) bool {
			// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
			return isNumBetween(value, 2020, 2030)
		},
		"hgt": func(value string) bool {
			// hgt (Height) - a number followed by either cm or in:
			// If cm, the number must be at least 150 and at most 193.
			// If in, the number must be at least 59 and at most 76.
			switch {
			case strings.HasSuffix(value, "cm"):
				return isNumBetween(strings.TrimSuffix(value, "cm"), 150, 193)
			case strings.HasSuffix(value, "in"):
				return isNumBetween(strings.TrimSuffix(value, "in"), 59, 76)
			default:
				return false
			}
		},
		"hcl": func(value string) bool {
			// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
			return hairColourRegex.MatchString(value)
		},
		"ecl": func(value string) bool {
			// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
			return validEyeColours[value]
		},
		"pid": func(value string) bool {
			// pid (Passport ID) - a nine-digit number, including leading zeroes.
			return pidRegex.MatchString(value)
		},
	}
)

func main() {
	input, err := util.GetInput("input")
	if err != nil {
		panic(err)
	}

	passports, err := ParseInput(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("part 1: %d\n", passports.Valid())
	fmt.Printf("part 2: %d\n", passports.StrictlyValid())
}

type Passport map[string]string

func (p Passport) StrictlyValid() bool {
	if !p.Valid() {
		return false
	}
	for field, validator := range validators {
		if val, ok := p[field]; !ok || !validator(val) {
			return false
		}
	}
	return true
}

func (p Passport) Valid() bool {
	if len(p) == 8 {
		return true
	}
	if _, ok := p["cid"]; !ok && len(p) == 7 {
		return true

	}
	return false
}

type Passports []Passport

func (p Passports) Valid() int {
	c := 0
	for _, passport := range p {
		if passport.Valid() {
			c++
		}
	}
	return c
}
func (p Passports) StrictlyValid() int {
	c := 0
	for _, passport := range p {
		if passport.StrictlyValid() {
			c++
		}
	}
	return c
}

func isNumBetween(s string, min, max int) bool {
	y, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	if y < min || y > max {
		return false
	}
	return true
}

func ParseInput(input []string) (Passports, error) {
	passports := make([]Passport, 0)
	p := make(Passport)
	for l, line := range input {
		if line == "" {
			passports = append(passports, p)
			p = make(Passport)
		}
		fields := strings.Fields(line)
		for _, field := range fields {
			parts := strings.Split(field, ":")
			if len(parts) != 2 {
				return nil, fmt.Errorf("unexpected key:value pair: %s in line: %d", field, l)
			}
			p[parts[0]] = parts[1]
		}
	}
	// We'll miss the last passport, 'cos we're only inserting when we encounter
	// an empty line
	passports = append(passports, p)

	return passports, nil
}
