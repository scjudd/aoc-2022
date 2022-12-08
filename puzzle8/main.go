package main

import (
	"bufio"
	"io"
	"strconv"

	"github.com/scjudd/aoc-2022/pkg/advent"
)

func main() {
	year, day := 2022, 8
	session := advent.MustLoadSession()
	trees := parseInput(advent.MustGetInput(session, year, day))
	advent.PrintResult(advent.Check(session, year, day, 1, partOne(trees)))
	advent.PrintResult(advent.Check(session, year, day, 2, partTwo(trees)))
}

func partOne(trees [][]int) int {
	count := 0
	for y := range trees {
		for x := range trees[y] {
			if visible(trees, x, y) {
				count += 1
			}
		}
	}
	return count
}

func partTwo(trees [][]int) int {
	bestScore := 0
	for y := range trees {
		for x := range trees[y] {
			score := scenicScore(trees, x, y)
			if score > bestScore {
				bestScore = score
			}
		}
	}
	return bestScore
}

func visible(trees [][]int, x, y int) bool {
	h := trees[y][x]
	left, top, right, bottom := true, true, true, true

	for i := 0; i < x; i++ {
		if trees[y][i] >= h {
			left = false
			break
		}
	}

	for i := 0; i < y; i++ {
		if trees[i][x] >= h {
			top = false
			break
		}
	}

	for i := x + 1; i < len(trees[y]); i++ {
		if trees[y][i] >= h {
			right = false
			break
		}
	}

	for i := y + 1; i < len(trees); i++ {
		if trees[i][x] >= h {
			bottom = false
			break
		}
	}

	return left || top || right || bottom
}

func scenicScore(trees [][]int, x, y int) int {
	h := trees[y][x]
	left, top, right, bottom := 0, 0, 0, 0

	for i := x - 1; i >= 0; i-- {
		left++
		if trees[y][i] >= h {
			break
		}
	}

	for i := y - 1; i >= 0; i-- {
		top++
		if trees[i][x] >= h {
			break
		}
	}

	for i := x + 1; i < len(trees[y]); i++ {
		right++
		if trees[y][i] >= h {
			break
		}
	}

	for i := y + 1; i < len(trees); i++ {
		bottom++
		if trees[i][x] >= h {
			break
		}
	}

	return left * top * right * bottom
}

func parseInput(input io.Reader) [][]int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanRunes)

	trees := [][]int{[]int{}}

	for scanner.Scan() {
		s := scanner.Text()
		row := len(trees) - 1

		if s == "\n" {
			trees = append(trees, []int{})
			continue
		}

		h, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		trees[row] = append(trees[row], h)
	}

	trees = trees[:len(trees)-1]

	return trees
}
