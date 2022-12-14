package main

import (
	"sort"

	"github.com/scjudd/aoc-2022/pkg/advent"
)

type op rune

const (
	opAdd      op = '+'
	opMultiply    = '*'
	opSquare      = '²'
)

type operation struct {
	op      op
	operand int
}

type monkey struct {
	items         []int
	operation     operation
	testDivisible int
	trueMonkey    int
	falseMonkey   int
	inspections   int
}

func main() {
	year, day := 2022, 11
	session := advent.MustLoadSession()
	advent.PrintResult(advent.Check(session, year, day, 1, partOne(initialMonkies())))
	advent.PrintResult(advent.Check(session, year, day, 2, partTwo(initialMonkies())))
}

func partOne(monkies []*monkey) int {
	for r := 0; r < 20; r++ {
		for _, m := range monkies {
			for _, i := range m.items {
				switch m.operation.op {
				case opAdd:
					i = i + m.operation.operand
				case opMultiply:
					i = i * m.operation.operand
				case opSquare:
					i = i * i
				}
				i /= 3
				if i%m.testDivisible == 0 {
					monkies[m.trueMonkey].items = append(monkies[m.trueMonkey].items, i)
				} else {
					monkies[m.falseMonkey].items = append(monkies[m.falseMonkey].items, i)
				}
				m.inspections++
			}
			m.items = []int{}
		}
	}
	sort.Slice(monkies, func(i, j int) bool {
		return monkies[i].inspections < monkies[j].inspections
	})
	end := len(monkies) - 1
	return monkies[end].inspections * monkies[end-1].inspections
}

func partTwo(monkies []*monkey) int {
	for r := 0; r < 10_000; r++ {
		common := 1
		for _, m := range monkies {
			common *= m.testDivisible
		}
		for _, m := range monkies {
			for _, i := range m.items {
				switch m.operation.op {
				case opAdd:
					i = (i + m.operation.operand) % common
				case opMultiply:
					i = (i * m.operation.operand) % common
				case opSquare:
					i = (i * i) % common
				}
				if i%m.testDivisible == 0 {
					monkies[m.trueMonkey].items = append(monkies[m.trueMonkey].items, i)
				} else {
					monkies[m.falseMonkey].items = append(monkies[m.falseMonkey].items, i)
				}
				m.inspections++
			}
			m.items = []int{}
		}
	}
	sort.Slice(monkies, func(i, j int) bool {
		return monkies[i].inspections < monkies[j].inspections
	})
	end := len(monkies) - 1
	return monkies[end].inspections * monkies[end-1].inspections
}

// Skip input parsing for now since my input is so small and I could do most of
// this with VIM macros.
func initialMonkies() []*monkey {
	return []*monkey{
		&monkey{
			items:         []int{91, 54, 70, 61, 64, 64, 60, 85},
			operation:     operation{op: opMultiply, operand: 13},
			testDivisible: 2,
			trueMonkey:    5,
			falseMonkey:   2,
		},
		&monkey{
			items:         []int{82},
			operation:     operation{op: opAdd, operand: 7},
			testDivisible: 13,
			trueMonkey:    4,
			falseMonkey:   3,
		},
		&monkey{
			items:         []int{84, 93, 70},
			operation:     operation{op: opAdd, operand: 2},
			testDivisible: 5,
			trueMonkey:    5,
			falseMonkey:   1,
		},
		&monkey{
			items:         []int{78, 56, 85, 93},
			operation:     operation{op: opMultiply, operand: 2},
			testDivisible: 3,
			trueMonkey:    6,
			falseMonkey:   7,
		},
		&monkey{
			items:         []int{64, 57, 81, 95, 52, 71, 58},
			operation:     operation{op: opSquare},
			testDivisible: 11,
			trueMonkey:    7,
			falseMonkey:   3,
		},
		&monkey{
			items:         []int{58, 71, 96, 58, 68, 90},
			operation:     operation{op: opAdd, operand: 6},
			testDivisible: 17,
			trueMonkey:    4,
			falseMonkey:   1,
		},
		&monkey{
			items:         []int{56, 99, 89, 97, 81},
			operation:     operation{op: opAdd, operand: 1},
			testDivisible: 7,
			trueMonkey:    0,
			falseMonkey:   2,
		},
		&monkey{
			items:         []int{68, 72},
			operation:     operation{op: opAdd, operand: 8},
			testDivisible: 19,
			trueMonkey:    6,
			falseMonkey:   0,
		},
	}
}
