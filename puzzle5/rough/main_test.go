package main

import (
	"testing"
)

func TestPartOneExample(t *testing.T) {
	stacks := []stack{
		[]rune("ZN"),
		[]rune("MCD"),
		[]rune("P"),
	}
	instructions := []instruction{
		{from: 2, to: 1, amount: 1},
		{from: 1, to: 3, amount: 3},
		{from: 2, to: 1, amount: 2},
		{from: 1, to: 2, amount: 1},
	}
	expected := "CMZ"
	got := partOne(stacks, instructions)
	if expected != got {
		t.Errorf("Expected %s, got %s\n", expected, got)
	}
}

func TestPartTwoExample(t *testing.T) {
	stacks := []stack{
		[]rune("ZN"),
		[]rune("MCD"),
		[]rune("P"),
	}
	instructions := []instruction{
		{from: 2, to: 1, amount: 1},
		{from: 1, to: 3, amount: 3},
		{from: 2, to: 1, amount: 2},
		{from: 1, to: 2, amount: 1},
	}
	expected := "MCD"
	got := partTwo(stacks, instructions)
	if expected != got {
		t.Errorf("Expected %s, got %s\n", expected, got)
	}
}
