package main

import (
	"fmt"
	"testing"
)

var testInput = []string{
	"nop +0",
	"acc +1",
	"jmp +4",
	"acc +3",
	"jmp -3",
	"acc -99",
	"acc +1",
	"jmp -4",
	"acc +6",
}
var testArgs = []int{0, 1, 4, 3, -3, -99, 1, -4, 6}

func TestExtractInstructions(t *testing.T) {
	instructions, err := ExtractInstructions(testInput)
	if err != nil {
		t.Error(err)
	}
	if len(instructions) != len(testInput) {
		t.Errorf("Got %d instructions, expected: %d", len(instructions), len(testInput))
	}
	for i, in := range instructions {
		if testArgs[i] != in.argument {
			t.Errorf("Case: %d - got instructions with arg: %d, expected: %d", i, in.argument, testArgs[i])
		}
	}
}
func TestExecute(t *testing.T) {
	instructions, err := ExtractInstructions(testInput)
	if err != nil {
		t.Error(fmt.Errorf("error extracting instructions: %s", err))
	}
	firstLoop, err := Execute(instructions)
	if err == nil {
		t.Errorf("Expected error due to infinite loop")
	}
	if firstLoop != 5 {
		t.Errorf("Got first loop at %d, expected: %d", firstLoop, 5)
	}
}
func TestBrokenInstruction(t *testing.T) {
	instructions, err := ExtractInstructions(testInput)
	if err != nil {
		t.Error(fmt.Errorf("error extracting instructions: %s", err))
	}
	instructions, err = FixBrokenInstruction(instructions)
	if err != nil {
		t.Errorf("got error looking for broken instruction: %s", err)
	}
	correctResult, err := Execute(instructions)
	if err != nil {
		t.Errorf("got error running fixed instruction: %s", err)
	}
	if correctResult != 8 {
		t.Errorf("Got right answer of %d, expected: %d", correctResult, 8)
	}
}
