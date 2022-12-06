package main

import (
	"errors"
	"github.com/scjudd/aoc-2022/pkg/advent"
	"io"
)

func main() {
	a := advent.MustFromEnv(2022, 6)

	input, err := advent.GetInput(a)
	if err != nil {
		panic(err)
	}
	defer input.Close()

	signal := parseInput(input)

	advent.PrintResult(advent.CheckPartOne(a, partOne(signal)))
	advent.PrintResult(advent.CheckPartTwo(a, partTwo(signal)))
}

func partOne(signal []rune) int {
	end := 4

shiftWindow:
	for start := 0; end < len(signal); start, end = start+1, end+1 {
		window := signal[start:end]
		for i := 0; i < len(window); i++ {
			for j := i + 1; j < len(window); j++ {
				if window[i] == window[j] {
					continue shiftWindow
				}
			}
		}
		break

	}

	return end
}

func partTwo(signal []rune) int {
	end := 14

shiftWindow:
	for start := 0; end < len(signal); start, end = start+1, end+1 {
		window := signal[start:end]
		for i := 0; i < len(window); i++ {
			for j := i + 1; j < len(window); j++ {
				if window[i] == window[j] {
					continue shiftWindow
				}
			}
		}
		break

	}

	return end
}

func parseInput(input io.Reader) []rune {
	signal := []rune{}

	b := make([]byte, 1)
	for {
		_, err := input.Read(b)
		if err != nil && !errors.Is(err, io.EOF) {
			panic(err)
		}

		signal = append(signal, rune(b[0]))

		if err != nil {
			break
		}
	}

	return signal
}
