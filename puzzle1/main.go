package main

import (
	"bufio"
	"io"
	"strconv"

	"github.com/scjudd/aoc-2022/pkg/advent"
)

func main() {
	year, day := 2022, 1
	session := advent.MustLoadSession()
	data := parseInput(advent.MustGetInput(session, year, day))
	advent.PrintResult(advent.Check(session, year, day, 1, partOne(data)))
	advent.PrintResult(advent.Check(session, year, day, 2, partTwo(data)))
}

func partOne(calorieList [][]int) int {
	maxCalories := 0
	for _, elf := range calorieList {
		calories := sum(elf)
		if maxCalories < calories {
			maxCalories = calories
		}
	}
	return maxCalories
}

func partTwo(calorieList [][]int) int {
	maxCalories := []int{0, 0, 0}
	for _, elf := range calorieList {
		calories := sum(elf)
		maxList(maxCalories, calories)
	}
	return sum(maxCalories)
}

func parseInput(input io.Reader) [][]int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	calorieList := [][]int{}
	elf := []int{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if len(elf) != 0 {
				calorieList = append(calorieList, elf)
				elf = []int{}
			}
			continue
		}

		calories, _ := strconv.Atoi(line)
		elf = append(elf, calories)
	}

	if len(elf) != 0 {
		calorieList = append(calorieList, elf)
	}

	return calorieList
}

func sum(list []int) int {
	total := 0
	for _, num := range list {
		total += num
	}
	return total
}

func maxList(list []int, newNum int) {
	for idx := range list {
		if list[idx] < newNum {
			newNum, list[idx] = list[idx], newNum
		}
	}
}
