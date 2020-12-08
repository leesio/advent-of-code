package main

import (
	"fmt"
	"time"
)

type ExecutionContext struct {
	instructions     Instructions
	cursor           int
	acc              int
	seenInstructions map[string]bool
	err              error
}

func NewExecutionContext(instructions Instructions) *ExecutionContext {
	return &ExecutionContext{
		instructions:     instructions,
		cursor:           0,
		acc:              0,
		seenInstructions: make(map[string]bool),
		err:              nil,
	}
}

func (e *ExecutionContext) cloneSeenInstructions() map[string]bool {
	m := make(map[string]bool)
	for k, v := range e.seenInstructions {
		m[k] = v
	}
	return m
}

func (e *ExecutionContext) Clone() *ExecutionContext {
	return &ExecutionContext{
		instructions:     e.instructions.Clone(),
		cursor:           e.cursor,
		seenInstructions: e.cloneSeenInstructions(),
		acc:              e.acc,
	}
}

func (e *ExecutionContext) Error() error {
	return e.err
}

func (e *ExecutionContext) NextInstruction() *Instruction {
	if e.cursor == len(e.instructions) {
		return nil
	}
	if e.cursor < 0 || e.cursor > len(e.instructions) {
		e.err = fmt.Errorf("invalidl cursor: %d", e.cursor)
		return nil
	}
	return e.instructions[e.cursor]
}

func (e *ExecutionContext) Execute() (int, error) {
	for e.Step() {
	}
	return e.acc, e.Error()
}

func (e *ExecutionContext) Step() bool {
	instruction := e.NextInstruction()
	if instruction == nil {
		return false
	}
	if e.seenInstructions[instruction.ID] {
		e.err = fmt.Errorf("started infinite loop at instruction: %d", e.cursor)
		return false
	}
	e.seenInstructions[instruction.ID] = true
	switch instruction.operation {
	case JMP:
		e.cursor = e.cursor + instruction.argument
		return true
	case ACC:
		e.acc = e.acc + instruction.argument
	case NOP:
	}
	e.cursor++
	return true
}

func (e ExecutionContext) FindBrokenInstruction() int {
	resultCh := make(chan int)
	go func() {
		for inst := e.NextInstruction(); inst != nil; inst = e.NextInstruction() {
			if inst.operation == JMP || inst.operation == NOP {
				alt := e.Clone()
				newInstruction := inst.Clone()
				if inst.operation == JMP {
					newInstruction.operation = NOP
				} else {
					newInstruction.operation = JMP
				}
				alt.instructions[alt.cursor] = newInstruction
				go func(resultCh chan int) {
					acc, err := alt.Execute()
					if err == nil {
						resultCh <- acc
					}
				}(resultCh)
			}
			e.Step()
		}
	}()
	select {
	case res := <-resultCh:
		return res
	case <-time.After(5 * time.Second):
		return -1
	}
}
