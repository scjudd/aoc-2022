package main

import (
	"bufio"
	"github.com/scjudd/aoc-2022/pkg/advent"
	"io"
	"strings"
)

func main() {
	a := advent.MustFromEnv(2022, 3)

	input := advent.MustGetInput(a)
	defer input.Close()

	sacks := parseInput(input)

	advent.PrintResult(advent.CheckPartOne(a, partOne(sacks)))
	advent.PrintResult(advent.CheckPartTwo(a, partTwo(sacks)))
}

func partOne(sacks []string) int {
	score := 0

	for _, sack := range sacks {
		mid, end := len(sack)/2, len(sack)
		left, right := sack[0:mid], sack[mid:end]

		for _, item := range left {
			if strings.ContainsRune(right, item) {
				score += priority(item)
				break
			}
		}
	}

	return score
}

func partTwo(sacks []string) int {
	const groupSize = 3
	score := 0

	for groupStart := 0; groupStart <= len(sacks)-groupSize; groupStart += groupSize {
		firstSack := sacks[groupStart]
		otherSacks := sacks[groupStart+1 : groupStart+groupSize]
	nextItem:
		for _, item := range firstSack {
			for _, sack := range otherSacks {
				if !strings.ContainsRune(sack, item) {
					continue nextItem
				}
			}
			score += priority(item)
			break
		}
	}

	return score
}

func priority(item rune) int {
	items := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return strings.IndexRune(items, item) + 1
}

func parseInput(input io.Reader) []string {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	sacks := []string{}

	for scanner.Scan() {
		sacks = append(sacks, scanner.Text())
	}

	return sacks
}
