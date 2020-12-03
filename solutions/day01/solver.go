package day01

import (
	"fmt"

	"github.com/drtchops/aoc2020/utils"
)

type Solver struct {
	input string
}

func New(input string) *Solver {
	return &Solver{input: input}
}

func (s *Solver) SolveA() string {
	entries := utils.ParseInputInts(s.input, "\n")

	for i, entry1 := range entries {
		for j := i + 1; j < len(entries); j++ {
			entry2 := entries[j]
			if entry1+entry2 == 2020 {
				return fmt.Sprint(entry1 * entry2)
			}
		}
	}

	return ""
}

func (s *Solver) SolveB() string {
	entries := utils.ParseInputInts(s.input, "\n")

	for i, entry1 := range entries {
		for j, entry2 := range entries {
			if i == j {
				continue
			}
			for k, entry3 := range entries {
				if k == i || k == j {
					continue
				}
				if entry1+entry2+entry3 == 2020 {
					return fmt.Sprint(entry1 * entry2 * entry3)
				}
			}
		}
	}

	return ""
}
