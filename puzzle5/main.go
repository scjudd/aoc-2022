package main

import (
	"bufio"
	"github.com/scjudd/aoc-2022/pkg/advent"
	"io"
	"strconv"
	"strings"
)

type stack []rune

type instruction struct {
	from   int
	to     int
	amount int
}

func main() {
	a := advent.MustFromEnv(2022, 5)

	input := advent.MustGetInput(a)
	defer input.Close()

	partOneStacks, instructions := parseInput(input)
	partTwoStacks := copyStacks(partOneStacks)

	advent.PrintResult(advent.CheckPartOne(a, partOne(partOneStacks, instructions)))
	advent.PrintResult(advent.CheckPartTwo(a, partTwo(partTwoStacks, instructions)))
}

func partOne(stacks []stack, instructions []instruction) string {
	for _, inst := range instructions {
		for i := 0; i < inst.amount; i++ {
			from, to := inst.from-1, inst.to-1
			last := len(stacks[from]) - 1
			stacks[to] = append(stacks[to], stacks[from][last])
			stacks[from] = stacks[from][:last]
		}
	}

	return string(topCrates(stacks))
}

func partTwo(stacks []stack, instructions []instruction) string {
	for _, inst := range instructions {
		from, to := inst.from-1, inst.to-1
		start, end := len(stacks[from])-inst.amount, len(stacks[from])
		stacks[to] = append(stacks[to], stacks[from][start:end]...)
		stacks[from] = stacks[from][:start]
	}

	return string(topCrates(stacks))
}

func topCrates(stacks []stack) []rune {
	crates := make([]rune, len(stacks))
	for i, stack := range stacks {
		crates[i] = stack[len(stack)-1]
	}
	return crates
}

func parseInput(input io.Reader) ([]stack, []instruction) {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	stacksLines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		stacksLines = append(stacksLines, line)
	}
	stacks := parseStacks(stacksLines)

	instructions := []instruction{}
	for scanner.Scan() {
		instructions = append(instructions, parseInstruction(scanner.Text()))
	}

	return stacks, instructions
}

func parseStacks(lines []string) []stack {
	numStacks := 0
	for _, line := range lines {
		n := (len(line) + 1) / 4
		if n > numStacks {
			numStacks = n
		}
	}

	stacks := make([]stack, numStacks)

	for l := len(lines) - 2; l >= 0; l-- {
		line := lines[l]
		for s := range stacks {
			idx := s*4 + 1
			if idx >= len(line) {
				break
			}

			r := rune(line[idx])
			if r == ' ' {
				continue
			}

			stacks[s] = append(stacks[s], r)
		}
	}

	return stacks
}

func parseInstruction(s string) instruction {
	pieces := strings.Split(s, " ")
	from := parseInt(pieces[3])
	to := parseInt(pieces[5])
	amount := parseInt(pieces[1])
	return instruction{from: from, to: to, amount: amount}
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func copyStacks(src []stack) []stack {
	dst := make([]stack, len(src))
	for i := range src {
		dst[i] = make(stack, len(src[i]))
		copy(dst[i], src[i])
	}
	return dst
}
