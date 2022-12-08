package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/scjudd/aoc-2022/pkg/advent"
)

type direction string

const (
	dirLeft  direction = "L"
	dirRight           = "R"
	dirUp              = "U"
	dirDown            = "D"
)

type motion struct {
	dir    direction
	amount int
}

type position struct {
	x int
	y int
}

func main() {
	year, day := 2022, 9
	session := advent.MustLoadSession()
	motions := parseInput(advent.MustGetInput(session, year, day))
	advent.PrintResult(advent.Check(session, year, day, 1, partOne(motions)))
	advent.PrintResult(advent.Check(session, year, day, 2, partTwo(motions)))
}

func partOne(motions []motion) int {
	var head, tail position
	visited := map[position]bool{position{0, 0}: true}

	for _, motion := range motions {
		for n := 0; n < motion.amount; n++ {
			head = moveHead(head, motion.dir)
			tail = moveTail(tail, head)
			visited[tail] = true
		}
	}

	count := 0
	for _ = range visited {
		count += 1
	}
	return count
}

func partTwo(motions []motion) int {
	var knots [10]position
	visited := map[position]bool{position{0, 0}: true}

	for _, motion := range motions {
		for n := 0; n < motion.amount; n++ {
			knots[0] = moveHead(knots[0], motion.dir)
			for i := 1; i < len(knots); i++ {
				knots[i] = moveTail(knots[i], knots[i-1])
			}
			visited[knots[len(knots)-1]] = true
		}
	}

	count := 0
	for _ = range visited {
		count += 1
	}
	return count
}

func moveHead(head position, dir direction) position {
	newHead := head

	switch dir {
	case dirLeft:
		newHead.x -= 1
	case dirRight:
		newHead.x += 1
	case dirUp:
		newHead.y -= 1
	case dirDown:
		newHead.y += 1
	}

	return newHead
}

func moveTail(tail position, head position) position {
	newTail := tail

	if (head.x == tail.x-2 && head.y == tail.y-1) ||
		(head.x == tail.x-1 && head.y == tail.y-2) ||
		(head.x == tail.x-2 && head.y == tail.y-2) {
		newTail.x -= 1
		newTail.y -= 1
	}

	if (head.x == tail.x+2 && head.y == tail.y-1) ||
		(head.x == tail.x+1 && head.y == tail.y-2) ||
		(head.x == tail.x+2 && head.y == tail.y-2) {
		newTail.x += 1
		newTail.y -= 1
	}

	if (head.x == tail.x+2 && head.y == tail.y+1) ||
		(head.x == tail.x+1 && head.y == tail.y+2) ||
		(head.x == tail.x+2 && head.y == tail.y+2) {
		newTail.x += 1
		newTail.y += 1
	}

	if (head.x == tail.x-2 && head.y == tail.y+1) ||
		(head.x == tail.x-1 && head.y == tail.y+2) ||
		(head.x == tail.x-2 && head.y == tail.y+2) {
		newTail.x -= 1
		newTail.y += 1
	}

	if head.x == tail.x && head.y == tail.y-2 {
		newTail.y -= 1
	}

	if head.y == tail.y && head.x == tail.x+2 {
		newTail.x += 1
	}

	if head.x == tail.x && head.y == tail.y+2 {
		newTail.y += 1
	}

	if head.y == tail.y && head.x == tail.x-2 {
		newTail.x -= 1
	}

	return newTail
}

func parseInput(input io.Reader) []motion {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	motions := []motion{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		var dir direction

		switch parts[0] {
		case "L":
			dir = dirLeft
		case "R":
			dir = dirRight
		case "U":
			dir = dirUp
		case "D":
			dir = dirDown
		}

		amount, _ := strconv.Atoi(parts[1])
		motions = append(motions, motion{dir, amount})
	}

	return motions
}
