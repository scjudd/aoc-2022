package main

import (
	"reflect"
	"testing"
)

func TestPartOneExample(t *testing.T) {
	expected := 26
	got := partOne(exampleScans(), 10)
	if expected != got {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestPartTwoExample(t *testing.T) {
	expected := 56000011
	got := partTwo(exampleScans(), 20)
	if expected != got {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestDistance(t *testing.T) {
	p1, p2 := position{x: 0, y: 1}, position{x: 5, y: 7}
	expected := 11
	got := distance(p1, p2)
	if expected != got {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestBorder(t *testing.T) {
	expected := []position{
		{x: 0, y: -3},
		{x: 3, y: 0},
		{x: 0, y: 3},
		{x: -3, y: 0},

		{x: 1, y: -2},
		{x: 2, y: 1},
		{x: -1, y: 2},
		{x: -2, y: -1},

		{x: 2, y: -1},
		{x: 1, y: 2},
		{x: -2, y: 1},
		{x: -1, y: -2},
	}
	got := border(position{x: 0, y: 0}, 3)
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func exampleScans() []scan {
	return []scan{
		{
			scanner: position{x: 2, y: 18},
			beacon:  position{x: -2, y: 15},
		},
		{
			scanner: position{x: 9, y: 16},
			beacon:  position{x: 10, y: 16},
		},
		{
			scanner: position{x: 13, y: 2},
			beacon:  position{x: 15, y: 3},
		},
		{
			scanner: position{x: 12, y: 14},
			beacon:  position{x: 10, y: 16},
		},
		{
			scanner: position{x: 10, y: 20},
			beacon:  position{x: 10, y: 16},
		},
		{
			scanner: position{x: 14, y: 17},
			beacon:  position{x: 10, y: 16},
		},
		{
			scanner: position{x: 8, y: 7},
			beacon:  position{x: 2, y: 10},
		},
		{
			scanner: position{x: 2, y: 0},
			beacon:  position{x: 2, y: 10},
		},
		{
			scanner: position{x: 0, y: 11},
			beacon:  position{x: 2, y: 10},
		},
		{
			scanner: position{x: 20, y: 14},
			beacon:  position{x: 25, y: 17},
		},
		{
			scanner: position{x: 17, y: 20},
			beacon:  position{x: 21, y: 22},
		},
		{
			scanner: position{x: 16, y: 7},
			beacon:  position{x: 15, y: 3},
		},
		{
			scanner: position{x: 14, y: 3},
			beacon:  position{x: 15, y: 3},
		},
		{
			scanner: position{x: 20, y: 1},
			beacon:  position{x: 15, y: 3},
		},
	}
}
