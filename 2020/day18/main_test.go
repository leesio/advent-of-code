package main

import (
	"testing"
)

var (
	testInputs = []struct {
		input       string
		exp         int
		expAdvanced int
	}{
		{input: "1 + 2 * 3 + 4 * 5 + 6", exp: 71, expAdvanced: 231},
		{input: "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", exp: 13632, expAdvanced: 23340},
	}
)

func TestEvaluate(t *testing.T) {
	for _, testInput := range testInputs {
		if res := Evaluate(testInput.input); res != testInput.exp {
			t.Errorf("Got %d, expected: %d", res, testInput.exp)
		}
	}
}
func TestEvaluateAdvanced(t *testing.T) {
	for _, testInput := range testInputs {
		if res := EvaluateAdvanced(testInput.input); res != testInput.expAdvanced {
			t.Errorf("Got %d, expected: %d", res, testInput.expAdvanced)
		}
	}
}
