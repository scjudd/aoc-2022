package main

import (
	"bufio"
	"io"
	"strings"

	"github.com/scjudd/aoc-2022/pkg/advent"
)

type round [2]string

type scoring struct {
	round round
	score int
}

func main() {
	year, day := 2022, 2
	session := advent.MustLoadSession()
	data := parseInput(advent.MustGetInput(session, year, day))
	advent.PrintResult(advent.Check(session, year, day, 1, partOne(data)))
	advent.PrintResult(advent.Check(session, year, day, 2, partTwo(data)))
}

func partOne(rounds []round) int {
	scoreTable := []scoring{
		{round{"A", "X"}, 1 + 3}, // rock -> rock (1) -> draw (3)
		{round{"A", "Y"}, 2 + 6}, // rock -> paper (2) -> win (6)
		{round{"A", "Z"}, 3 + 0}, // rock -> scissors (3) -> lose (0)
		{round{"B", "X"}, 1 + 0}, // paper -> rock (1) -> lose (0)
		{round{"B", "Y"}, 2 + 3}, // paper -> paper (2) -> draw (3)
		{round{"B", "Z"}, 3 + 6}, // paper -> scissors (3) -> win (6)
		{round{"C", "X"}, 1 + 6}, // scissors -> rock (1) -> win (6)
		{round{"C", "Y"}, 2 + 0}, // scissors -> paper (2) -> lose (0)
		{round{"C", "Z"}, 3 + 3}, // scissors -> scissors (3) -> draw (3)
	}

	return score(scoreTable, rounds)
}

func partTwo(rounds []round) int {
	scoreTable := []scoring{
		{round{"A", "X"}, 0 + 3}, // rock -> lose (0) -> scissors (3)
		{round{"A", "Y"}, 3 + 1}, // rock -> draw (3) -> rock (1)
		{round{"A", "Z"}, 6 + 2}, // rock -> win (6) -> paper (2)
		{round{"B", "X"}, 0 + 1}, // paper -> lose (0) -> rock (1)
		{round{"B", "Y"}, 3 + 2}, // paper -> draw (3) -> paper (2)
		{round{"B", "Z"}, 6 + 3}, // paper -> win (6) -> scissors (3)
		{round{"C", "X"}, 0 + 2}, // scissors -> lose (0) -> paper (2)
		{round{"C", "Y"}, 3 + 3}, // scissors -> draw (3) -> scissors (3)
		{round{"C", "Z"}, 6 + 1}, // scissors -> win (6) -> rock (1)
	}

	return score(scoreTable, rounds)
}

func score(scoreTable []scoring, rounds []round) int {
	total := 0
	for _, round := range rounds {
		for _, scoring := range scoreTable {
			if round == scoring.round {
				total += scoring.score
				break
			}
		}
	}
	return total
}

func parseInput(input io.Reader) []round {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	rounds := []round{}

	for scanner.Scan() {
		r := strings.Split(scanner.Text(), " ")
		round := round{r[0], r[1]}
		rounds = append(rounds, round)
	}

	return rounds
}
