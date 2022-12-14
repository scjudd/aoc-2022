package main

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPartOneExample(t *testing.T) {
	expected := 24
	got := partOne(exampleCave())
	if expected != got {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestPartTwoExample(t *testing.T) {
	expected := 93
	got := partTwo(parseInput(strings.NewReader("498,4 -> 498,6 -> 496,6\n503,4 -> 502,4 -> 502,9 -> 494,9\n")))
	if expected != got {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestParseInput(t *testing.T) {
	input := strings.NewReader("498,4 -> 498,6 -> 496,6\n503,4 -> 502,4 -> 502,9 -> 494,9\n")
	expected := exampleCave()
	got := parseInput(input)
	if diff := cmp.Diff(expected, got); diff != "" {
		t.Errorf("Input parsing failed, (-expected,+got)\n%s", diff)
	}
}

func exampleCave() map[position]entity {
	return map[position]entity{
		position{498, 4}: entityRock,
		position{498, 5}: entityRock,
		position{498, 6}: entityRock,
		position{497, 6}: entityRock,
		position{496, 6}: entityRock,
		position{503, 4}: entityRock,
		position{502, 4}: entityRock,
		position{502, 5}: entityRock,
		position{502, 6}: entityRock,
		position{502, 7}: entityRock,
		position{502, 8}: entityRock,
		position{502, 9}: entityRock,
		position{501, 9}: entityRock,
		position{500, 9}: entityRock,
		position{499, 9}: entityRock,
		position{498, 9}: entityRock,
		position{497, 9}: entityRock,
		position{496, 9}: entityRock,
		position{495, 9}: entityRock,
		position{494, 9}: entityRock,
	}
}
