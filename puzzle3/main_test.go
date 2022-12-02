package main

import (
	"testing"
)

func TestPartOneExample(t *testing.T) {
	sacks := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}
	expected := 157
	got := partOne(sacks)
	if expected != got {
		t.Errorf("Expected %d, got %d\n", expected, got)
	}
}

func TestPartTwoExample(t *testing.T) {
	sacks := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}
	expected := 70
	got := partTwo(sacks)
	if expected != got {
		t.Errorf("Expected %d, got %d\n", expected, got)
	}
}
