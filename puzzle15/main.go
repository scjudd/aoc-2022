package main

import (
	"bufio"
	"io"
	"regexp"
	"strconv"

	"github.com/scjudd/aoc-2022/pkg/advent"
)

type position struct{ x, y int }

type scan struct {
	scanner position
	beacon  position
}

func main() {
	year, day := 2022, 15
	session := advent.MustLoadSession()
	scans := parseInput(advent.MustGetInput(session, year, day))
	advent.PrintResult(advent.Check(session, year, day, 1, partOne(scans, 2_000_000)))
	advent.PrintResult(advent.Check(session, year, day, 2, partTwo(scans, 4_000_000)))
}

func partOne(scans []scan, row int) int {
	impossible := make(map[int]struct{})

	for s := range scans {
		r := distance(scans[s].scanner, scans[s].beacon)
		if scans[s].scanner.y+r <= row || scans[s].scanner.y-r >= row {
			continue
		}

		r -= abs(scans[s].scanner.y - row)
		for x := scans[s].scanner.x - r; x <= scans[s].scanner.x+r; x++ {
			if scans[s].beacon.x == x && scans[s].beacon.y == row {
				continue
			}
			impossible[x] = struct{}{}
		}
	}

	count := 0
	for range impossible {
		count++
	}
	return count
}

func partTwo(scans []scan, size int) int {
	for s1 := range scans {
		r1 := distance(scans[s1].scanner, scans[s1].beacon)
	nextPosition:
		for _, pos := range border(scans[s1].scanner, r1+1) {
			if pos.x < 0 || pos.y < 0 || pos.x > size || pos.y > size {
				continue
			}
			for s2 := range scans {
				if s1 == s2 {
					continue
				}
				r2 := distance(scans[s2].scanner, scans[s2].beacon)
				if distance(scans[s2].scanner, pos) <= r2 {
					continue nextPosition
				}
			}
			return pos.x*4_000_000 + pos.y
		}
	}

	panic("distress beacon not found")
}

func border(p position, r int) []position {
	positions := []position{}
	for i := 0; i < r; i++ {
		positions = append(positions,
			position{x: p.x + i, y: p.y - r + i},
			position{x: p.x + r - i, y: p.y + i},
			position{x: p.x - i, y: p.y + r - i},
			position{x: p.x - r + i, y: p.y - i},
		)
	}
	return positions
}

func distance(p1, p2 position) int {
	return abs(p2.x-p1.x) + abs(p2.y-p1.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func parseInput(input io.Reader) []scan {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	re := regexp.MustCompile("[-0-9]+")
	scans := []scan{}

	for scanner.Scan() {
		var s scan
		line := scanner.Text()
		matches := re.FindAllString(line, -1)
		s.scanner.x, _ = strconv.Atoi(matches[0])
		s.scanner.y, _ = strconv.Atoi(matches[1])
		s.beacon.x, _ = strconv.Atoi(matches[2])
		s.beacon.y, _ = strconv.Atoi(matches[3])
		scans = append(scans, s)
	}

	return scans

}
