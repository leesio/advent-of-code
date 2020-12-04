package main

import (
	"testing"
)

var (
	testInput = []string{
		"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
		"byr:1937 iyr:2017 cid:147 hgt:183cm",
		"",
		"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884",
		"hcl:#cfa07d byr:1929",
		"",
		"hcl:#ae17e1 iyr:2013",
		"eyr:2024",
		"ecl:brn pid:760753108 byr:1931",
		"hgt:179cm",
		"",
		"hcl:#cfa07d eyr:2025 pid:166559648",
		"iyr:2011 ecl:brn hgt:59in",
	}

	strictlyInvalid = []string{
		"eyr:1972 cid:100",
		"hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926",
		"",
		"iyr:2019",
		"hcl:#602927 eyr:1967 hgt:170cm",
		"ecl:grn pid:012533040 byr:1946",
		"",
		"hcl:dab227 iyr:2012",
		"ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277",
		"",
		"hgt:59cm ecl:zzz",
		"eyr:2038 hcl:74454a iyr:2023",
		"pid:3556412378 byr:2007",
	}
	strictlyValid = []string{
		"pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980",
		"hcl:#623a2f",
		"",
		"eyr:2029 ecl:blu cid:129 byr:1989",
		"iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm",
		"",
		"hcl:#888785",
		"hgt:164cm byr:2001 iyr:2015 cid:88",
		"pid:545766238 ecl:hzl",
		"eyr:2022",
		"",
		"iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719",
	}
)

func TestParseInput(t *testing.T) {
	_, err := ParseInput(testInput)
	if err != nil {
		t.Error(err)
	}
}
func TestValid(t *testing.T) {
	p, err := ParseInput(testInput)
	if err != nil {
		t.Error(err)
	}
	if v := p.Valid(); v != 2 {
		t.Errorf("Got %d valid passports, expected: %d", v, 2)
	}
}

func TestStrictlyValid(t *testing.T) {
	p, err := ParseInput(strictlyValid)
	if err != nil {
		t.Error(err)
	}
	if v := p.StrictlyValid(); v != len(p) {
		t.Errorf("Got %d strictly valid passports, expected: %d", v, len(p))
	}
}
func TestStrictlyInValid(t *testing.T) {
	p, err := ParseInput(strictlyInvalid)
	if err != nil {
		t.Error(err)
	}
	if v := p.StrictlyValid(); v != 0 {
		t.Errorf("Got %d strictly valid passports, expected: %d", v, 0)
	}
}

func TestValidators(t *testing.T) {
	cases := []struct {
		key   string
		valid bool
		val   string
	}{
		{"byr", true, "2002"},
		{"byr", false, "2003"},
		{"hgt", true, "60in"},
		{"hgt", true, "190cm"},
		{"hgt", false, "190in"},
		{"hgt", false, "190"},
		{"hcl", true, "#123abc"},
		{"hcl", false, "#123abz"},
		{"hcl", false, "123abc"},
		{"ecl", true, "brn"},
		{"ecl", false, "wat"},
		{"pid", true, "000000001"},
		{"pid", false, "0123456789"},
	}
	for _, c := range cases {
		if res := validators[c.key](c.val); res != c.valid {
			t.Errorf("got: %v, expected: %v for key: %s, val: %s", res, c.valid, c.key, c.val)
		}
	}
}
