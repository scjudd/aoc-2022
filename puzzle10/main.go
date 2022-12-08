package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/scjudd/aoc-2022/pkg/advent"
)

type instruction struct {
	command string
	operand int
}

func main() {
	year, day := 2022, 10
	session := advent.MustLoadSession()
	instructions := parseInput(advent.MustGetInput(session, year, day))
	advent.PrintResult(advent.Check(session, year, day, 1, partOne(instructions)))
	fmt.Printf("Part 2:\n%s", partTwo(instructions))
}

func partOne(instructions []instruction) int {
	pc := 0
	register := 1

	addxInProgress := false
	addxOperand := 0

	signalStrength := 0

	for cycle := 1; pc < len(instructions); cycle++ {
		if cycle%40 == 20 {
			signalStrength += cycle * register
		}

		inst := instructions[pc]
		if inst.command == "addx" && !addxInProgress {
			addxInProgress, addxOperand = true, inst.operand
			continue
		} else if addxInProgress {
			register += addxOperand
			addxInProgress = false
		}

		pc += 1
	}

	return signalStrength
}

func partTwo(instructions []instruction) string {
	pc := 0
	register := 1

	addxInProgress := false
	addxOperand := 0

	var framebuffer strings.Builder

	for cycle := 1; pc < len(instructions); cycle++ {
		position := (cycle - 1) % 40
		if position == register-1 || position == register || position == register+1 {
			framebuffer.WriteString("#")
		} else {
			framebuffer.WriteString(".")
		}
		if position == 39 {
			framebuffer.WriteString("\n")
		}

		inst := instructions[pc]
		if inst.command == "addx" && !addxInProgress {
			addxInProgress, addxOperand = true, inst.operand
			continue
		} else if addxInProgress {
			register += addxOperand
			addxInProgress = false
		}

		pc += 1
	}

	return framebuffer.String()
}

func parseInput(input io.Reader) []instruction {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	instructions := []instruction{}

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		inst := instruction{command: parts[0]}
		if len(parts) == 2 {
			operand, _ := strconv.Atoi(parts[1])
			inst.operand = operand
		}
		instructions = append(instructions, inst)
	}

	return instructions
}
