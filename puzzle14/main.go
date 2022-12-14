package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/scjudd/aoc-2022/pkg/advent"
)

type position struct{ x, y int }

type entity int

const (
	entityAir entity = iota
	entityRock
	entitySand
)

func main() {
	year, day := 2022, 14
	session := advent.MustLoadSession()

	partOneCave := parseInput(advent.MustGetInput(session, year, day))
	partTwoCave := make(map[position]entity)
	for k, v := range partOneCave {
		partTwoCave[k] = v
	}

	advent.PrintResult(advent.Check(session, year, day, 1, partOne(partOneCave)))
	advent.PrintResult(advent.Check(session, year, day, 2, partTwo(partTwoCave)))
}

func partOne(cave map[position]entity) int {
	var lowest int
	for position, entity := range cave {
		if entity == entityRock && position.y > lowest {
			lowest = position.y
		}
	}

	origin := position{500, 0}
	rested := 0

	for {
		last := origin
		for {
			position := dropUnit(cave, last, 0)
			if position.y > lowest {
				return rested
			}
			if position == last {
				cave[position] = entitySand
				rested++
				break
			}
			last = position
		}
	}

	return rested
}

func partTwo(cave map[position]entity) int {
	var floor int
	for position, entity := range cave {
		if entity == entityRock && position.y > floor {
			floor = position.y
		}
	}
	floor += 2

	origin := position{500, 0}
	rested := 0

	for {
		last := origin
		for {
			position := dropUnit(cave, last, floor)
			if position == last {
				cave[position] = entitySand
				rested++
				if position == origin {
					return rested
				}
				break
			}
			last = position
		}
	}

	return rested
}

func dropUnit(cave map[position]entity, start position, floor int) position {
	possibilities := []position{
		position{start.x, start.y + 1},
		position{start.x - 1, start.y + 1},
		position{start.x + 1, start.y + 1},
	}

	for i := range possibilities {
		if cave[possibilities[i]] == entityAir {
			if floor == 0 || possibilities[i].y < floor {
				return possibilities[i]
			}
		}
	}

	return start
}

func parseInput(input io.Reader) map[position]entity {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	cave := make(map[position]entity)

	for scanner.Scan() {
		line := scanner.Text()
		points := []position{}
		for _, s := range strings.Split(line, " -> ") {
			var point position
			v := strings.Split(s, ",")
			point.x, _ = strconv.Atoi(v[0])
			point.y, _ = strconv.Atoi(v[1])
			points = append(points, point)
		}
		drawRocks(cave, points)
	}

	return cave
}

func drawRocks(cave map[position]entity, points []position) {
	for i, j := 0, 1; j < len(points); i, j = i+1, j+1 {
		one, two := points[i], points[j]

		if one.x != two.x {
			for x := min(one.x, two.x); x <= max(one.x, two.x); x++ {
				cave[position{x: x, y: one.y}] = entityRock
			}
		} else if one.y != two.y {
			for y := min(one.y, two.y); y <= max(one.y, two.y); y++ {
				cave[position{x: one.x, y: y}] = entityRock
			}
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
