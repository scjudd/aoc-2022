package main

import (
	"bufio"
	"io"
	"sort"
	"strconv"
	"strings"

	"github.com/scjudd/aoc-2022/pkg/advent"
)

type packetType int

const (
	typeInt packetType = iota
	typeList
)

type packet struct {
	t packetType
	i int
	l []*packet
}

func main() {
	year, day := 2022, 13
	session := advent.MustLoadSession()
	packets := parseInput(advent.MustGetInput(session, year, day))
	advent.PrintResult(advent.Check(session, year, day, 1, partOne(packets)))
	advent.PrintResult(advent.Check(session, year, day, 2, partTwo(packets)))
}

func partOne(packets []*packet) int {
	sum := 0

	for n, i, j := 1, 0, 1; j < len(packets); n, i, j = n+1, i+2, j+2 {
		if ordered, _ := compare(packets[i], packets[j]); ordered {
			sum += n
		}
	}

	return sum
}

func partTwo(packets []*packet) int {
	markerOne, markerTwo := l(l(i(2))), l(l(i(6)))
	packets = append(packets, markerOne, markerTwo)

	sort.Slice(packets, func(i, j int) bool {
		ordered, _ := compare(packets[i], packets[j])
		return ordered
	})

	var positionOne, positionTwo int

	for i := range packets {
		if packets[i] == markerOne {
			positionOne = i + 1
		} else if packets[i] == markerTwo {
			positionTwo = i + 1
		}
	}

	return positionOne * positionTwo
}

func compare(left, right *packet) (ordered, equal bool) {
	if left.t == typeInt && right.t == typeInt {
		return left.i < right.i, left.i == right.i
	}

	if left.t == typeInt {
		return compare(l(left), right)
	}

	if right.t == typeInt {
		return compare(left, l(right))
	}

	for n := 0; n < min(len(left.l), len(right.l)); n++ {
		ordered, equal = compare(left.l[n], right.l[n])
		if !equal {
			return
		}
	}

	return len(left.l) < len(right.l), len(left.l) == len(right.l)
}

func i(v int) *packet {
	return &packet{t: typeInt, i: v}
}

func l(v ...*packet) *packet {
	return &packet{t: typeList, l: v}
}

func min(x, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}

func parseInput(input io.Reader) []*packet {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	packets := []*packet{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		packets = append(packets, parsePacket(line))
	}

	return packets
}

func parsePacket(s string) *packet {
	tokens := tokenize(s)
	lists := []*packet{}
	for _, token := range tokens {
		switch token {
		case "[":
			p := &packet{t: typeList, l: []*packet{}}
			if len(lists) > 0 {
				lists[len(lists)-1].l = append(lists[len(lists)-1].l, p)
			}
			lists = append(lists, p)
		case "]":
			if len(lists) > 1 {
				lists = lists[:len(lists)-1]
			}
		default:
			i, _ := strconv.Atoi(token)
			p := &packet{t: typeInt, i: i}
			if len(lists) == 0 {
				return p
			}
			lists[len(lists)-1].l = append(lists[len(lists)-1].l, p)
		}
	}
	return lists[0]
}

func tokenize(s string) []string {
	tokens := []string{}
	var sb strings.Builder

	for _, char := range s {
		switch char {
		case '[':
			tokens = append(tokens, string(char))
		case ']':
			if sb.Len() > 0 {
				tokens = append(tokens, sb.String())
				sb.Reset()
			}
			tokens = append(tokens, string(char))
		case ',':
			if sb.Len() > 0 {
				tokens = append(tokens, sb.String())
				sb.Reset()
			}
		default:
			sb.WriteRune(char)
		}
	}

	return tokens
}
