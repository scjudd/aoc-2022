package main

import (
	"fmt"
	"testing"
)

func TestPartOneExample(t *testing.T) {
	cases := []struct {
		signal string
		marker int
	}{
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("example %d", i+1), func(t *testing.T) {
			got := partOne(c.signal)
			if got != c.marker {
				t.Errorf("Expected %d, got %d", c.marker, got)
			}
		})
	}
}

func TestPartTwoExample(t *testing.T) {
	cases := []struct {
		signal string
		marker int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("example %d", i+1), func(t *testing.T) {
			got := partTwo(c.signal)
			if got != c.marker {
				t.Errorf("Expected %d, got %d", c.marker, got)
			}
		})
	}
}
