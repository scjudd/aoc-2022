package main

import (
	"fmt"
	"testing"
)

func TestPartOneExample(t *testing.T) {
	cases := []struct {
		str []rune
		pos int
	}{
		{[]rune("bvwbjplbgvbhsrlpgdmjqwftvncz"), 5},
		{[]rune("nppdvjthqldpwncqszvftbrmjlhg"), 6},
		{[]rune("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"), 10},
		{[]rune("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"), 11},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("example %d", i+1), func(t *testing.T) {
			got := partOne(c.str)
			if got != c.pos {
				t.Errorf("Expected %d, got %d", c.pos, got)
			}
		})
	}
}

func TestPartTwoExample(t *testing.T) {
	cases := []struct {
		str []rune
		pos int
	}{
		{[]rune("mjqjpqmgbljsphdztnvjfqwrcgsmlb"), 19},
		{[]rune("bvwbjplbgvbhsrlpgdmjqwftvncz"), 23},
		{[]rune("nppdvjthqldpwncqszvftbrmjlhg"), 23},
		{[]rune("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"), 29},
		{[]rune("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"), 26},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("example %d", i+1), func(t *testing.T) {
			got := partTwo(c.str)
			if got != c.pos {
				t.Errorf("Expected %d, got %d", c.pos, got)
			}
		})
	}
}
