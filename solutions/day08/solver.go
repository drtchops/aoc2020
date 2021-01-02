package day08

import (
	"fmt"
	"strconv"
	"strings"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

type operation string

var (
	NOP operation = "nop"
	ACC operation = "acc"
	JMP operation = "jmp"
)

type instruction struct {
	Operation operation
	Offset    int
}

func parse(input string) []instruction {
	lines := strings.Split(input, "\n")
	instructions := make([]instruction, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, " ")
		offset, _ := strconv.Atoi(parts[1])
		instructions[i] = instruction{
			Operation: operation(parts[0]),
			Offset:    offset,
		}
	}

	return instructions
}

func run(instructions []instruction) (int, bool) {
	i := 0
	acc := 0
	visited := make(map[int]bool)
	loops := false

	for {
		if visited[i] {
			loops = true
			break
		}
		if i >= len(instructions) {
			break
		}

		visited[i] = true
		inst := instructions[i]

		switch inst.Operation {
		case NOP:
			i++
		case ACC:
			acc += inst.Offset
			i++
		case JMP:
			i += inst.Offset
		}
	}

	return acc, loops
}

func (s *Solver) SolveA(input string) string {
	instructions := parse(input)
	acc, _ := run(instructions)
	return fmt.Sprint(acc)
}

func (s *Solver) SolveB(input string) string {
	instructions := parse(input)
	goodAcc := 0

	for i, inst := range instructions {
		if inst.Operation != NOP && inst.Operation != JMP {
			continue
		}

		newInstructions := make([]instruction, len(instructions))
		copy(newInstructions, instructions)
		newOp := NOP
		if inst.Operation == NOP {
			newOp = JMP
		}
		newInstructions[i] = instruction{
			Operation: newOp,
			Offset:    inst.Offset,
		}

		acc, looped := run(newInstructions)
		if !looped {
			goodAcc = acc
			break
		}
	}

	return fmt.Sprint(goodAcc)
}
