package main

import (
	"github.com/scjudd/aoc-2022/pkg/advent"
	"strings"
	"strconv"
	"io"
	"bufio"
)

func main() {
	a := advent.MustFromEnv(2022, 4)

	input := advent.MustGetInput(a)
	defer input.Close()

	data := parseInput(input)

	advent.PrintResult(advent.CheckPartOne(a, partOne(data)))
	advent.PrintResult(advent.CheckPartTwo(a, partTwo(data)))
}

func partOne(data []string) int {
	count := 0

	for _, line := range data {
		pair := strings.Split(line, ",")
		left, right := pair[0], pair[1]
		pair = strings.Split(left, "-")
		leftMinStr, leftMaxStr := pair[0], pair[1]
		leftMin, _ := strconv.Atoi(leftMinStr)
		leftMax, _ := strconv.Atoi(leftMaxStr)
		pair = strings.Split(right, "-")
		rightMinStr, rightMaxStr := pair[0], pair[1]
		rightMin, _ := strconv.Atoi(rightMinStr)
		rightMax, _ := strconv.Atoi(rightMaxStr)

		if rightMin >= leftMin && rightMax <= leftMax {
			count++
		} else if  leftMin >= rightMin && leftMax <= rightMax {
			count++
		}
	}

	return count
}

func partTwo(data []string) int {
	count := 0

	for _, line := range data {
		pair := strings.Split(line, ",")
		left, right := pair[0], pair[1]
		pair = strings.Split(left, "-")
		leftMinStr, leftMaxStr := pair[0], pair[1]
		leftMin, _ := strconv.Atoi(leftMinStr)
		leftMax, _ := strconv.Atoi(leftMaxStr)
		pair = strings.Split(right, "-")
		rightMinStr, rightMaxStr := pair[0], pair[1]
		rightMin, _ := strconv.Atoi(rightMinStr)
		rightMax, _ := strconv.Atoi(rightMaxStr)

		if rightMin >= leftMin && rightMin <= leftMax {
			count++
		} else if leftMin >= rightMin && leftMin <= rightMax {
			count++
		} else if rightMax <= leftMax && rightMax >= leftMin {
			count++
		} else if leftMax <= rightMax && leftMax >= rightMin {
			count++
		}
	}

	return count
}

func parseInput(input io.Reader) []string {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	data := []string{}

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data
}
