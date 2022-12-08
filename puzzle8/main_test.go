package main

import "testing"

func TestPartOneExample(t *testing.T) {
	trees := [][]int{
		[]int{3, 0, 3, 7, 3},
		[]int{2, 5, 5, 1, 2},
		[]int{6, 5, 3, 3, 2},
		[]int{3, 3, 5, 4, 9},
		[]int{3, 5, 3, 9, 0},
	}
	expected := 21
	got := partOne(trees)
	if expected != got {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestPartTwoExample(t *testing.T) {
	trees := [][]int{
		[]int{3, 0, 3, 7, 3},
		[]int{2, 5, 5, 1, 2},
		[]int{6, 5, 3, 3, 2},
		[]int{3, 3, 5, 4, 9},
		[]int{3, 5, 3, 9, 0},
	}
	expected := 8
	got := partTwo(trees)
	if expected != got {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}
