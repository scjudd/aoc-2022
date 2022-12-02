package main

import (
	"bufio"
	"github.com/scjudd/aoc-2022/pkg/advent"
	"io"
	"strconv"
	"strings"
)

type span struct {
	min int
	max int
}

func (this span) contains(other span) bool {
	return this.min <= other.min && this.max >= other.max
}

func (this span) overlaps(other span) bool {
	return this.max >= other.min && this.min <= other.max
}

func main() {
	a := advent.MustFromEnv(2022, 4)

	input, err := advent.GetInput(a)
	if err != nil {
		panic(err)
	}
	defer input.Close()

	pairs := parseInput(input)

	advent.PrintResult(advent.CheckPartOne(a, partOne(pairs)))
	advent.PrintResult(advent.CheckPartTwo(a, partTwo(pairs)))
}

func partOne(pairs [][2]span) int {
	count := 0

	for _, pair := range pairs {
		if pair[0].contains(pair[1]) || pair[1].contains(pair[0]) {
			count++
		}
	}

	return count
}

func partTwo(pairs [][2]span) int {
	count := 0

	for _, pair := range pairs {
		if pair[0].overlaps(pair[1]) {
			count++
		}
	}

	return count
}

func parseInput(input io.Reader) [][2]span {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	pairs := [][2]span{}

	for scanner.Scan() {
		pairStrings := strings.Split(scanner.Text(), ",")
		leftStrings := strings.Split(pairStrings[0], "-")
		rightStrings := strings.Split(pairStrings[1], "-")
		pairs = append(pairs, [2]span{
			{min: parseInt(leftStrings[0]), max: parseInt(leftStrings[1])},
			{min: parseInt(rightStrings[0]), max: parseInt(rightStrings[1])},
		})
	}

	return pairs
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
