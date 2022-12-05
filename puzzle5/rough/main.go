package main

import (
	"bufio"
	"fmt"
	"github.com/scjudd/aoc-2022/pkg/advent"
	"io"
	"regexp"
	"strconv"
)

func main() {
	a := advent.MustFromEnv(2022, 5)

	input, err := advent.GetInput(a)
	if err != nil {
		panic(err)
	}
	defer input.Close()

	stacks, instructions := parseInput(input)
	stacksCopy := make([]stack, len(stacks))
	copy(stacksCopy, stacks)

	fmt.Printf("Part 1: %s\n", partOne(stacks, instructions))
	fmt.Printf("Part 2: %s\n", partTwo(stacksCopy, instructions))

	// advent.PrintResult(advent.CheckPartOne(a, partOne(data)))
	// advent.PrintResult(advent.CheckPartTwo(a, partTwo(data)))
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

	s := ""
	for _, stack := range stacks {
		s += string(stack[len(stack)-1])
	}
	return s
}

func partTwo(stacks []stack, instructions []instruction) string {
	for _, inst := range instructions {
		from, to := inst.from-1, inst.to-1
		start, end := len(stacks[from])-inst.amount, len(stacks[from])
		stacks[to] = append(stacks[to], stacks[from][start:end]...)
		stacks[from] = stacks[from][:start]
	}

	s := ""
	for _, stack := range stacks {
		s += string(stack[len(stack)-1])
	}
	return s
}

type stack []rune

type instruction struct {
	from   int
	to     int
	amount int
}

func parseInput(input io.Reader) ([]stack, []instruction) {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		lines = append(lines, line)
	}

	stacks := parseStacks(lines)

	instructions := []instruction{}

	for scanner.Scan() {
		instructions = append(instructions, parseInstruction(scanner.Text()))
	}

	return stacks, instructions
}

func parseStacks(lines []string) []stack {
	stackCount := 0

	for _, line := range lines {
		numStacks := (len(line) + 1) / 4
		if numStacks > stackCount {
			stackCount = numStacks
		}
	}

	stacks := make([]stack, stackCount)

	for l := len(lines) - 2; l >= 0; l-- {
		line := []rune(lines[l])
		for s := range stacks {
			i := s*4 + 1
			if i >= len(line) {
				break
			}

			r := line[i]
			if r == ' ' {
				continue
			}

			stacks[s] = append(stacks[s], r)
		}
	}

	return stacks
}

func parseInstruction(s string) instruction {
	r := regexp.MustCompile("move ([0-9]+) from ([0-9]+) to ([0-9]+)")
	matches := r.FindStringSubmatch(s)
	amount, _ := strconv.Atoi(matches[1])
	from, _ := strconv.Atoi(matches[2])
	to, _ := strconv.Atoi(matches[3])
	return instruction{from: from, to: to, amount: amount}
}
