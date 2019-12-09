package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	POSITIONAL = iota
	IMMEDIATE
)

func main() {
	f, err := os.Open("./input")
	if err != nil {
		fmt.Println("error opening file", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("error reading file", err)
		os.Exit(1)
	}
	s := strings.Split(strings.Trim(string(b), "\n "), ",")
	run(s, 0)(1)
	run(s, 0)(5)
}
func run(baseList []string, pos int) func(input int) {
	list := make([]string, len(baseList))
	copy(list, baseList)
	for pos < len(list) {
		instruction, err := parseInstruction(list[pos])
		if err != nil {
			fmt.Println("error parsing instruction", err)
			os.Exit(1)
		}
		paramCount := instruction.ParamCount()
		params := list[pos+1 : pos+paramCount+1]
		debugf("%d params: %s (opcode %d)\n", paramCount, params, instruction.Opcode)
		switch instruction.Opcode {
		case 1:
			a := stringToInt(params[0])
			if instruction.ParamType(0) == POSITIONAL {
				a = stringToInt(list[a])
			}
			b := stringToInt(params[1])
			if instruction.ParamType(1) == POSITIONAL {
				b = stringToInt(list[b])
			}
			idx := stringToInt(params[2])
			list[idx] = intToString(a + b)
			pos = pos + 1 + paramCount
			continue
		case 2:
			a := stringToInt(params[0])
			if instruction.ParamType(0) == POSITIONAL {
				a = stringToInt(list[a])
			}
			b := stringToInt(params[1])
			if instruction.ParamType(1) == POSITIONAL {
				b = stringToInt(list[b])
			}
			idx := stringToInt(params[2])
			list[idx] = intToString(a * b)
		case 3:
			idx := stringToInt(params[0])
			return func(val int) {
				list[idx] = intToString(val)
				pos = pos + 1 + paramCount
				run(list, pos)
			}
		case 4:
			a := stringToInt(params[0])
			if instruction.ParamType(0) == POSITIONAL {
				a = stringToInt(list[a])
			}
			fmt.Println("**********")
			fmt.Println(a)
			fmt.Println("**********")
		case 5:
			a := stringToInt(params[0])
			if instruction.ParamType(0) == POSITIONAL {
				a = stringToInt(list[a])
			}
			if a != 0 {
				b := stringToInt(params[1])
				if instruction.ParamType(1) == POSITIONAL {
					b = stringToInt(list[b])
				}
				pos = b
				continue
			}
		case 6:
			a := stringToInt(params[0])
			if instruction.ParamType(0) == POSITIONAL {
				a = stringToInt(list[a])
			}
			if a == 0 {
				b := stringToInt(params[1])
				if instruction.ParamType(1) == POSITIONAL {
					b = stringToInt(list[b])
				}
				pos = b
				continue
			}
		case 7:
			a := stringToInt(params[0])
			if instruction.ParamType(0) == POSITIONAL {
				a = stringToInt(list[a])
			}
			b := stringToInt(params[1])
			if instruction.ParamType(1) == POSITIONAL {
				b = stringToInt(list[b])
			}
			idx := stringToInt(params[2])
			if a < b {
				list[idx] = intToString(1)
			} else {
				list[idx] = intToString(0)
			}
		case 8:
			a := stringToInt(params[0])
			if instruction.ParamType(0) == POSITIONAL {
				a = stringToInt(list[a])
			}
			b := stringToInt(params[1])
			if instruction.ParamType(1) == POSITIONAL {
				b = stringToInt(list[b])
			}
			idx := stringToInt(params[2])
			if a == b {
				list[idx] = intToString(1)
			} else {
				list[idx] = intToString(0)
			}
		case 99:
			fmt.Println("finished")
			return nil
		}
		pos = pos + 1 + paramCount
		debug("moving to pos", pos)

	}
	fmt.Println("reached end of list")
	return nil
}

type Instruction struct {
	Opcode     int
	ParamTypes []int
	Params     []string

	raw string
}

func (i *Instruction) ParamType(param int) int {
	if len(i.ParamTypes) <= param {
		return POSITIONAL
	}
	return i.ParamTypes[param]
}
func (i *Instruction) ParamCount() int {
	switch i.Opcode {
	case 1:
		return 3
	case 2:
		return 3
	case 3:
		return 1
	case 4:
		return 1
	case 5:
		return 2
	case 6:
		return 2
	case 7:
		return 3
	case 8:
		return 3
	case 99:
		return 0
	}
	return 0
}

func parseInstruction(i string) (*Instruction, error) {
	digits := strings.Split(i, "")
	var opcodeStr string
	paramTypes := make([]int, 0)
	if len(digits) > 2 {
		opcodeStr = strings.Join(digits[len(digits)-2:], "")
		paramTypesS := digits[:len(digits)-2]
		paramTypes = make([]int, len(paramTypesS))
		for p, paramType := range paramTypesS {
			idx := len(paramTypes) - 1 - p
			if paramType == "0" {
				paramTypes[idx] = POSITIONAL
			} else if paramType == "1" {
				paramTypes[idx] = IMMEDIATE
			}
		}
	} else {
		opcodeStr = i
	}
	opcode, err := strconv.Atoi(opcodeStr)
	if err != nil {
		return nil, err
	}

	return &Instruction{
		Opcode:     opcode,
		raw:        i,
		ParamTypes: paramTypes,
	}, nil

}
func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("can't convert: %s, to integer: %w", s, err)
		os.Exit(1)
	}
	return i
}
func intToString(i int) string {
	s := strconv.Itoa(i)
	return s
}
func debug(s ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		fmt.Println(s...)
	}
}
func debugf(s string, v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		fmt.Printf(s, v...)
	}
}
