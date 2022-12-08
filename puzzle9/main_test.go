package main

import (
	"reflect"
	"strings"
	"testing"
)

var exampleOneMotions = []motion{{dirRight, 4}, {dirUp, 4}, {dirLeft, 3}, {dirDown, 1}, {dirRight, 4}, {dirDown, 1}, {dirLeft, 5}, {dirRight, 2}}
var exampleTwoMotions = []motion{{dirRight, 5}, {dirUp, 8}, {dirLeft, 8}, {dirDown, 3}, {dirRight, 17}, {dirDown, 10}, {dirLeft, 25}, {dirUp, 20}}

func TestPartOneExample(t *testing.T) {
	motions := exampleOneMotions
	expected := 13
	got := partOne(motions)
	if expected != got {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestPartTwoExample(t *testing.T) {
	motions := exampleTwoMotions
	expected := 36
	got := partTwo(motions)
	if expected != got {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestParseInput(t *testing.T) {
	input := strings.NewReader("R 4\nU 4\nL 3\nD 1\nR 4\nD 1\nL 5\nR 2")
	expected := exampleOneMotions
	got := parseInput(input)
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func TestMoveTail(t *testing.T) {
	tail := position{5, 5}

	testCases := []struct {
		head     position
		expected position
		desc     string
	}{
		{position{4, 3}, position{4, 4}, "top left (up)"},
		{position{3, 4}, position{4, 4}, "top left (left)"},
		{position{3, 3}, position{4, 4}, "top left (diagonal)"},
		{position{6, 3}, position{6, 4}, "top right (up)"},
		{position{7, 4}, position{6, 4}, "top right (right)"},
		{position{7, 3}, position{6, 4}, "top right (diagonal)"},
		{position{3, 6}, position{4, 6}, "bottom left (left)"},
		{position{4, 7}, position{4, 6}, "bottom left (down)"},
		{position{3, 7}, position{4, 6}, "bottom left (diagonal)"},
		{position{7, 6}, position{6, 6}, "bottom right (right)"},
		{position{6, 7}, position{6, 6}, "bottom right (down)"},
		{position{7, 7}, position{6, 6}, "bottom right (diagonal)"},
		{position{5, 3}, position{5, 4}, "top"},
		{position{7, 5}, position{6, 5}, "right"},
		{position{5, 7}, position{5, 6}, "down"},
		{position{3, 5}, position{4, 5}, "left"},
		{position{4, 4}, position{5, 5}, "touching (top left)"},
		{position{5, 4}, position{5, 5}, "touching (top)"},
		{position{6, 4}, position{5, 5}, "touching (top right)"},
		{position{6, 5}, position{5, 5}, "touching (right)"},
		{position{6, 6}, position{5, 5}, "touching (bottom right)"},
		{position{5, 6}, position{5, 5}, "touching (bottom)"},
		{position{4, 6}, position{5, 5}, "touching (bottom left)"},
		{position{4, 5}, position{5, 5}, "touching (left)"},
		{position{5, 5}, position{5, 5}, "touching (overlap)"},
	}

	for _, c := range testCases {
		t.Run(c.desc, func(t *testing.T) {
			got := moveTail(tail, c.head)
			if c.expected != got {
				t.Errorf("Expected %v, got %v", c.expected, got)
			}
		})
	}
}
