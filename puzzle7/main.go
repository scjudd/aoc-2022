package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/scjudd/aoc-2022/pkg/advent"
)

type fileType int

const (
	typeDir fileType = iota
	typeFile
)

type file struct {
	fileType fileType
	name     string
	size     int
	parent   *file
	children []*file
}

func newRoot() *file {
	return &file{
		fileType: typeDir,
		name:     "/",
		children: []*file{},
	}
}

func (f *file) cd(name string) *file {
	if name == ".." {
		return f.parent
	}
	if name == "/" {
		root := f
		for root.parent != nil {
			root = root.parent
		}
		return root
	}
	for _, child := range f.children {
		if child.name == name && child.fileType == typeDir {
			return child
		}
	}
	return nil
}

func (f *file) mkdir(name string) *file {
	child := &file{
		fileType: typeDir,
		name:     name,
		parent:   f,
		children: []*file{},
	}
	f.children = append(f.children, child)
	return child
}

func (f *file) addFile(name string, size int) *file {
	child := &file{
		fileType: typeFile,
		name:     name,
		size:     size,
		parent:   f,
	}
	f.children = append(f.children, child)
	return child
}

func (f *file) totalSize() int {
	total := f.size
	for _, child := range f.children {
		total += child.totalSize()
	}
	return total
}

func (f *file) String() string {
	return f.buildString(0)
}

func (f *file) buildString(indents int) string {
	const indentSize = 2

	var b strings.Builder

	for i := 0; i < indents; i++ {
		b.WriteString(" ")
	}

	b.WriteString("- ")
	b.WriteString(f.name)

	if f.fileType == typeDir {
		b.WriteString(" (dir)\n")

		for _, child := range f.children {
			b.WriteString(child.buildString(indents + indentSize))
		}
	}

	if f.fileType == typeFile {
		b.WriteString(" (file, size=")
		b.WriteString(strconv.Itoa(f.size))
		b.WriteString(")\n")
	}

	return b.String()
}

func main() {
	year, day := 2022, 7
	session := advent.MustLoadSession()
	data := parseInput(advent.MustGetInput(session, year, day))
	fmt.Println(data)
	advent.PrintResult(advent.Check(session, year, day, 1, partOne(data)))
	advent.PrintResult(advent.Check(session, year, day, 2, partTwo(data)))
}

func partOne(f *file) int {
	if f.fileType != typeDir {
		return 0
	}

	sum, totalSize := 0, f.totalSize()
	if totalSize <= 100_000 {
		sum += totalSize
	}
	for _, child := range f.children {
		sum += partOne(child)
	}

	return sum
}

func partTwo(root *file) int {
	diskSize := 70_000_000
	required := 30_000_000
	used := root.totalSize()
	available := diskSize - used
	needed := required - available

	var walk func(*file, int) int

	walk = func(f *file, best int) int {
		if f.fileType != typeDir {
			return best
		}

		totalSize := f.totalSize()
		if totalSize >= needed && totalSize < best {
			best = totalSize
		}

		for _, child := range f.children {
			childBest := walk(child, best)
			if childBest < best {
				best = childBest
			}
		}

		return best
	}

	return walk(root, used)
}

func parseInput(input io.Reader) *file {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	root := newRoot()
	cwd := root

	for scanner.Scan() {
	afterScan:
		line := scanner.Text()

		if strings.Contains(line, "$ cd") {
			parts := strings.Split(line, " ")
			dir := parts[2]
			cwd = cwd.cd(dir)
			continue
		}

		if strings.Contains(line, "$ ls") {
			for scanner.Scan() {
				line := scanner.Text()

				if line == "" {
					return root
				}

				if strings.Contains(line, "$") {
					goto afterScan
				}

				parts := strings.Split(line, " ")

				if parts[0] == "dir" {
					dir := parts[1]
					cwd.mkdir(dir)
					continue
				}

				name := parts[1]
				size, err := strconv.Atoi(parts[0])
				if err != nil {
					panic(err)
				}

				cwd.addFile(name, size)
			}
		}
	}

	return root
}
