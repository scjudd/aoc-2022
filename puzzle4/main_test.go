package main

import (
	"testing"
)

func TestPartOneExample(t *testing.T) {
	pairs := [][2]span{
		{{2, 4}, {6, 8}},
		{{2, 3}, {4, 5}},
		{{5, 7}, {7, 9}},
		{{2, 8}, {3, 7}},
		{{6, 6}, {4, 6}},
		{{2, 6}, {4, 8}},
	}
	expected := 2
	got := partOne(pairs)
	if expected != got {
		t.Errorf("Expected %d, got %d\n", expected, got)
	}
}

func TestPartTwoExample(t *testing.T) {
	pairs := [][2]span{
		{{2, 4}, {6, 8}},
		{{2, 3}, {4, 5}},
		{{5, 7}, {7, 9}},
		{{2, 8}, {3, 7}},
		{{6, 6}, {4, 6}},
		{{2, 6}, {4, 8}},
	}
	expected := 4
	got := partTwo(pairs)
	if expected != got {
		t.Errorf("Expected %d, got %d\n", expected, got)
	}
}
