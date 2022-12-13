package main

import (
	"bufio"
	"io"

	"github.com/scjudd/aoc-2022/pkg/advent"
)

type position struct {
	x, y int
}

type heightmap struct {
	start position
	end   position
	w     int
	h     int
	data  map[position]rune
}

func (m heightmap) height(p position) int {
	return int(m.data[p])
}

func main() {
	year, day := 2022, 12
	session := advent.MustLoadSession()
	m := parseInput(advent.MustGetInput(session, year, day))
	advent.PrintResult(advent.Check(session, year, day, 1, partOne(m)))
	advent.PrintResult(advent.Check(session, year, day, 2, partTwo(m)))
}

func partOne(m heightmap) int {
	return search(m, m.start, m.end)
}

func partTwo(m heightmap) int {
	var best int
	for start, r := range m.data {
		if r != 'a' {
			continue
		}
		distance := search(m, start, m.end)
		if distance == 0 {
			continue
		}
		if best == 0 || distance < best {
			best = distance
		}
	}
	return best
}

// Compute this shortest distance from start to end. If there is no valid path,
// zero is returned.
func search(m heightmap, start, end position) int {
	queue := []position{start}
	distance := map[position]int{start: 0}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, n := range neighbors(m, current) {
			if m.height(n) > m.height(current)+1 {
				continue
			}
			if d, ok := distance[n]; ok && d <= distance[current]+1 {
				continue
			}
			distance[n] = distance[current] + 1
			// fmt.Printf("%d, %d -> %d\n", n.x, n.y, distance[n])
			queue = append(queue, n)
		}
	}

	return distance[end]
}

func neighbors(m heightmap, p position) []position {
	neighbors := make([]position, 0, 4)
	if p.y > 0 {
		neighbors = append(neighbors, position{x: p.x, y: p.y - 1})
	}
	if p.x < m.w-1 {
		neighbors = append(neighbors, position{x: p.x + 1, y: p.y})
	}
	if p.y < m.h-1 {
		neighbors = append(neighbors, position{x: p.x, y: p.y + 1})
	}
	if p.x > 0 {
		neighbors = append(neighbors, position{x: p.x - 1, y: p.y})
	}
	return neighbors
}

func parseInput(input io.Reader) heightmap {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanRunes)

	var p position
	m := heightmap{h: 1, data: make(map[position]rune)}

	for scanner.Scan() {
		r := []rune(scanner.Text())[0]

		if r == '\n' {
			m.w = p.x
			m.h++
			p = position{x: 0, y: p.y + 1}
			continue
		}

		if r == 'S' {
			r = 'a'
			m.start = p
		}

		if r == 'E' {
			r = 'z'
			m.end = p
		}

		m.data[p] = r
		p.x++
	}

	return m
}
