package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestPartOneExample(t *testing.T) {
	expected := 31
	got := partOne(exampleHeightMap())
	if expected != got {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestPartTwoExample(t *testing.T) {
	expected := 29
	got := partTwo(exampleHeightMap())
	if expected != got {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestParseInput(t *testing.T) {
	input := strings.NewReader("Sabqponm\nabcryxxl\naccszExk\nacctuvwj\nabdefghi")
	expected := exampleHeightMap()
	got := parseInput(input)
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Expected: %#+v\ngot: %#+v", expected, got)
	}
}

func exampleHeightMap() heightmap {
	return heightmap{
		start: position{0, 0},
		end:   position{5, 2},
		w:     8,
		h:     5,
		data: map[position]rune{
			position{0, 0}: 'a',
			position{1, 0}: 'a',
			position{2, 0}: 'b',
			position{3, 0}: 'q',
			position{4, 0}: 'p',
			position{5, 0}: 'o',
			position{6, 0}: 'n',
			position{7, 0}: 'm',
			position{0, 1}: 'a',
			position{1, 1}: 'b',
			position{2, 1}: 'c',
			position{3, 1}: 'r',
			position{4, 1}: 'y',
			position{5, 1}: 'x',
			position{6, 1}: 'x',
			position{7, 1}: 'l',
			position{0, 2}: 'a',
			position{1, 2}: 'c',
			position{2, 2}: 'c',
			position{3, 2}: 's',
			position{4, 2}: 'z',
			position{5, 2}: 'z',
			position{6, 2}: 'x',
			position{7, 2}: 'k',
			position{0, 3}: 'a',
			position{1, 3}: 'c',
			position{2, 3}: 'c',
			position{3, 3}: 't',
			position{4, 3}: 'u',
			position{5, 3}: 'v',
			position{6, 3}: 'w',
			position{7, 3}: 'j',
			position{0, 4}: 'a',
			position{1, 4}: 'b',
			position{2, 4}: 'd',
			position{3, 4}: 'e',
			position{4, 4}: 'f',
			position{5, 4}: 'g',
			position{6, 4}: 'h',
			position{7, 4}: 'i',
		},
	}
}
