package main

import (
	"testing"
)

func TestPartOneExample(t *testing.T) {
	data := exampleFilesystem()
	expected := 95437
	got := partOne(data)
	if expected != got {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestPartTwoExample(t *testing.T) {
	data := exampleFilesystem()
	expected := 24933642
	got := partTwo(data)
	if expected != got {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func exampleFilesystem() *file {
	root := newRoot()
	cwd := root

	cwd.mkdir("a")
	cwd.addFile("b.txt", 14848514)
	cwd.addFile("c.dat", 8504156)
	cwd.mkdir("d")
	cwd = cwd.cd("a")
	cwd.mkdir("e")
	cwd.addFile("f", 29116)
	cwd.addFile("g", 2557)
	cwd.addFile("h.lst", 62596)
	cwd = cwd.cd("e")
	cwd.addFile("i", 584)
	cwd = cwd.cd("..")
	cwd = cwd.cd("..")
	cwd = cwd.cd("d")
	cwd.addFile("j", 4060174)
	cwd.addFile("d.log", 8033020)
	cwd.addFile("d.ext", 5626152)
	cwd.addFile("k", 7214296)

	return root
}
