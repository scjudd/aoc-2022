package main

import (
	"github.com/scjudd/aoc-2022/pkg/advent"
	"io"
	"strings"
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

func partOne(signal string) int {
	return markerIndex(signal, 4)
}

func partTwo(signal string) int {
	return markerIndex(signal, 14)
}

func markerIndex(signal string, numUnique int) int {
	end := numUnique

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

func parseInput(input io.Reader) string {
	buf := new(strings.Builder)

	_, err := io.Copy(buf, input)
	if err != nil {
		panic(err)
	}

	return buf.String()
}
