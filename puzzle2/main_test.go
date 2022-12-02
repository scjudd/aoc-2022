package main

import (
	"testing"
)

func TestPartOneExample(t *testing.T) {
	rounds := []round{
		{"A", "Y"},
		{"B", "X"},
		{"C", "Z"},
	}
	expected := 15
	got := partOne(rounds)
	if got != expected {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func TestPartTwoExample(t *testing.T) {
	rounds := []round{
		{"A", "Y"},
		{"B", "X"},
		{"C", "Z"},
	}
	expected := 12
	got := partTwo(rounds)
	if got != expected {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}
