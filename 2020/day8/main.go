package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/leesio/advent-of-code/2020/util"
)

const (
	JMP Operation = "jmp"
	ACC Operation = "acc"
	NOP Operation = "nop"
)

func main() {
	input, err := util.GetInput("input")
	if err != nil {
		panic(err)
	}
	instructions, err := ExtractInstructions(input)
	if err != nil {
		panic(err)
	}
	acc, err := Execute(instructions)
	if err != nil {
		fmt.Printf("part 1: %d\n", acc)
	} else {
		panic(fmt.Errorf("didn't find first loop"))
	}
	a := time.Now()
	fixedInstructions, err := FixBrokenInstruction(instructions)
	if err != nil {
		panic(err)
	}
	acc, err = Execute(fixedInstructions)
	if err != nil {
		panic(fmt.Errorf("fixed instructions failed to execute: %s", err))
	}
	fmt.Printf("part 2 (1): %d (%v)\n", acc, time.Now().Sub(a))

	a = time.Now()
	e := NewExecutionContext(instructions)
	fmt.Printf("part 2 (2): %d (%v)\n", e.FindBrokenInstruction(), time.Now().Sub(a))
}

type Operation string

func (o Operation) Valid() bool {
	switch o {
	case JMP:
	case ACC:
	case NOP:
	default:
		return false
	}
	return true
}

type Instruction struct {
	ID        string
	operation Operation
	argument  int
}

func (i *Instruction) Clone() *Instruction {
	return &Instruction{
		ID:        i.ID,
		operation: i.operation,
		argument:  i.argument,
	}
}

type Instructions []*Instruction

func (in Instructions) Clone() Instructions {
	clone := make(Instructions, len(in))
	for i, instruction := range in {
		clone[i] = instruction.Clone()
	}
	return clone
}

func ExtractInstructions(input []string) (Instructions, error) {
	instructions := make(Instructions, len(input))
	for i, line := range input {
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			return nil, fmt.Errorf("unexpected input structure: %s", line)
		}
		argument, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		operation := Operation(parts[0])
		if !operation.Valid() {
			return nil, fmt.Errorf("unknown operation: %s", parts[0])
		}
		instructions[i] = &Instruction{
			operation: operation,
			argument:  argument,
			ID:        strconv.Itoa(i),
		}
	}
	return instructions, nil
}

func FixBrokenInstruction(instructions Instructions) (Instructions, error) {
	for cursor := 0; cursor < len(instructions); {
		inst := instructions[cursor]
		switch inst.operation {
		case JMP:
			altInstructions := instructions.Clone()
			altInstructions[cursor].operation = NOP
			if _, err := Execute(altInstructions); err == nil {
				return altInstructions, nil
			}
		case NOP:
			altInstructions := instructions.Clone()
			altInstructions[cursor].operation = JMP
			if _, err := Execute(altInstructions); err == nil {
				return altInstructions, nil
			}
		case ACC:
		}
		cursor++
	}
	return nil, fmt.Errorf("couldn't find broken instruction")
}

func Execute(instructions Instructions) (int, error) {
	seenInstructions := make(map[string]bool)
	acc := 0
	for cursor := 0; cursor < len(instructions); {
		inst := instructions[cursor]
		if seenInstructions[inst.ID] {
			return acc, fmt.Errorf("started infinite loop at instruction: %d", cursor)
		}
		seenInstructions[inst.ID] = true
		switch inst.operation {
		case JMP:
			cursor = cursor + inst.argument
			continue
		case ACC:
			acc = acc + inst.argument
		case NOP:
		}
		cursor++
	}
	return acc, nil
}
