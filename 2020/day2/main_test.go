package main

import (
	"testing"
)

var testInput = []string{
	"1-3 a: abcde",
	"1-3 b: cdefg",
	"2-9 c: ccccccccc",
}

var testItems = []struct {
	Item       *Item
	Result     bool
	ResultOTCA bool
}{
	{
		Item: &Item{
			Password: "abcde",
			Policy:   &PasswordPolicy{MinOccurences: 1, MaxOccurences: 3, Pattern: "a"},
		},
		Result:     true,
		ResultOTCA: true,
	},
	{
		Item: &Item{
			Password: "cdefg",
			Policy:   &PasswordPolicy{MinOccurences: 1, MaxOccurences: 3, Pattern: "b"},
		},
		Result:     false,
		ResultOTCA: false,
	},
	{
		Item: &Item{
			Password: "ccccccccc",
			Policy:   &PasswordPolicy{MinOccurences: 2, MaxOccurences: 9, Pattern: "c"},
		},
		Result:     true,
		ResultOTCA: false,
	},
}

func TestParseInput(t *testing.T) {
	items, err := parseInput(testInput)
	if err != nil {
		t.Fatal(err)
	}
	if len(items) != len(testInput) {
		t.Fatal("The simplest comparison possible failed")
	}
	for n, parsedItem := range items {
		testItem := testItems[n].Item
		if parsedItem.Password != testItem.Password {
			t.Fail()
		}
		if parsedItem.Policy.MinOccurences != testItem.Policy.MinOccurences {
			t.Fail()
		}
		if parsedItem.Policy.MaxOccurences != testItem.Policy.MaxOccurences {
			t.Fail()
		}
		if parsedItem.Policy.Pattern != testItem.Policy.Pattern {
			t.Fail()
		}
	}
}
func TestValidation(t *testing.T) {
	for i := 0; i < len(testItems); i++ {
		testCase := testItems[i]
		if testCase.Item.Valid() != testCase.Result {
			t.Fatal("Case", i, "failed")
		}
	}
}
func TestValidationOTCA(t *testing.T) {
	for i := 0; i < len(testItems); i++ {
		testCase := testItems[i]
		if testCase.Item.ValidOTCA() != testCase.ResultOTCA {
			t.Fatal("Case", i, "failed")
		}
	}
}
